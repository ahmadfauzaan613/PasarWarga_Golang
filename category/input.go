package category

type CategoryUserInput struct {
	CategoriName string `binding:"required"`
	CategoriSlug string `binding:"required"`
}

type GetCategoryDetail struct {
	ID int `uri:"id" binding:"required"`
}
