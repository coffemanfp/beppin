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
    users(language_id, username, password, name, last_name, birthday, theme)
VALUES
    (1, 'coffemanfp', '1234', 'Franklin', 'Peñaranda', NOW(), 'light')
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