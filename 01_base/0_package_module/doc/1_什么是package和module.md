# 1.package
## 1.1.什么是package
在go语言中，package是用来组织源码（即多个go源码的集合），作为代码复用的基础(比如fmt、os、io)。<br/> 
每个源码文件都需要声明其所属的package(package所属可以和文件夹名不同--支持自定义)，并且同一个文件夹下的源码文件都具有相同的package。<br/> 
## 1.2.import的使用
同文件夹下无需导入可直接使用，import主要是解决不同文件夹下的导入问题。主要有如下方式：
- 导入包路径
- 导入包路径并定义别名: 为了解决不同路径下package名重复的问题
- 导入包路径下所有（即别名为.）: 可以达到同文件夹下使用的效果，但不推荐使用
- 导入匿名包路径（即别名为_）: 可以解决不使用报错的问题以及用于初始化

# 2.module
## 2.1.什么是module
一个模块是一组相关的Go包的集合，它是依赖管理的基本单元。一个项目要成为模块，只需在根目录下包含一个go.mod文件——该文件记录了模块的标识、依赖项
及其版本、替换规则等核心信息。它是自动维护的，不要去随意修改。<br/>
模块的关键特性：
- 项目可放在任意目录（无需局限于GOPATH/src）；
- 依赖版本明确可控，支持多版本并存；
- 通过语义化版本管理依赖兼容性。

## 2.2.设置国内镜像代理
查看go环境信息：
```shell
go env
```
设置：
```shell
# 设置 GO111MODULE = on
go env -w GO111MODULE=on

# 设置国内镜像源 GOPROXY
# 阿里云：https://mirrors.aliyun.com/goproxy/
# 七牛云：https://goproxy.cn
# 原地址：https://goproxy.com
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```
## 2.3.go mod的相关命令
```shell
# 查看当前项目的所有依赖
go list -m all

# 查看某个具体依赖的可用版本
go list -m -versions github.com/gin-gonic/gin

# 修改项目中的某个依赖的具体版本:  使用@符号后面跟版本号（比如当前github.com/gin-gonic/gin 使用版本为v1.11.0）
# go get命令会修改go.mod文件，可在项目的go.mod文件中查看
go get github.com/gin-gonic/gin@v1.10.1

# 使用go get不指定具体版本，默认下载最新的依赖
go get github.com/go-redis/redis/v8
# 使用go get -u 升级到最新的次要版本或修订版本
go get -u github.com/gin-gonic/gin
# 使用go get -u=patch 升级到最新的修订版本
go get -u=patch github.com/gin-gonic/gin

# go mod的帮助
go mod help

# 使用go mod查看依赖
go mod graph

# 使用go mod初始化go.mod管理文件
go mod init

# 更常用的tidy： 用于整理项目中的所有依赖--去除没有使用的依赖并下载
go mod tidy

# 远程仓库名为project-A， 项目中的go.mod配置的module为github.com/bobby/A
# 当两者不一致时，我们如果想通过module配置的去下载，就可以通过replace的方式
# 也就是说go get github.com/bobby/A实际会是go get github.com/bobby/project-A@1.0.0
go mod edit -replace github.com/bobby/A=github.com/bobby/project-A@v1.0.0
```