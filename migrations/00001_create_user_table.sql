-- +goose Up
CREATE TABLE user (
    id int NOT NULL AUTO_INCREMENT,
    name text,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE user;
