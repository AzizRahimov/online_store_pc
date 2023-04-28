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