package article

type ArticleUserInput struct {
	Title      string `binding:"required"`
	Slug       string `binding:"required"`
	CategoryId int    `binding:"required"`
	Content    string `binding:"required"`
}

type GetArticleDetail struct {
	ID int `uri:"id" binding:"required"`
}
