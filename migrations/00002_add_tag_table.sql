-- +goose Up
CREATE TABLE tag (
    id int NOT NULL AUTO_INCREMENT,
    user_id int,
    name text,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE tag;
