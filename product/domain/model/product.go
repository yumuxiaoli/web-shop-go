package model

type Product struct {
	ID                 int64          `gorm:"primary_key;not_null;auto_increment"`
	ProductName        string         `json:"product_name"`
	ProductSku         string         `gorm:"unique_index;not_null" json:"product_sku"`
	ProductPrice       float64        `json:"product_description"`
	ProductDescription string         `json:"product_description"`
	ProductImage       []ProductImage `gorm:"ForeignKey:ImageProduct/id" json:"product_image"`
	ProductSize        []ProductImage `gorm:"ForeignKey:SizeProductID" json:"product_seo"`
	ProductSeo         ProductSeo     `gorm:"ForeignKey:SeoProductID" json:"product_seo"`
}
