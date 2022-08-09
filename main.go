package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	err := makeConstructor()
	if err != nil {
		fmt.Printf("make-constructor: [ERROR] %v\n", err)
		os.Exit(1)
	}
}

func makeConstructor() error {
	pkg, err := GetPackageInfo(".")
	if err != nil {
		return err
	}
	// skip if generated recently
	genFilename := "./constructor_gen.go"
	if isGeneratedRecently(genFilename) {
		return nil
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
		return nil
	}
	code, err := GenerateCode(pkg.Name, allImports, allResults)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(genFilename, []byte(code), 0644)
	if err != nil {
		return err
	}
	genFilepath, err := filepath.Abs(genFilename)
	if err != nil {
		return err
	}
	fmt.Printf("make-constructor: [INFO] wrote %v\n", genFilepath)
	return nil
}

func isGeneratedRecently(genFilename string) bool {
	stat, err := os.Stat(genFilename)
	if err != nil {
		return false
	}
	return time.Now().Sub(stat.ModTime()) < 5*time.Second
}
