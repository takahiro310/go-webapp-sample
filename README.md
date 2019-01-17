# go-webapp-sample
GolangでWebアプリを作ってみるの巻

### Using Middlewares

| Target          | Middleware |
| :--             | :--        |
| Web Framework   | Gin        |
| Migration Tool  | goose      |

goose
---
> go get bitbucket.org/liamstask/goose/cmd/goose

Windows環境を利用している方は、あらかじめ[GCC](http://tdm-gcc.tdragon.net/)をインストールが必要。
go get後にgooseコマンドが見つからない場合は、go envでGOPATHを確認し、環境変数のPATHにGOPATHの値を設定する。

xo
---
```
$ go get -u golang.org/x/tools/cmd/goimports
$ go get -u github.com/xo/xo
$ go get -u github.com/go-sql-driver/mysql
$ mkdir src/model
$ mkdir -p db/xo/templates
$ cp $GOPATH/src/github.com/xo/xo/templates/mysql* db/xo/templates/
$ cp $GOPATH/src/github.com/xo/xo/templates/xo* db/xo/templates/
$ xo mysql://root:root@localhost/sample -o src/model/ --template-path db/xo/templates/
```

なぜかモデル出力でエラーが発生する。
```
error: model/goosedbversion.xo.go:147:1: expected declaration, found 'IDENT' postgres
model/goosesample.xo.go:140:1: expected declaration, found 'IDENT' postgres
model/test.xo.go:145:1: expected declaration, found 'IDENT' postgres
```
対象のファイルの末尾に "postgres.index.go.tpl" が挿入されている行がコンパイルエラーになるので消すと、一応は動く・・・。

環境は go version go1.9.2 windows/amd64 なんですけど、だれかご存知の方おしえてください(;´Д｀)
