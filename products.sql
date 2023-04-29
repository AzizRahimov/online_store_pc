CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          category_id INTEGER NOT NULL,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          price NUMERIC(10,2) NOT NULL,
                          FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TABLE product_attributes (
                                    id SERIAL PRIMARY KEY,
                                    product_id INTEGER NOT NULL,
                                    first_attribute_name VARCHAR(255) NOT NULL,
                                    first_attribute_value VARCHAR(255) NOT NULL,
                                    second_attribute_name VARCHAR(255) NOT NULL,
                                    second_attribute_value VARCHAR(255) NOT NULL,
                                    FOREIGN KEY (product_id) REFERENCES products (id)
);

INSERT INTO products (name, price, category_id)
VALUES ('Процессор Intel Core i9-10940X OEM', 999.99, 3)
RETURNING id;
-- вот этот пункт, обязателен


-- Вставляем запись в таблицу product_attributes
INSERT INTO product_attributes (product_id, first_attribute_name, first_attribute_value, second_attribute_name, second_attribute_value )
VALUES (1,  'Общее количество ядер', '14', 'Базовая частота процессора ГГц', '3.3' );



SELECT p.id,c.name, p.name, pa.first_attribute_name, pa.first_attribute_value, pa.second_attribute_name, pa.second_attribute_value, p.price FROM  products p
Join categories c on c.id = p.category_id
JOIN product_attributes pa on p.id = pa.product_id;