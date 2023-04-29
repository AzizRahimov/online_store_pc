package repository

import (
	"database/sql"
	"fmt"
	models "github.com/AzizRahimov/online_store_pc/model"
)

type ProductImp struct {
	db *sql.DB
}

func NewProductImp(db *sql.DB) *ProductImp {
	return &ProductImp{db: db}
}

func (p *ProductImp) AddProduct(product models.Product) error {

	fmt.Println("product", product)

	addProduct := fmt.Sprintf("INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3) RETURNING id")
	addProductAtr := fmt.Sprintf(`INSERT INTO product_attributes (product_id, first_attribute_name, first_attribute_value, second_attribute_name, second_attribute_value)
VALUES($1, $2, $3, $4, $5)`)

	// добавление товара
	err := p.db.QueryRow(addProduct, product.Name, product.Price, product.CatID).Scan(&product.ID)

	if err != nil {

		return err
	}

	// добавление описание к продукту
	_, err = p.db.Exec(addProductAtr,
		product.ID,
		product.ProductAttributes.FirstAttributeName,
		product.ProductAttributes.FirstAttributeValue,
		product.ProductAttributes.SecondAttributeName,
		product.ProductAttributes.SecondAttributeValue)
	if err != nil {

		return err
	}

	return nil

}

func (p *ProductImp) DeleteProduct(id int) error {
	deleteProductAtr := fmt.Sprintf("DELETE FROM product_attributes WHERE product_id = $1")
	deleteProduct := fmt.Sprintf("DELETE FROM products WHERE id = $1")

	_, err := p.db.Exec(deleteProductAtr, id)
	if err != nil {
		fmt.Println("ошибка при удалении атрибутов продукта:", err)
		return err
	}

	_, err = p.db.Exec(deleteProduct, id)
	if err != nil {
		fmt.Println("ошибка при удалении продукта:", err)
		return err
	}

	return nil
}

func (p *ProductImp) UpdateProduct(product models.Product) error {
	updateProduct := fmt.Sprintf("UPDATE products SET name = $1, price = $2, category_id = $3 WHERE id = $4")
	updateProductAtr := fmt.Sprintf(`UPDATE product_attributes SET first_attribute_name = $1, first_attribute_value = $2, second_attribute_name = $3, second_attribute_value = $4 WHERE product_id = $5`)

	// обновление информации о продукте
	_, err := p.db.Exec(updateProduct, product.Name, product.Price, product.CatID, product.ID)
	if err != nil {
		fmt.Println("ошибка при обновлении продукта:", err)
		return err
	}

	// обновление атрибутов продукта
	_, err = p.db.Exec(updateProductAtr,
		product.ProductAttributes.FirstAttributeName,
		product.ProductAttributes.FirstAttributeValue,
		product.ProductAttributes.SecondAttributeName,
		product.ProductAttributes.SecondAttributeValue,
		product.ID)
	if err != nil {
		fmt.Println("ошибка при обновлении атрибутов продукта:", err)
		return err
	}

	return nil
}

func (p *ProductImp) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	selectAllProducts := `
        SELECT p.id, c.name, p.name, pa.first_attribute_name, pa.first_attribute_value, pa.second_attribute_name, pa.second_attribute_value, p.price
        FROM products p
        JOIN categories c ON c.id = p.category_id
        JOIN product_attributes pa ON p.id = pa.product_id;
    `
	rows, err := p.db.Query(selectAllProducts)
	if err != nil {
		fmt.Println("ошибка при выборке всех продуктов:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Category.Title,
			&product.Name,
			&product.ProductAttributes.FirstAttributeName,
			&product.ProductAttributes.FirstAttributeValue,
			&product.ProductAttributes.SecondAttributeName,
			&product.ProductAttributes.SecondAttributeValue,
			&product.Price,
		)
		if err != nil {
			fmt.Println("ошибка при сканировании строк из базы данных:", err)
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
