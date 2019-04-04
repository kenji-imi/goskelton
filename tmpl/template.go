package tmpl

// [projeect_root]/main.go
const MainGoTmpl = `package main

import "github.com/{{ .User }}/{{ .Project }}/src/hello"

func main() {
	hello.SayHello()
}
`

// [projeect_root]/src/hello/hello.go
const HelloGoTmpl = `package hello

import "fmt"

func SayHello() {
	fmt.Println("Hello!")
}

func GetHello() string {
	return "Hello"
}
`

// [projeect_root]/src/hello/hello_test.go
const HelloTestGoTmpl = `package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	hello := GetHello()
	assert.Equal(t, "Hello", hello)
}
`
