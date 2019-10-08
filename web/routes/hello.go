package routes

import (
	"errors"
	"github.com/kataras/iris/hero"
)

var badNameErr = errors.New("bad name")

var badName = hero.Response{Err: badNameErr, Code: 400}

var helloView = hero.View{
	Name: "hello/index.html"
	Data: map[string]interface{} {
		"Title" :		"hello world"
		"Message":		"welcome to my world"
	}
}


func Hello() hero.Result {
	return helloView
}

func HelloName(name string) hero.Result {
	if name == "x" {
		return badName
	}

	return hero.View {
		Name: "hello/name.html"
		Data: name
	}
}
