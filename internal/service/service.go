package service

import "github.com/TomL-dev/bazel_demo/internal/repository"

type service struct {
	repository *repository.Instance
}

func NewBookService(r *repository.Instance) *service {
	return &service{
		repository: r,
	}
}

type Book struct {
	Title string
}

func (s *service) GetBook(title string) *Book {
	return &Book{
		Title: title,
	}
}
