INSERT INTO
    languages(code, status)
VALUES
    ('en-EN', 'available')
ON CONFLICT DO
    NOTHING;

INSERT INTO
    users(language_id, username, password, name, last_name, birthday, theme)
VALUES
    (1, 'coffemanfp', '1234', 'Franklin', 'Peñaranda', NOW(), 'light')
ON CONFLICT DO
    NOTHING;