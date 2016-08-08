# ApkInfGo
Golang parser for android apk files. It extracts version information, package name, etc.

##Example
```go
package main

import "fmt"
import "github.com/Mr4Mike4/ApkInfGo"

func main() {
	aapt := "d:\\Android\\sdk\\build-tools\\24.0.0\\aapt.exe"
    keytool := "c:\\Program Files\\Java\\jdk1.8.0_92\\bin\\keytool.exe"

    // without information on the certificate
	//apk := ApkInfGo.ApkInfo(app).File("d:\\downloads\\file.apk")

	apk := ApkInfGo.ApkInfo(app).CertKeyTool(keytool).File("d:\\downloads\\file.apk")
	fmt.Printf("%q\n", *apk)

	apks := ApkInfGo.ApkInfo(app).CertKeyTool(keytool).Folder("d:\\downloads", true)
	if apks != nil {
		for _, file := range *apks {
			fmt.Printf("%q\n", file)
		}
		fmt.Printf("Count - %d\n", len(*apks))
	}

	cert := ApkInfGo.ApkCertificate(keytool).File("d:\\downloads\\file.apk")
    fmt.Printf("%q\n", *cert)
}
```
