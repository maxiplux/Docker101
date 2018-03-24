--VERSION 0.5 BASE DE DATOS WHOPER
--AUTOR: CARLOS EDUARDO ARANGO <ArangoGutierrez>
--August 2017

--
-- Database: `whoperdb`
--

-- --------------------------------------------------------

--
-- Table admin
--
CREATE TABLE IF NOT EXISTS admin (
	admin_ID SERIAL PRIMARY KEY,
	username VARCHAR(11) NOT NULL,
	password VARCHAR(20) NOT NULL
);
---
--- Table users
---
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL,
    user_login VARCHAR(30) NOT NULL,
    pwd VARCHAR(200) NOT NULL,
    email VARCHAR(100) NOT NULL,
    gender VARCHAR(20) DEFAULT 'Human',
    location point NULL,
    salt VARCHAR(64) NOT NULL,
    status bit NOT NULL,
    image VARCHAR(200) DEFAULT 'https://upload.wikimedia.org/wikipedia/commons/6/67/User_Avatar.png',
	UNIQUE(user_login,email)
);
---
--- Table groups
---
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(240) NOT NULL,
    status bit NOT NULL,
    admin INTEGER REFERENCES users(id),
    imagen VARCHAR(200) DEFAULT 'https://upload.wikimedia.org/wikipedia/commons/b/b9/Group_font_awesome.svg',
	CONSTRAINT group_name_admin UNIQUE (admin,group_name)
);
---
---	Table posts
---
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    post_date date DEFAULT CURRENT_DATE,
    rating INTEGER NOT NULL,
    price	numeric NOT NULL DEFAULT 0,
    location point NULL,
		caption VARCHAR(240) NOT NULL,
    status bit NOT NULL,
    post_owner INTEGER REFERENCES users(id),
	CONSTRAINT users_post_owner UNIQUE (id,post_owner)
);
--
-- Table Post_Rating
--
CREATE TABLE IF NOT EXISTS post_rating (
	id SERIAL PRIMARY KEY,
	posts_id INTEGER REFERENCES posts(id) ON UPDATE CASCADE,
	user_id INTEGER REFERENCES users(id) ON UPDATE CASCADE,
	rating INTEGER NOT NULL,
	CONSTRAINT post_rating_user UNIQUE (user_id)
);
---
---	Table comment
---
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    comment_date date DEFAULT CURRENT_DATE,
    comment_txt VARCHAR(240) NOT NULL,
    status bit NOT NULL,
    comment_owner INTEGER REFERENCES users(id),
    post INTEGER REFERENCES posts(id),
	CONSTRAINT users_comment_owner UNIQUE (comment_owner)
);
---
---	Table images
---
CREATE TABLE IF NOT EXISTS images (
    id SERIAL PRIMARY KEY,
    url VARCHAR(200) NOT NULL,
    status bit NOT NULL
);
--
-- Table Following
--
CREATE TABLE IF NOT EXISTS following (
	user_id INTEGER REFERENCES users(id) ON UPDATE CASCADE,
	follows_user_id INTEGER REFERENCES users(id) ON UPDATE CASCADE,
	follow_date date DEFAULT CURRENT_DATE,
	CONSTRAINT following_tuple UNIQUE (user_id,follows_user_id)
);
---
---	Table users in group
---
CREATE TABLE IF NOT EXISTS users_groups (
    user_id INTEGER REFERENCES users(id) ON UPDATE CASCADE,
    group_id INTEGER REFERENCES groups(id) ON UPDATE CASCADE,
		CONSTRAINT users_groups_tuple UNIQUE (user_id,group_id)
);
---
---	Table post on group
---
CREATE TABLE IF NOT EXISTS posts_groups (
    post_id INTEGER REFERENCES posts(id) ON UPDATE CASCADE,
    group_id INTEGER REFERENCES groups(id) ON UPDATE CASCADE
);
---
--- Table images of post
---
CREATE TABLE IF NOT EXISTS images_post  (
    posts_id  INTEGER REFERENCES posts(id) ON UPDATE CASCADE,
    image_id INTEGER REFERENCES images(id) ON UPDATE CASCADE
		CONSTRAINT images_post_tuple UNIQUE (posts_id,image_id)
);
-- --------------------------------------------------------
--
-- Dumping data for table admin
INSERT INTO admin (username, password)
	VALUES	('Whoperadmin', 'montanoarango');
--EOF
