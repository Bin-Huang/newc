package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go/ast"
	"go/types"

	"go/parser"
	"go/token"

	"golang.org/x/tools/go/packages"
)

var fset = token.NewFileSet()

func main() {
	pkg, err := GetPackageInfo(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg.Name, pkg.GoFiles)
	for _, filename := range pkg.GoFiles {
		has, err := IncludeMakeMark(filename)
		if err != nil {
			panic(err)
		}
		if has {
			result, err := ParseFile(filename)
			if err != nil {
				panic(err)
			}
			if len(result) == 0 {
				continue
			}
			fmt.Println(result)
		}
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

// BuildAST ...
func BuildAST(filename string) (*ast.File, error) {
	astFile, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to build AST from file(%v): %w", filename, err)
	}
	return astFile, nil
}

// ResultField ...
type ResultField struct {
	Name string
	Type string
}

// Result ...
type Result struct {
	StructName string
	Fields     []ResultField
}

// ParseFile ...
func ParseFile(filename string) ([]Result, error) {
	ret := []Result{}
	astFile, err := BuildAST(filename)
	if err != nil {
		return ret, err
	}
	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.TYPE {
			continue
		}

		needGen := false
		for _, doc := range genDecl.Doc.List {
			if IsMakeComment(doc.Text) {
				needGen = true
				break
			}
		}
		if !needGen {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			resultFields := []ResultField{}
			for _, field := range structType.Fields.List {
				fieldType := types.ExprString(field.Type)
				var fieldName string
				if len(field.Names) > 0 {
					fieldName = field.Names[0].Name
				} else {
					// handle embeded struct cases just like this:
					// 		type Foo struct {
					//  		pkg.Struct,
					// 		}
					items := strings.Split(fieldType, ".")
					fieldName = items[len(items) - 1]
					// handle pointer cases just like this:
					// 		type Foo struct {
					//  		*pkg.Struct,
					// 		}
					fieldName = strings.TrimPrefix(fieldName, "*")
				}
				resultFields = append(resultFields, ResultField{
					Type: fieldType,
					Name: fieldName,
				})
			}
			ret = append(ret, Result{
				StructName: typeSpec.Name.Name,
				Fields:     resultFields,
			})
		}
	}
	return ret, nil
}
