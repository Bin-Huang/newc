package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

func main() {
	pkg, err := GetPackageInfo(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg.Name, pkg.GoFiles)
	// TODO: improve performance by multi tasks
	for _, filepath := range pkg.GoFiles {
		has, err := IncludeMakeMark(filepath)
		if err != nil {
			panic(err)
		}
		fmt.Println(filepath, has)
	}
}

// GetPackageInfo get the Go package information in the dir
func GetPackageInfo(dir string) (*packages.Package, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode:  packages.NeedName | packages.NeedFiles,
		Tests: false,
	}, dir)
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("cannot find any package in %v", dir)
	}
	return pkgs[0], nil
}

// IncludeMakeMark ...
func IncludeMakeMark(filepath string) (bool, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if IsMakeComment(line) {
			return true, nil
		}
	}
	return false, nil
}

// IsMakeComment ...
func IsMakeComment(s string) bool {
	s = strings.TrimSpace(s)
	return strings.HasPrefix(s, "//go:generate") && strings.Contains(s, "make-constructor")
}
