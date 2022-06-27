package main

import (
	"fmt"

	"github.com/TomL-dev/bazel_demo/internal/repository"
	"github.com/TomL-dev/bazel_demo/internal/service"
)

func main() {
	r := repository.New()
	service := service.NewBookService(r)

	book := service.GetBook("The Water Dancer")
	fmt.Println(book.Title)
}
