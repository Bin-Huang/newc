package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	pkg, err := GetPackageInfo(".")
	if err != nil {
		panic(err)
	}
	// skip if generated recently
	genFilename := "./constructor_gen.go"
	if isGeneratedRecently(genFilename) {
		return
	}
	allImports := []ImportInfo{}
	allResults := []StructInfo{}
	for _, filename := range pkg.GoFiles {
		has, err := IncludeMakeMark(filename)
		if err != nil {
			panic(err)
		}
		if !has {
			continue
		}
		results, imports, err := ParseCodeFile(filename)
		if err != nil {
			panic(err)
		}
		if len(results) == 0 {
			continue
		}
		allImports = append(allImports, imports...)
		allResults = append(allResults, results...)
	}
	if len(allResults) == 0 {
		return
	}
	code, err := GenerateCode(pkg.Name, allImports, allResults)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(genFilename, []byte(code), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("make-constructor: %v: wrote %v\n", pkg.PkgPath, genFilename)
}

func isGeneratedRecently(genFilename string) bool {
	stat, err := os.Stat(genFilename)
	if err != nil {
		return false
	}
	return time.Now().Sub(stat.ModTime()) < 5*time.Second
}
