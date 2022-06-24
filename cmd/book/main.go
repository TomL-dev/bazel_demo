package main

import (
	"fmt"

	"github.com/TomL-dev/bazel_demo/internal/service"
)

func main() {

	service := service.NewBookService()

	book := service.GetBook("The Water Dancer")
	fmt.Println(book.Title)
}
