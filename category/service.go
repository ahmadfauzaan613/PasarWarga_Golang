package category

type Service interface {
	RegisterCategory(input CategoryUserInput) (Categori, error)
	GetAllCategory() ([]Categori, error)
	GetCategoryByID(input GetCategoryDetail) (Categori, error)
	Update(inputID GetCategoryDetail, inputData CategoryUserInput) (Categori, error)
	Delete(inputID GetCategoryDetail) (Categori, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterCategory(input CategoryUserInput) (Categori, error) {
	category := Categori{}
	category.CategoryName = input.CategoriName
	category.CategorySlug = input.CategoriSlug

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) GetAllCategory() ([]Categori, error) {
	categoris, err := s.repository.FindAll()
	if err != nil {
		return categoris, err
	}

	return categoris, nil

}

func (s *service) GetCategoryByID(input GetCategoryDetail) (Categori, error) {
	categoris, err := s.repository.FindByID(input.ID)
	if err != nil {
		return categoris, err
	}

	return categoris, nil
}

func (s *service) Update(inputID GetCategoryDetail, inputData CategoryUserInput) (Categori, error) {
	category, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return category, err
	}
	category.CategoryName = inputData.CategoriName
	category.CategorySlug = inputData.CategoriSlug

	updateNewCategori, err := s.repository.UpdateCategory(category)
	if err != nil {
		return updateNewCategori, err
	}
	return updateNewCategori, nil
}

func (s *service) Delete(inputID GetCategoryDetail) (Categori, error) {
	category, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return category, err
	}
	DeletedCategori, err := s.repository.CategoryDelete(category)
	if err != nil {
		return DeletedCategori, err
	}
	return DeletedCategori, nil
}
