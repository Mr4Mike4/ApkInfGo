# ApkInfGo
Golang parser for android apk files. It extracts version information, package name, etc.

##Example
```go
package main

import "fmt"
import "github.com/Mr4Mike4/ApkInfGo"

func main() {
	app := "d:\\Android\\sdk\\build-tools\\24.0.0\\aapt.exe"

	apk := ApkInfGo.ApkInfo(app).File("d:\\downloads\\file.apk")
	fmt.Printf("%q\n", *apk)

	apks := ApkInfGo.ApkInfo(app).Folder("d:\\downloads", true)
	if apks != nil {
		for _, file := range *apks {
			fmt.Printf("%q\n", file)
		}
		fmt.Printf("Count - %d\n", len(*apks))
	}
}
```
