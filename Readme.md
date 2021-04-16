# Gin swagger demo

本项目为 gin swagger 结合使用的 demo，在此案例中，添加了相关的使用案例；具体注释格式参考来自：[https://github.com/swaggo/swag](https://github.com/swaggo/swag)

## 特性

- [x] [Air](https://github.com/cosmtrek/air) 方式自动监听文件变动，进行热重启
- [x] 使用 `tag` 方式动态选择是否在编译中携带 `swagger`，在本地或测试服情况可以使用 `go build -tags "doc" ."` 方式编译，提供接口文档；在生产模式下，可以移除 `tag`，构建无 `swagger` 的主程序；
- [x] 接入 GORM，使用 swaggertype 更改字段类型
- [x] 添加 Security，实现登录
- [x] 对 swagger 相关路由添加 Basic Auth 校验
- [x] 编写若干个相关示例

## 如何运行

进入到项目目录 `github.com/happyphper/gin-swagger-demo`，执行：

```
air
```