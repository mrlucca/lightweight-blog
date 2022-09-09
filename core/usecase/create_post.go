package usecase

import "lightweight-blog/core/domain"

type createPostUseCase struct {
	postRepository   domain.IPostRepository
	authorRepository domain.IAuthorRepository
}

type NewPostInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatorName string `json:"creatorName"`
}

func NewCreatePostUseCase(pr domain.IPostRepository, ar domain.IAuthorRepository) *createPostUseCase {
	return &createPostUseCase{
		postRepository:   pr,
		authorRepository: ar,
	}
}

func (u createPostUseCase) Execute(input NewPostInput) error {
	creator, err := u.authorRepository.GetFromName(input.CreatorName)

	if err != nil {
		return err
	}
	post := domain.NewPostWithAutoCompleteSlug(
		input.Title,
		input.Description,
		nil,
		creator,
	)
	err = u.postRepository.Save(*post)
	if err != nil {
		return err
	}
	return nil
}
