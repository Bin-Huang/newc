package pkg1

import (
	"fmt"
	"io"
	. "os"
	stringutils "strings"
)

// F ...
func F(a interface{}) {
	s := fmt.Sprintln(a)
	s = stringutils.ToLower(s)
	fmt.Println(s)
}

// Service ...
//go:generate go run ../../../make-constructor
type Service struct {
	Name string
}

// PostService ...
//go:generate go run ../../../make-constructor
type PostService struct {
	Service
	Version int
}

// AgeService ...
//go:generate go run ../../../make-constructor
type AgeService struct {
	Service
	Age      int
	Writer   io.Writer
	File     File
	AA       stringutils.Builder
	Stringer fmt.Stringer
}
