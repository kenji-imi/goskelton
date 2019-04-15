# goskelton

## install

```sh
$ go get -u github.com/kenji-imi/goskelton
```

## help

```sh
$ goskelton -h
NAME:
   goskelton - goskelton

USAGE:
   goskelton [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --project value  generated project name
   --user value     git user name for package import path in skelton file [$GOSKELTON_USER]
   --dest value     path for under which directory to create project skelton (default: ".") [$GOSKELTON_DEST_DIR]
   --help, -h       show help
   --version, -v    print the version
```

## setup project

create project

```sh
$ goskelton --user kenji-imi --dest . --project sample-prj
2019/04/15 10:50:35 [INFO] Created ./sample-prj
2019/04/15 10:50:35 [INFO] Created ./sample-prj/Makefile
2019/04/15 10:50:35 [INFO] Created ./sample-prj/main.go
2019/04/15 10:50:35 [INFO] Created ./sample-prj/src/hello
2019/04/15 10:50:35 [INFO] Created ./sample-prj/src/hello/hello.go
2019/04/15 10:50:35 [INFO] Created ./sample-prj/src/hello/hello_test.go
$ cd sample-prj
$ tree
.
├── Makefile
├── go.mod
├── main.go
└── src
    └── hello
        ├── hello.go
        └── hello_test.go

2 directories, 5 files
```

init modules

```sh
$ make init_mod
go mod init github.com/kenji-imi/sample-prj
go: creating new go.mod: module github.com/kenji-imi/sample-prj
```

## run

```sh
$ go run main.go
Hello!
```
```sh
$ make test_unit
go test -v ./src/...
go: finding github.com/stretchr/testify/assert latest
?       github.com/kenji-imi/sample-prj [no test files]
=== RUN   TestGetHello
--- PASS: TestGetHello (0.00s)
PASS
ok      github.com/kenji-imi/sample-prj/src/hello       0.024s
```