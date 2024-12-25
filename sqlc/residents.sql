-- name: GetResidentByKey :one
select * from residents where `key` = sqlc.arg(userKey);
