# Go programming language in action


## Go get started

## Go mod 依赖管理
将当前目录变成一个go 项目:go mod init xxxx,后续项目内所有的包引用以它为前缀

例如 go mod init hello.com，想引用 ch1/echo 包

```
    import hello.com/ch1/echo
```

## build and run 
``` 
 go mod tidy
 
 go build hello.com/ch2/cf 
 or
 go build ./ch2/cf
 
 ./cf 32
```
or
``` 
go run hello.com/ch2/cf 32
```
