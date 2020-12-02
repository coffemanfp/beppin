INSERT INTO
    languages(code, status)
VALUES
    ('en-EN', 'available')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/8987987237462387642photo_2019-10-11_18-27-08.jpg')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/16055147830696068331.webp')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/3403453453069506832s-1229.webp')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/4545349829387492837q.webp')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/5334538237987129837photo_2019-10-11_18-27-08.jpg')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/0982087928374598273photo_2019-10-11_18-27-08.jpg')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/2983749287349798938photo_2019-10-11_18-27-08.jpg')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files(path)
VALUES
    ('assets/7898789876787654234photo_2019-10-11_18-27-08.jpg')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    languages(code, status)
VALUES
    ('es-ES', 'in-progress')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(language, username, password, email, avatar_id, name, last_name, birthday, theme, currency)
VALUES
    ('es-ES', 'coffemanfp', '1234', 'coffemanfp@gmail.com', 1, 'Franklin', 'Peñaranda', NOW(), 'dark', 'EUR')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(username, password, email, name, last_name)
VALUES
    ('glendysanez', '1234', 'example@gmail.com', 'Glendys', 'Anez')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    products(user_id, name, description, price)
VALUES
    (1, 'Cool shirt', 'Awesome shirt', 887654308.08)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    products(user_id, name, description, price)
VALUES
    (1, 'Super cool jacket', 'Yeah', 76)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    products(user_id, name, description, price)
VALUES
    (2, 'Awesome lamp', 'A awesome.... lamp?', 53.02)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    offers(product_id, type, value, description, expirated_at)
VALUES
    (1, '%', '%50', 'Just for now!!', NOW())
ON CONFLICT DO
    NOTHING;

INSERT INTO
    offers(product_id, type, value, description)
VALUES
    (2, 'x', '2x1', 'Unlimited!!')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    categories(name, description)
VALUES
    ('Home', 'Common home products')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    categories(name, description)
VALUES
    ('Clothes', 'Build and Break your style!')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    categories(name, description)
VALUES
    ('Illuminosity', 'Illuminate your site')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    categories(name)
VALUES
    ('Audio')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    categories(name, description)
VALUES
    ('Devices', 'Computers... SmartPhones... the whole future')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    product_categories(category_id, product_id)
VALUES
    (2, 1)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    product_categories(category_id, product_id)
VALUES
    (2, 2)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    product_categories(category_id, product_id)
VALUES
    (3, 3)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files_products(file_id, product_id)
VALUES
    (2, 2)
ON CONFLICT DO
    NOTHING;
    
INSERT INTO
    files_products(file_id, product_id)
VALUES
    (6, 2)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files_products(file_id, product_id)
VALUES
    (3, 1)
ON CONFLICT DO
    NOTHING;

INSERT INTO
    files_products(file_id, product_id)
VALUES
    (4, 3)
ON CONFLICT DO
    NOTHING;
