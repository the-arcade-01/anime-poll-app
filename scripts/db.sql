create database if not exists anime_poll_app;

use anime_poll_app;

create table if not exists anime_details (
    mal_id int PRIMARY KEY UNIQUE NOT NULL,
    title varchar(255) NOT NULL,
    image_link varchar(255) NOT NULL
);

create table if not exists anime_votes (
    mal_id int PRIMARY KEY UNIQUE NOT NULL,
    vote int NOT NULL DEFAULT 0,
    FOREIGN KEY (mal_id) REFERENCES anime_details (mal_id)
);