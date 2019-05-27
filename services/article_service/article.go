package article_service

import "ohflutter-server/models"

type Article struct {
	UserID int
	Title  string
}

func (u *Article) Add() error {

	article := map[string]interface{}{
		"user_id": u.UserID,
		"title":   u.Title,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}
