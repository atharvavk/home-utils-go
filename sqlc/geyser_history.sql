-- name: InsertGeyserHistory :exec
insert into geyser_history (`action`, resident_key) values (sqlc.arg(actionValue), sqlc.arg(userKey));

-- name: GetGeyserHistoryPaginated :many
select gh.* , r.display_name as display_name from geyser_history gh join residents r on r.`key` = gh.resident_key order by created_at desc limit? offset?;

-- name: GetGeyserHistoryCount :one
select count(*) from geyser_history;
