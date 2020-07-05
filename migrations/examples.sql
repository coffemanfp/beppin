INSERT INTO
    languages(code, status)
VALUES
    ('en-EN', 'available')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    languages(code, status)
VALUES
    ('es-ES', 'in-progress')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(language, username, password, email, name, last_name, birthday, theme)
VALUES
    ('es-ES', 'coffemanfp', '1234', 'coffemanfp@gmail.com', 'Franklin', 'Peñaranda', NOW(), 'light')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    products(user_id, name, description, categories)
VALUES
    (1, 'Cool t-shirt', 'Awesome t-shirt', ARRAY['clothes', 't-shirts']);

INSERT INTO
    products(user_id, name, description, categories)
VALUES
    (1, 'Super cool pants', 'Yeah', ARRAY['clothes', 'pants']);

INSERT INTO
    products(user_id, name, description, categories)
VALUES
    (1, 'Awesome lamp', 'A awesome.... lamp?', ARRAY['luminosity', 'house', 'office']);