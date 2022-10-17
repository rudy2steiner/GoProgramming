# Go programming language book practice


## Go get started

## Go mod 依赖管理
将当前目录变成一个go 项目:go mod init xxxx,后续项目内所有的包引用以它为前缀

例如 go mod init hello.com，想引用 ch1/echo 包

```
    import hello.com/ch1/echo
```
