# make-constructor

A command-line tool to generate constructor code for a struct. It don't need manual installation, just add a comment line to the struct then it works.

## How to use?

It don't need a manual installation. Just add this comment line to the struct you want to generate a constructor.

```go
//go:generate go run github.com/Bin-Huang/make-constructor@v0.1.0
```

For example:

```go
// user_servie.go

//go:generate go run github.com/Bin-Huang/make-constructor@v0.1.0
type UserService struct {
	baseService
	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository
}
```

after `go generate ./...`, `go test` or `go build`, you get this:

```go
// user_servie_gen.go

// NewUserService Create a new UserService
func NewUserService(baseService baseService, userRepository *repositories.UserRepository, proRepository *repositories.ProRepository) *UserService {
	return &UserService{
		baseService:    baseService,
		userRepository: userRepository,
		proRepository:  proRepository,
	}
}
```

See [more examples here](https://github.com/Bin-Huang/make-constructor/tree/master/test)

## Can it be installed locally?

Actually Go will automatically installs it locally, but you can also install manually. 

```go
go get -u github.com/Bin-Huang/make-constructor
```

Now you can use it like this:

```go
//go:generate make-constructor
type UserService struct {
	baseService
	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository
}
```
