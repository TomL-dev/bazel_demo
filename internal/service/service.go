package service

type service struct {
}

func NewBookService() *service {
	return &service{}
}

type Book struct {
	Title string
}

func (s *service) GetBook(title string) *Book {
	return &Book{
		Title: title,
	}
}
