package models

type Product struct {
	ID                int               `json:"id"`
	CatID             int               `json:"category_id,omitempty"`
	Name              string            `json:"name"`
	Price             int               `json:"price"`
	ProductAttributes ProductAttributes `json:"product_attributes"`
	Category          Category          `json:"category,omitempty"`
}

type ProductAttributes struct {
	ID                   int    `json:"id,omitempty"`
	ProductID            int    `json:"product_id,omitempty"`
	FirstAttributeName   string `json:"first_attribute_name"`
	FirstAttributeValue  string `json:"first_attribute_value"`
	SecondAttributeName  string `json:"second_attribute_name"`
	SecondAttributeValue string `json:"second_attribute_value"`
}
