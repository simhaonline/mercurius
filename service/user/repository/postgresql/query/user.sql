-- name: ListUsers :many
SELECT * FROM "user";

-- name: CreateUser :exec
INSERT INTO "user" (id,firstname,lastname,login) VALUES ($1,$2,$3,$4);

-- name: DeleteUser :exec
DELETE FROM "user" WHERE id = $1;