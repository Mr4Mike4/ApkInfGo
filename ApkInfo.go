package ApkInfGo

import (
	"log"
	"os/exec"
	"strings"
	"regexp"
	"io/ioutil"
	"strconv"
	"os"
)
type ApkInfoSt struct {
	Name string
	VersionCode uint32
	VersionName string
	Label string
	Icon string
	SdkVersion uint16
	TargetSdkVersion uint16
	FileSize int64
	FilePath string
	Cert ApkCertSt
}

type Conf struct {
	aapt string
	cert *ConfCert
}

func ApkInfo(aaptApp string) *Conf {
	c := &Conf{aapt:aaptApp, cert:nil}
	return c
}

func (c *Conf) CertKeyTool(keytoolApp string) *Conf {
	app := ApkCertificate(keytoolApp)
	c.cert = app
	return c
}

func (c *Conf) File(apk string) *ApkInfoSt {
	d := parse(c, apk)
	if d != nil {
		file, _ := os.Stat(apk)
		d.FileSize = file.Size()
	}
	return d
}

func parse(c *Conf, apk string) *ApkInfoSt {
	out, err := exec.Command(c.aapt, "dump", "badging", apk).Output()
	if err != nil {
		log.Printf("err: %q, file: %q", err, apk)
		return nil
	}
	//log.Printf("apk file - %q\n", apk)
	data := strings.Split(string(out), "\n")
	info := ApkInfoSt{FilePath:apk}
	for _, s := range data{
		arr := strings.Split(s, ":")
		if len(arr) != 2 {
			//log.Printf("error split - %q\n", s)
			continue
		}
		switch arr[0] {
		case "package":
			//log.Printf("package - %q\n", arr[1])
			re := regexp.MustCompile("name='([^']+)?' versionCode='(\\d*)?' versionName='([^']+)?'")
			packageInfo := re.FindStringSubmatch(arr[1])
			info.Name = packageInfo[1]
			info.VersionName = packageInfo[3]
			versionCode, _ := strconv.ParseUint(packageInfo[2], 0, 32)
			info.VersionCode = uint32(versionCode)
			break
		case "sdkVersion":
			//log.Printf("sdkVersion - %q\n", arr[1])
			sdkVersion, _ := strconv.ParseUint(strings.Trim(arr[1], "'"), 0, 16)
			info.SdkVersion = uint16(sdkVersion)
			break
		case "targetSdkVersion":
			//log.Printf("targetSdkVersion - %q\n", arr[1])
			targetSdkVersion, _ := strconv.ParseUint(strings.Trim(arr[1], "'"), 0, 16)
			info.TargetSdkVersion = uint16(targetSdkVersion)
			break
		case "application":
			//log.Printf("application - %q\n", arr[1])
			re2 := regexp.MustCompile("label='([^']+)?' icon='([^']+)?'")
			d := re2.FindStringSubmatch(arr[1])
			info.Label = d[1]
			info.Icon = d[2]
			break
		//default:
		//	log.Printf("%q - %q\n", arr[0], arr[1])
		//	break
		}
	}
	if c.cert != nil {
		info.Cert = *c.cert.File(apk)
	}
	return &info
}

func (c *Conf) Folder(dirname string, recurcive bool) *[]ApkInfoSt {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Printf("err: %q", err)
		return nil
	}
	//log.Printf("apk folder - %q\n", dirname)
	infoArr := make([]ApkInfoSt, 0)
	re := regexp.MustCompile(".*\\.apk")
	for _, file := range files {
		if re.MatchString(file.Name()) {
			dir := dirname + string(os.PathSeparator) + file.Name()
			a := parse(c, dir)
			if a != nil {
				a.FileSize = file.Size()
				infoArr = append(infoArr, *a)
			}
		} else if file.IsDir() && recurcive {
			dir := dirname + string(os.PathSeparator) + file.Name()
			//log.Printf("apk subfolder - %q\n", dir)
			arr := (c).Folder(dir, true)
			if arr != nil {
				for _, a := range *arr {
					infoArr = append(infoArr, a)
				}
			}
		}
	}
	return &infoArr
}
