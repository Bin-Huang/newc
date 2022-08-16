# newc

Doc: **English** | [中文](README_zh.md)

----------

A cli tool to generate constructor code for a Golang struct.

## Installation

```bash
go install github.com/Bin-Huang/newc@latest
```

## Usage

Add a `go:generate` command line to the struct which you want to generate a constructor.

```go
//go:generate newc
type UserService struct {
	baseService
	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository
}
```

After executing `go generate ./...` the constructor code generated:

```go
// constructor_gen.go

// NewUserService Create a new UserService
func NewUserService(baseService baseService, userRepository *repositories.UserRepository, proRepository *repositories.ProRepository) *UserService {
	return &UserService{
		baseService:    baseService,
		userRepository: userRepository,
		proRepository:  proRepository,
	}
}
```

See [more examples here](https://github.com/Bin-Huang/newc/tree/master/test)

## Usage without manual installation

**Recommended for team collaboration**

Without manual installation, just add this comment line to the struct. Go will automatically install this tool if missing.

```go
//go:generate go run github.com/Bin-Huang/newc@v0.8.3
```

For example:

```go
//go:generate go run github.com/Bin-Huang/newc@v0.8.3
type UserService struct {
	baseService
	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository
}
```

This is very useful, especially in teamwork. **It can run without manual installation. It doesn't break the work of other people who don't have installed this tool in collaboration.**

## How to return value instead reference?

Add `--value` parameter

```go
//go:generate newc --value
type Config struct {
	debug  bool
}
```

Generated code:

```go
// constructor_gen.go

// NewConfig Create a new Config
func NewConfig(debug bool) Config {
	return Config{
		debug:  debug,
	}
}
```

## How to call an initializer in constructor?

1. Add `--init` parameter
2. Write an `init` method for the struct

```go
//go:generate newc --init
type Controller struct {
	logger *zap.Logger
	debug  bool
}

func (c *Controller) init() {
	c.logger = c.logger.With(zap.String("tag", "controller-debugger"))
	c.debug = true
}
```

Generated code:

```go
// constructor_gen.go

// NewController Create a new Controller
func NewController(logger *zap.Logger, debug bool) *Controller {
	s := &Controller{
		logger: logger,
		debug:  debug,
	}
	s.init()
	return s
}
```

## How to ignore some fields when generating constructor code?

Add a tag `newc:"-"` to fields that need to be ignored

```go
type Forbidden struct {
	Msg    string
	Status int    `newc:"-"`
}
```

Generated code:

```go
// NewForbidden Create a new Forbidden
func NewForbidden(msg string) *Forbidden {
	return &Forbidden{
		Msg: msg,
	}
}
```

## If you think the `go:generate` comment is too long...

Some suggestions:

1. Add a code snippest in your editor/IDE for the tool (suggested)
2. ......

## Features & Motivation

**1. It makes your code easier to write and maintain**.

Writing and updating constructor code for many structs can be laborious and error-prone, especially if you have a huge codebase. These should be handed over to automatic tools like this tool.

And it also works well with these dependency injection tools like [`wire`](https://github.com/google/wire). If you use `wire` in your project, you may need this tool very much.

**2. It takes care of the generated code**.

Don't worry about the imports, variable naming, and code style in the generated code.

**3. It is more suitable for teamwork**.

It doesn't break the work of other people who don't have installed this tool in collaboration. Go will automatically install this tool if missing.

```go
//go:generate go run github.com/Bin-Huang/newc@v0.8.3
```

## Sponsoring

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://buymeacoffee.com/benn)

![](./doc/donate.png)

## License

MIT
