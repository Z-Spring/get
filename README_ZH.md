![img.png](img.png)
# Get

更便捷的导入包

[English](https://github.com/Z-Spring/get/blob/master/README.md)

你是否有这样的困扰，当你想要开始一个新项目时，首先要导入各种包，比如说 
gin, redis, gorm 等等。

比如: 

*go get github.com/gin-gonic/gin/*  <p>
*go get gorm.io/gorm/*  <p>
*go get github.com/go-redis/redis/v9/*  <p>

但是每次导入新包时，必须要记住包的路径，或者去网上查找，很是麻烦。 <p>

现在你可以用 getcli 来简化这个步骤，当你想导入gin时，直接用***get gin*** 就可以了


## 安装
```bash
go install github.com/z-spring/get
```
## 用法
> 有两个命令可以用
 * get search [package]
 * get [package]


你可以用 get [package]来导入包  <p>

```bash
get gin
get redis
...
```
也可以用 get search [package]命令来搜索相关的包

```bash
get search gin

$ get search gin
NAME            PKG                                             IMPORTED
gin             github.com/gin-gonic/gin                        31,327
cors            github.com/gin-contrib/cors                     1,054
ginSwagger      github.com/swaggo/gin-swagger                   620
gzip            github.com/gin-contrib/gzip                     199
jwt             github.com/appleboy/gin-jwt/v2                  166
pprof           github.com/gin-contrib/pprof                    259
sessions        github.com/gin-contrib/sessions                 602
gin             gopkg.in/gin-gonic/gin.v1                       212
static          github.com/gin-contrib/static                   273
cache           github.com/gin-contrib/cache                    42
gin             github.com/luraproject/lura/v2/router/gin       25
gin             github.com/fixbanking/gin                       24
ginzap          github.com/gin-contrib/zap                      84
ginprometheus   github.com/zsais/go-gin-prometheus              55
requestid       github.com/gin-contrib/requestid                31
gintemplate     github.com/foolin/gin-template                  44
multitemplate   github.com/gin-contrib/multitemplate            66
favicon         github.com/thinkerou/favicon                    65
```