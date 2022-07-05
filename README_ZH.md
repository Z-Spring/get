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
也可以用 get search [package]命令来搜索相关的包

