-- +goose Up
-- +goose StatementBegin
ALTER TABLE `events` ADD INDEX `message_owner_event_type_index` (`message_owner`, `event_type`);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX message_owner_event_type_index on events;
-- +goose StatementEnd