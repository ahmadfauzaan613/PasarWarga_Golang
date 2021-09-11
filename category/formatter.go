package category

type CategoryFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func FormatCategory(category Categori) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:   category.ID,
		Name: category.CategoryName,
		Slug: category.CategorySlug,
	}
	return formatter
}
