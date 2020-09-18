-- +migrate Up
CREATE TABLE data (
    video_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    picture VARCHAR(255) NOT NULL,
    channel_title VARCHAR(255) NOT NULL,
    publish_time DATETIME NOT NULL
);