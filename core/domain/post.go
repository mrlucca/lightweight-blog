package domain

import (
	"errors"
	"strings"
	"time"
)

type PostStep byte

const (
	PostCreatedStep          = PostStep('C')
	PostAvailableForReadStep = PostStep('R')
)

var (
	PostAlreadyExistsError = errors.New("post already exists")
)

type Post struct {
	Title       string
	Description string
	Slug        string
	Creator     Author
	Tags        []Tag
	Step        PostStep
	CreatedAt   time.Time
	UpdateAt    time.Time
}

type IPostRepository interface {
	Save(post Post) error
	GetFromSlug(slug string) error
	UpdateDescriptionFromSlug(slug, description string) error
	DeleteFromSlug(slug string) error
	AddTags(tags []Tag)
	AddTag(tag Tag)
}

func NewPostWithAutoCompleteSlug(title, description string, tags []Tag, creator Author) *Post {
	postCreatedTime := time.Now()
	titleWithoutSpaces := strings.TrimSpace(title)

	return &Post{
		Title:       titleWithoutSpaces,
		Slug:        getSlugFromTitle(titleWithoutSpaces),
		Description: description,
		Tags:        tags,
		Creator:     creator,
		CreatedAt:   postCreatedTime,
		UpdateAt:    postCreatedTime,
		Step:        PostCreatedStep,
	}
}

func getSlugFromTitle(title string) string {
	titleInLowerCaseFormat := strings.ToLower(title)
	slug := strings.Replace(titleInLowerCaseFormat, " ", "-", -1)
	return slug
}
