![img.png](img.png)
# Get

you can get go packages easily

[中文文档](https://github.com/Z-Spring/get/blob/master/README_ZH.md)

Do you have this problem which you want to begin a new go project, first,
you must go get some packages, like `gin`, `redis`, `gorm` and so on...

example:

***go get github.com/gin-gonic/gin/***  <p>
***go get gorm.io/gorm/***  <p>
***go get github.com/go-redis/redis/v9/***  <p>

Every time get new packages, you must remember its whole package path,
that's so terrible. <p>

Now, you can use getcli to simplify this process.   <p>
When you want to get gin, just `get gin`, that's so easy.


## Install
```bash
go install github.com/z-spring/get@latest
```
## Usage
> two commands you can use
* `get search` [package]
* `get` [package] <p>

you can use `get [package]` command to get  packages  <p>

```bash
get gin
```
you can use `get search [package]` command to search  packages

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
