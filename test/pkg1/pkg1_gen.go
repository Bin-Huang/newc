package pkg1

import (
	"fmt"

	"io"

	. "os"

	stringutils "strings"
)

// NewService Create a new Service
func NewService(Name string) *Service {
	return &Service{
		Name: Name,
	}
}

// NewPostService Create a new PostService
func NewPostService(Service Service, Version int) *PostService {
	return &PostService{
		Service: Service,
		Version: Version,
	}
}

// NewAgeService Create a new AgeService
func NewAgeService(Service Service, Age int, Writer io.Writer, File File, AA stringutils.Builder, Stringer fmt.Stringer) *AgeService {
	return &AgeService{
		Service:  Service,
		Age:      Age,
		Writer:   Writer,
		File:     File,
		AA:       AA,
		Stringer: Stringer,
	}
}
