create database if not exists anime_poll_app;

use anime_poll_app;

create table anime_details (
    mal_id int PRIMARY KEY UNIQUE NOT NULL,
    title varchar(255) NOT NULL,
    image_link varchar(255) NOT NULL
);