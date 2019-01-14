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
