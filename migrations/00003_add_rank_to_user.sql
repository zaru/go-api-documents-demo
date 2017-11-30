-- +goose Up
ALTER TABLE user ADD rank int not null;

-- +goose Down
ALTER TABLE user DROP rank;
