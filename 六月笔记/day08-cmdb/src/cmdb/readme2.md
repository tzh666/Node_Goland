程序部署:
go build (得把非go的代码复制到同一级别目录下,例如conf,views,static,sql(建议))

交叉编译:
    windows上编译,部署到Linux上:
        1、set GOOS=ilnux, 改回去set GOOS=windows
        2、go build (得把非go的代码复制到同一级别目录下,例如conf,views,static,sql(建议))

配置HTTPS:
    https://www.jianshu.com/p/640bf4bce823  （p12证书,不在浏览器导入可以忽略）
    
    app.conf 配置:
        EnableHTTP=true
        EnableHTTPS=true
        HTTPSCertFile=conf/ssl/server.crt
        HTTPSKeyFile=conf/ssl/server.key

cobra:  
子命令实现
    go run .\main.go db --help
    go run .\main.go user
    go run .\main.go web
全局命令
    rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
    go run .\main.go web --help


用户权限:


  
