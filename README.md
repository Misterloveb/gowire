# gowire

***`gowire` 是基于 **wire** (***依赖注入***) 构建的一个web项目，`wire`基于代码生成技术，编译期审查代码，更加容易发现问题，无反射，运行效率更高***
## 功能
- **Wire**: https://github.com/google/wire
- **Gin**: https://github.com/gin-gonic/gin
- **Gorm**: https://github.com/go-gorm/gorm
- **Viper**: https://github.com/spf13/viper
- **Go-redis**: https://github.com/redis/go-redis/v9
- **logrus**: https://github.com/sirupsen/logrus
- **excelize**: https://github.com/xuri/excelize/v2
- **lumberjack**: https://gopkg.in/natefinch/lumberjack.v2
- **base64Captcha**: https://github.com/mojocn/base64Captcha
- **carbon**: https://github.com/golang-module/carbon/v2
- **session**: https://github.com/gin-contrib/sessions
- **layui**: https://github.com/layui/layui

## 目录结构

借鉴 `go` 标准目录结构以及 `mvc` ：

- `cmd`：可执行文件目录，例如 `server/main.go`
    - `server`： 服务名称，每个服务对应一个目录
        - `wire`：每个服务用于 `wire` 生成 `wire_gen.go` 的入口目录
- `config`：项目配置文件，例如 `config.yaml`
- `internal`：项目内部代码，不对外暴露
    - `common`：公共使用的`函数/struct等数据`
    - `controller`：业务逻辑核心目录(包含了路由注册)
    - `dao`：数据访问对象(Data Access Object)，每张表对应一个dao
    - `middleware`：业务中的中间件
    - `model`：表结构到go结构体的映射，每张表对应一个数据模型
    - `provider`：用于`wire.Build()`进行依赖注入
    - `server`：构建 HTTP 服务器的代码
    - `test`：项目测试文件
    - `view`：html模板文件
- `pkg`：可重用的代码，对外暴露
    - `config`：用于读取配置文件信息(基于viper)
    - `http`：启动http(gin.Engine)用于支持Shutdown服务器优雅退出
    - `verificationcode`：图形验证码插件
    - `writer`：用于实现`gin`,`gorm`,`logrus`的`io.Writer`接口，支持自定义文件大小，文件保存周期，文件分割等功能，具体可以参考 [lumberjack](https://gopkg.in/natefinch/lumberjack.v2) 
- `static`：静态资源文件(css/js等)
- `upload`：存放用户上传的附件
- `log`：项目的日志文件(启动项目后自动生成)：
    - `db`：用于记录数据库操作日志(gorm)
    - `http`：用于记录HTTP请求日志(gin)
    - `server`：用于记录项目业务日志(controller等业务模块)
  
## 所需环境

* Golang >= 1.16
* Git
* MySQL >= 5.7
* Redis（可选 | 6.0）
* docker&docker-compose（可选）
* 安装 wire 并确保将 $GOPATH/bin 添加到 $PATH 中
```go
    go install github.com/google/wire/cmd/wire@latest
```
## 运行项目
- `本地`：
    - 在 `gowire` 项目根目录下执行 `go run cmd/server/main.go`
    - 打开浏览器访问 127.0.0.1:8080
- `docker`：
    - `Linux` 环境需要安装 `docker-compose`，`windows` 的`Docker Desktop` 默认已安装，下面是 Linux 安装方法，具体信息可以参考[docker官网](https://docs.docker.com/compose/install/linux/)
      - `Ubuntu 和 Debian`
      ``` shell
       sudo apt-get update
       sudo apt-get install docker-compose-plugin
      ```
      - ` RPM-based `
      ``` shell
      sudo yum update
      sudo yum install docker-compose-plugin
      ```
    - 在 `gowire` 项目根目录下执行 `docker-compose up -d`，等待构建完毕。。。🍵☕🧋
    - Linux本机浏览器访问 `127.0.0.1:8080` ，虚拟机则需在宿主浏览器访问 `虚拟机IP:8080`
## 开发流程
- `注册路由`： 路由注册统一放在 `internal/controller/router.go`，业务 `controller` 如果需要注册路由，只需各自实现 `RegisterRouter` 方法，系统会自动注册路由
- `添加controller`：
    - 一个业务模块一个文件，均放在 `internal/controller/` 目录下
    - 新增或删除一个 `contoller` 文件
        - 1： 修改 `internal/controller/router.go`文件中的 `RegisterController` 结构体(用于自动注册路由)
        - 2： 修改 `internal/provider/controller.go`，添加对应 `controller` 的provider，以及修改 `provider.CommonController` ( `CommonController` 仅仅用于减少 `wire.build()` 的代码量，使其看起来比较”优雅“🙂)
- `添加dao`：一张表一个dao文件，文件名为数据库表名，均放在`internal/dao/`目录下
    - 每新增或删除一个`dao`文件:
        - 1：修改 `internal/provider/provider.go` ，提供对应dao的provider,具体请参照 `provider.go` 
        - 2(可选)：修改 `common/server.CommonDao` (controller共用的dao)
- `重新构建项目`：如果更改了 `cmd/server/wire/wire.go` 文件或其中 `wire.Build` 所需要的 `provider` 被修改，均需要重新执行 `wire` 命令重新生成 `wire_gen.go` 文件（除第一次调用 `wire` ，以后只需要对 `wire_gen.go` 执行 `go generate` 即可🙂）

## 许可证

gowire是根据MIT许可证发布的。有关更多信息，请参见[LICENSE](LICENSE)文件。
