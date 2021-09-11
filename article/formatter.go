package article

type ArticleFormatter struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	CategoryId int    `json:"categoryid"`
	Content    string `json:"content"`
}

func FormatArticle(article Article) ArticleFormatter {
	formatter := ArticleFormatter{
		ID:         article.ID,
		Title:      article.Title,
		Slug:       article.Slug,
		CategoryId: article.CategoryId,
		Content:    article.Content,
	}
	return formatter
}

type ArticleDetailFormatter struct {
	ID         int               `json:"id"`
	Title      string            `json:"title"`
	Slug       string            `json:"slug"`
	CategoryId int               `json:"categoryid"`
	Content    string            `json:"content"`
	Categori   CategoriDetailAll `json:"categori"`
}

type CategoriDetailAll struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func AllFormatarticle(article Article) ArticleDetailFormatter {
	formatallArticle := ArticleDetailFormatter{}
	formatallArticle.ID = article.ID
	formatallArticle.Title = article.Title
	formatallArticle.Slug = article.Slug
	formatallArticle.CategoryId = article.CategoryId
	formatallArticle.Content = article.Content

	category := article.Categori
	categoriDetail := CategoriDetailAll{}
	// -----------------------
	categoriDetail.Name = category.CategoryName
	categoriDetail.Slug = category.CategorySlug

	// -------------------
	formatallArticle.Categori = categoriDetail

	return formatallArticle
}
