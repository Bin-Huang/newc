package main

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func main() {
	pkg, err := GetPackageInfo(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg.Name, pkg.GoFiles)
}

// GetPackageInfo ...
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
