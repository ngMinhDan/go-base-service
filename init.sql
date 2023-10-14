-- I use PostgreSQL for this sample service : Users 
-- Below, There is a SQL code
-- You can run this or copy and paste into DB Console

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    profile_picture_url VARCHAR(255),
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE messages (
    id serial PRIMARY KEY,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- ALL PASSWORD ARE : 
INSERT INTO users (username, email, password, profile_picture_url, role, is_active)
VALUES ('admin', 'admin@example.com', '$2a$10$y58WHfkmnSvu6ebmdML2Q.26tZytcWPbQcKlFDbm0pjO7P2lkUFOu', 
'https://res.cloudinary.com/dtmebo99b/image/upload/c_scale,w_0.45,h_0.45/nmdan.com/media/banners/bat.jpeg', 'admin', true);

INSERT INTO users (username, email, password, profile_picture_url, role, is_active)
VALUES ('mod', 'mod@example.com', '$2a$10$y58WHfkmnSvu6ebmdML2Q.26tZytcWPbQcKlFDbm0pjO7P2lkUFOu', 
'https://res.cloudinary.com/dtmebo99b/image/upload/c_scale,w_0.45,h_0.45/nmdan.com/media/banners/bat.jpeg', 'mod', true);

INSERT INTO users (username, email, password, profile_picture_url, role, is_active)
VALUES ('user', 'user@example.com', '$2a$10$y58WHfkmnSvu6ebmdML2Q.26tZytcWPbQcKlFDbm0pjO7P2lkUFOu', 
'https://res.cloudinary.com/dtmebo99b/image/upload/c_scale,w_0.45,h_0.45/nmdan.com/media/banners/bat.jpeg', 'user', true);

INSERT INTO users (username, email, password, profile_picture_url, role, is_active)
VALUES ('root', 'root@example.com', '$2a$10$y58WHfkmnSvu6ebmdML2Q.26tZytcWPbQcKlFDbm0pjO7P2lkUFOu', 
'https://res.cloudinary.com/dtmebo99b/image/upload/c_scale,w_0.45,h_0.45/nmdan.com/media/banners/bat.jpeg', 'admin', true);

INSERT INTO messages (content, created_at)
VALUES ('Hello, world!', '2023-10-03 12:00:00');

INSERT INTO messages (content)
VALUES ('This is another message.');

INSERT INTO messages (content, created_at)
VALUES ('日本人が大好きですよ！', '2023-10-03 14:30:00');