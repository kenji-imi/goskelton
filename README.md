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