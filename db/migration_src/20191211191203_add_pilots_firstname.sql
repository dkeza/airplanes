-- +goose Up
-- +goose StatementBegin
ALTER TABLE pilots ADD COLUMN firstname text DEFAULT '';
UPDATE pilots SET firstname = 'John' WHERE id = 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE pilots DROP COLUMN firstname;
-- +goose StatementEnd
