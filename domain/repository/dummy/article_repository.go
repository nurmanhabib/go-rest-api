package dummy

import (
    "fmt"
    "github.com/nurmanhabib/go-rest-api/domain/model"
)

type articleRepository struct {
	articles []model.Article
}

func (repo *articleRepository) All() []model.Article {
	return repo.articles
}

func (repo *articleRepository) FindByID(ID int) (*model.Article, error) {
	for _, v := range repo.articles {
		if v.ID == ID {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("ResponseData Not Found")
}

func NewArticleRepository() model.ArticleRepository {
	articles := []model.Article {
		{ID: 101, Title: "Article 1", Author: "John Doe", Content: "Sample content article 1"},
		{ID: 102, Title: "Article 2", Author: "Tarjono", Content: "Sample content article 2"},
		{ID: 103, Title: "Article 3", Author: "Tara Basro", Content: "Sample content article 3"},
		{ID: 104, Title: "Article 4", Author: "Ahmad", Content: "Sample content article 4"},
	}

	return &articleRepository{articles: articles}
}
