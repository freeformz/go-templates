package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	base := filepath.Join(wd, "shared", "templates")

	// The returned template is named "a.gotmpl" and contains the following templates:
	// 'b.gotmpl' == the entirety of the file.
	// 'bb' == the defined template in the `b.gotmpl` file.
	// associated template named "bb", which was loaded from b.gotmpl
	a, err := template.ParseFiles(filepath.Join(base, "a.gotmpl"), filepath.Join(base, "b.gotmpl"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Template from a.gotmpl:", a.Name())

	b := a.Lookup("b.gotmpl")
	fmt.Println("Template from b.gotmpl:", b.Name())

	bb := a.Lookup("bb")
	fmt.Println("Template defined inside of b.gotmpl:", bb.Name())
	fmt.Println("Output of Executing 'a'")
	fmt.Println("-------------------------------------------")
	if err := a.Execute(os.Stdout, "no data"); err != nil {
		panic(err)
	}
}
