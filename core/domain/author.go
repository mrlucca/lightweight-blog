package domain

import "errors"

type Author struct {
	Name string
	Id   string
	Bio  string
}

type IAuthorRepository interface {
	GetFromName(name string) (Author, error)
}

var (
	AuthorNotExists     = errors.New("author not exists")
	AuthorAlreadyExists = errors.New("author already exists")
)
