# make-constructor

Doc: [English](README.md) | **中文**

----------

一个自动为 Go 结构体生成类似 `NewXXX` 构造器函数代码的命令行工具。它不需要手动安装，只需要在结构体上加一行代码注释就能工作。

## 如何使用？

它的使用方式非常简单，不用专门手动安装，只需要在结构体添加下面这行代码注释就能工作。

```go
//go:generate go run github.com/Bin-Huang/make-constructor@v0.7.0
```

举个例子：

```go
//go:generate go run github.com/Bin-Huang/make-constructor@v0.7.0
type UserService struct {
	baseService
	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository
}
```

当你执行 `go generate ./...`, `go test` 或者 `go build` 后，你就能得到下面生成的代码文件：

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

这里可以[查看更多例子](https://github.com/Bin-Huang/make-constructor/tree/master/test)

## 它可以本地安装吗？

其实如果用上面的方法，Go 会自动帮你本地安装，当然你也可以手动安装：

```go
go get -u github.com/Bin-Huang/make-constructor
```

然后你就可以这么使用它：

```go
//go:generate make-constructor
type UserService struct {
	baseService
	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository
}
```

## 想在构造时做些初始化?

1. 加上 `--init` 参数
2. 为结构体实现一个 `init` 方法

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

生成代码：

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

## 如果你觉得这条注释太长……

一些建议：

1. (推荐）把它加进你的编辑器/IDE的快捷代码片段（code snippest）里
2. 手动安装这个工具

## 功能特性与设计理念

**1. 它能让你的代码更容易编写和维护**.

不管是编写还是更新构造器代码，都是一个费力且容易出错的事情，尤其当代码量很大的时候。这些繁琐易错的工作应该交给自动程序来完成，比如这个工具。

同时，这个工具还能完美兼容像[**wire**](https://github.com/google/wire)这种依赖注入工具。如果你的项目中也使用了 **wire**，那你可能非常需要这个工具。**wire** 在 **make-constructor** 的“加持”下会变得更加好用。

**2. 你不需要担心自动生成的代码**.

这个工具在生成代码时会非常小心，会帮你考虑所有代码细节，包括引用依赖、变量命名，甚至还有代码风格。

**3. 它不需要手动安装，也不需要引用其他依赖**.

只要有 GO 环境和网络的地方，这个工具就能正常工作。你在项目中使用这个工具不会影响到其他同事，就算他们没有安装这个工具，代码的自动生成也不会有任何问题。

（这个工具运行在本地，网络只用于在必要时下载这个工具）

## 赞赏

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://buymeacoffee.com/benn)

![](./doc/donate.png)

## License

MIT
