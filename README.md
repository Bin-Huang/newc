# make-constructor

Doc: **English** | [中文](README_zh.md)

A command-line tool to generate constructor code for a struct. It don't need manual installation, just add a comment line to the struct then it works.

## How to use?

It don't need a manual installation. Just add this comment line to the struct you want to generate a constructor.

```go
//go:generate go run github.com/Bin-Huang/make-constructor@v0.7.0
```

For example:

```go
//go:generate go run github.com/Bin-Huang/make-constructor@v0.7.0
type UserService struct {
	baseService
	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository
}
```

After `go generate ./...` you will get this:

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

## Wanna initialize something in the constructor?

1. Add `--init` parameter
2. Write an `init` method for the struct

```go
//go:generate go run github.com/Bin-Huang/make-constructor@v0.7.0 --init
type Controller struct {
	logger *zap.Logger
    debug  bool
}

func (c *Controller) init() {
	c.logger = c.logger.With(zap.String("tag", "this-special-controller"))
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

## If you think the "magic comment" is too long...

Some suggestions:
1. Add a code snippest in your editor/IDE for the tool (suggested)
2. [Install this tool manually](#can-it-be-installed-locally)

## Features & Motivation

**It makes your code easier to write and maintain**.

Writing and updating constructor code for many structs can be laborious and error-prone, especially if you have a huge codebase. These should be handed over to automatic tools like this tool.

And it also works well with these dependency injection tool like [`wire`](https://github.com/google/wire). That is to say, if you use `wire` in your project, you may need this tool very much.


**It take care of the generated code**.

Don't worry about the imports, variable naming and code style in the generated code.

**It dont't need manual installation and other dependency**.

It works anywhere there is a GO runtime and network. It don't broke the work of other people who don't have installed this tool in collaboration.

## Sponsoring

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://buymeacoffee.com/benn)

![](./doc/donate.png)

## License

MIT
