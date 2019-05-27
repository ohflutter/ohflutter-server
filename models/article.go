package models

type Article struct {
	Model

	User   User 	 `gorm:"ForeignKey:UserID;association_foreignkey:ID"`
	UserID int

	Title     string `json:"title"`
}

func AddArticle(data map[string]interface{}) error {
	article := Article{
		UserID: data["user_id"].(int),
		Title:  data["title"].(string),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}
