package main

import (
	"fmt"

	"github.com/toml-dev/bazel_demo/internal/service"
)

func main() {

	service := service.NewBookService()

	book := service.GetBook("The Water Dancer")
	fmt.Println(book.Title)
}
