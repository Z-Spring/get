![img.png](img.png)
# Get

you can get go packages easily

[中文文档](https://github.com/Z-Spring/get/blob/master/README_ZH.md)

Do you have this problem which you want to begin a new go project, first,
you must go get some packages, like gin, redis, gorm and so on...

example:

*go get github.com/gin-gonic/gin/*  <p>
*go get gorm.io/gorm/*  <p>
*go get github.com/go-redis/redis/v9/*  <p>
every time get new packages, you must remember its whole package path,
that's so terrible. <p>

Now, you can use getcli to simplify this process.   <p>
When you want to get gin, just ***get gin***, that's so easy.


## Install
```bash
go install github.com/z-spring/get
```
## Usage
> two commands you can use
* get search [package]
* get [package]

you can use get [package] command to get  packages  <p>
you can use get search [package] command to search  packages
