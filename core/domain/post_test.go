package domain_test

import (
	"lightweight-blog/core/domain"
	"testing"
)

var post = domain.NewPostWithAutoCompleteSlug(
	"Test novo post  ",
	"oi mais uma description",
	nil,
	domain.Author{
		Name: "test",
		Bio:  "test",
	},
)

func TestNewPostWithAutoCompleteIfTitleIsCorrect(t *testing.T) {
	if post.Title != "Test novo post" {
		t.Errorf("invalid title field %v", post.Title)
	}
}

func TestNewPostWithAutoCompleteIfSlugIsCorrect(t *testing.T) {
	if post.Slug != "test-novo-post" {
		t.Errorf("ivalid slug field %v", post.Slug)
	}
}
