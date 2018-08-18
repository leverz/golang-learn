# golang-learn
学习 golang 产生的代码

# long, long ago
golang 项目有一个固定的目录结构：
假设项目名 foo, foo 下必须要有 src、pkg、bin 三个子目录
src 用来放源码
pkg 用来放置编译的中间文件
bin 用来放编译后的可执行文件

golang 的包管理是个大坑，目前的解决方案是 govendor ，跟我大 node 的 npm 比起来简直垃圾。

```
go get github.com/kardianos/govendor
```

安装 govendor 到全局

第一步设置项目的 $GOPATH 和 $GOBIN

```
export GOPATH=/xx/xx.../foo
export GOBIN=$GOPATH/bin
```

第二步

```
cd foo/src
mkdir bar
cd bar
govendor init
```

其中 bar 就是真正的代码目录，govendor 会在其中产生一个 vendor 目录，用来放置依赖的包的代码，类似于 node 中的 node_modules



# 2018.08.18
golang 的 $GOPATH 可以设置多个，在 unix 系统中用 : 分割，当有多个GOPATH时，默认会将go get的内容放在第一个目录下。

又发现一个比较不错的 golang 包管理工具，star 数比 govendor 还多，叫 [glide](https://glide.sh/)

Go 语言标准库中有一个代码包专门用于接收和解析命令参数。这个代码包的名字叫 **flag**

同一个目录下的源码文件都需要被声明为属于同一个代码包。

源码文件声明的包名可以与其所在目录的名称不同，只要这些文件声明的包名一致就好。

### 问题：import 两个不同路径，但是最后一级名称相同的包，是否会发生冲突？
先举个例子：dep/lib/flag 和 flag 这两个包，最后一级名称都是 flag

回答：如果两个包中定义了相同的 package 名，则会冲突，需要在 import 的时候起别名，比如：

```
import (
    "dep/lib/flag"
    f "flag"
)
```

如果两个包定义了不同的 package 名，则不会冲突

