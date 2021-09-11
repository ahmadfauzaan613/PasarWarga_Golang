package article

type Service interface {
	RegisterArticle(input ArticleUserInput) (Article, error)
	GetAllArticle() ([]Article, error)
	GetArticleByID(input GetArticleDetail) (Article, error)
	Update(inputID GetArticleDetail, inputData ArticleUserInput) (Article, error)
	Delete(inputID GetArticleDetail) (Article, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterArticle(input ArticleUserInput) (Article, error) {
	article := Article{}
	article.Title = input.Title
	article.Slug = input.Slug
	article.CategoryId = input.CategoryId
	article.Content = input.Content

	newArticle, err := s.repository.Save(article)
	if err != nil {
		return newArticle, err
	}

	return newArticle, nil
}

func (s *service) GetAllArticle() ([]Article, error) {
	articles, err := s.repository.FindAll()
	if err != nil {
		return articles, err
	}

	return articles, nil

}

func (s *service) Update(inputID GetArticleDetail, inputData ArticleUserInput) (Article, error) {
	articles, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return articles, err
	}
	articles.Title = inputData.Title
	articles.Slug = inputData.Slug
	articles.Content = inputData.Slug

	updateNewArticle, err := s.repository.UpdateArticle(articles)
	if err != nil {
		return updateNewArticle, err
	}
	return updateNewArticle, nil
}

func (s *service) Delete(inputID GetArticleDetail) (Article, error) {
	articles, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return articles, err
	}
	DeletedArticle, err := s.repository.ArticleDelete(articles)
	if err != nil {
		return DeletedArticle, err
	}
	return DeletedArticle, nil
}

func (s *service) GetArticleByID(input GetArticleDetail) (Article, error) {
	articles, err := s.repository.FindByID(input.ID)
	if err != nil {
		return articles, err
	}
	return articles, nil
}
