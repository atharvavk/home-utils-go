-- name: GetGeyserStatus :one
select * from geyser_status g join residents r on g.action_by = r.`key`;

-- name: TurnOnGeyser :execrows
update geyser_status set is_on = true, action_by = sqlc.arg(userKey) where is_on = false;

-- name: TurnOffGeyser :execrows
update geyser_status set is_on = false where action_by = sqlc.arg(userKey) and is_on = true;
