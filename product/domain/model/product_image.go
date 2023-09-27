package model

type ProductImage struct {
	ID         int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ImagesName string `json:"image_name"`
	ImagesCode string `gorm:"unique_index;not_null" json:"image_code"`
	ImagesUrl  string `json:"image_url"`
	ImagesProductId int64 `json:"image_url"`
}
