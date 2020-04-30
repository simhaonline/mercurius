/* name: ListUsers :many */
SELECT * FROM user;

/* name: CreateUser :exec */
INSERT INTO user (id,firstname,lastname,login) VALUES (?,?,?,?);

/* name: DeleteUser :exec */
DELETE FROM user WHERE id = ?;