SELECT id, name, parent_id FROM categories ORDER BY id ASC;


SELECT id, name, parent_id FROM categories WHERE parent_id = 0;



CREATE table categories(
    id serial primary key,
    name varchar not null,
    parent_id interval
);

INSERT INTO categories(name, parent_id)
VALUES
    ('Комплектующие для ПК', 0),
    ('Основные комплектующие для ПК', 1),
    ('Процессоры', 2);


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
                                    attribute_name VARCHAR(255) NOT NULL,
                                    attribute_value VARCHAR(255) NOT NULL,
                                    FOREIGN KEY (product_id) REFERENCES products (id)
);

 drop table products;





CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255),
                          price NUMERIC(10,2),
                          category_id INTEGER REFERENCES categories(id)
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255),
                          price NUMERIC(10,2),
                          category_id INTEGER REFERENCES categories(id)
);

CREATE TABLE attribute_names (
                                 id SERIAL PRIMARY KEY,
                                 name VARCHAR(255)
);

CREATE TABLE product_attributes (
                                    id SERIAL PRIMARY KEY,
                                    product_id INTEGER REFERENCES products(id),
                                    attribute_name_id INTEGER REFERENCES attribute_names(id),
                                    attribute_value VARCHAR(255)
);


-- Вставляем запись в таблицу products
INSERT INTO products (name, price, category_id)
VALUES ('Процессор Intel Core i9-10940X OEM', 999.99, 1)
RETURNING id;

-- Вставляем запись в таблицу attributes
INSERT INTO attribute_names (name)
VALUES ('Характеристики процессора')
RETURNING id;

-- Вставляем запись в таблицу product_attributes
INSERT INTO product_attributes (product_id, attribute_name_id, first_value, second_value)
VALUES (
           (SELECT id FROM products WHERE name = 'Процессор Intel Core i9-10940X OEM'),
           (SELECT id FROM attribute_names WHERE name = 'Характеристики процессора'),
           'Общее количество ядер: 14',
           'Базовая частота процессора ГГц: 3.3'
       );


INSERT INTO product_attributes (product_id, attribute_name_id, attribute_value)
VALUES (
           1,
           1,
           '14:Базовая частота процессора ГГц:3.3'
       );






INSERT INTO attributes (name) VALUES ('Общее количество ядер');
INSERT INTO attributes (name) VALUES ('Базовая частота процессора ГГц');
INSERT INTO attributes (name) VALUES ('Диагональ экрана');
INSERT INTO attributes (name) VALUES ('Технология изготовления матрицы');


INSERT INTO products (name, price, category_id) VALUES ('Процессор Intel Core i9-10940X OEM', 75000.00, 1);

INSERT INTO product_attributes (product_id, attribute_id, value) VALUES (1, 1, '14');
INSERT INTO product_attributes (product_id, attribute_id, value) VALUES (1, 2, '3.3');


SELECT p.name as product_name,
       string_agg(pa.attribute_name || ': ' || pa.attribute_value, ', ') as attributes,
       p.price
FROM products p
         JOIN product_categories pc ON p.category_id = pc.id
         JOIN product_attributes pa ON p.id = pa.product_id
WHERE pc.name = 'Процессоры'
GROUP BY p.name, p.price;





CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255),
                          price NUMERIC(10,2),
                          category_id INTEGER REFERENCES categories(id)
);

CREATE TABLE attributes (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255)
);

CREATE TABLE product_attributes (
                                    id SERIAL PRIMARY KEY,
                                    product_id INTEGER REFERENCES products(id),
                                    attribute_id INTEGER REFERENCES attributes(id),
                                    first_value VARCHAR(255),
                                    second_value VARCHAR(255)
);

INSERT INTO products (name, price, category_id)
VALUES ('Процессор Intel Core i9-10940X OEM', 60000, 1);

INSERT INTO attributes (name)
VALUES ('Количество ядер'), ('Базовая частота процессора ГГц');

INSERT INTO product_attributes (product_id, attribute_id, first_value, second_value)
VALUES (
           1,
           1,
           '14',
           '3.3'
       ),
       (
           1,
           2,
           NULL,
           NULL
       );


CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255),
                          price NUMERIC(10,2),
                          category_id INTEGER REFERENCES categories(id)
);

CREATE TABLE attributes (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255)
);

CREATE TABLE product_attributes (
                                    id SERIAL PRIMARY KEY,
                                    product_id INTEGER REFERENCES products(id),
                                    attribute_id INTEGER REFERENCES attributes(id),
                                    first_value VARCHAR(255),
                                    second_value VARCHAR(255)
);

INSERT INTO products (name, price, category_id)
VALUES ('Процессор Intel Core i9-10940X OEM', 60000, 1);

INSERT INTO attributes (name)
VALUES ('Количество ядер'), ('Базовая частота процессора ГГц');

INSERT INTO product_attributes (product_id, attribute_id, first_value, second_value)
VALUES (
           1,
           1,
           '14',
           '3.3'
       ),
       (
           1,
           2,
           NULL,
           NULL
       );


SELECT p.name AS product_name,
       p.price AS product_price,
       c.name AS category_name,
       a1.name AS attribute1_name,
       pa1.first_value AS attribute1_value,
       a2.name AS attribute2_name,
       pa2.second_value AS attribute2_value
FROM products p
         JOIN categories c ON p.category_id = c.id
         LEFT JOIN product_attributes pa1 ON p.id = pa1.product_id AND pa1.attribute_id = 1
         LEFT JOIN attributes a1 ON pa1.attribute_id = a1.id
         LEFT JOIN product_attributes pa2 ON p.id = pa2.product_id AND pa2.attribute_id = 2
         LEFT JOIN attributes a2 ON pa2.attribute_id = a2.id;



