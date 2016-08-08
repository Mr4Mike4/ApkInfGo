package ApkInfGo

import (
	"log"
	"os/exec"
	"strings"
)
type ApkCertSt struct {
	CertMD5 string
	CertSHA1 string
	CertSHA256 string
	SignAlgorithm string
	FilePath string
}

type ConfCert struct {
	keytool string
}

func ApkCertificate(keytoolApp string) *ConfCert {
	c := &ConfCert{keytool:keytoolApp}
	return c
}

func (c *ConfCert) File(apk string) *ApkCertSt {
	out, err := exec.Command(c.keytool, "-list", "-printcert", "-jarfile", apk).Output()
	if err != nil {
		log.Printf("err: %q, file: %q", err, apk)
		return nil
	}
	//log.Printf("apk file - %q\n", apk)
	data := strings.Split(string(out), "\n")
	cert := ApkCertSt{FilePath:apk}
	for _, s := range data{
		arr := strings.Split(s, ": ")
		if len(arr) != 2 {
			//log.Printf("error split - %q\n", s)
			continue
		}
		arr[0] = strings.Trim(arr[0], " \t")
		switch arr[0] {
		case "MD5":
			//log.Printf("MD5 - %q\n", arr[1])
			cert.CertMD5 = strings.Trim(arr[1], " ")
			break
		case "SHA1":
			//log.Printf("SHA1 - %q\n", arr[1])
			cert.CertSHA1 = arr[1]
			break
		case "SHA256":
			//log.Printf("SHA256 - %q\n", arr[1])
			cert.CertSHA256 = arr[1]
			break
		case "Signature algorithm name":
			//log.Printf("Signature algorithm name - %q\n", arr[1])
			cert.SignAlgorithm = arr[1]
			break
		//default:
		//	log.Printf("%q - %q\n", arr[0], arr[1])
		//	break
		}
	}
	return &cert
}