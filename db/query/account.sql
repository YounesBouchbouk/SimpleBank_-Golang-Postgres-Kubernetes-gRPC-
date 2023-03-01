-- name: account :one
INSERT INTO account (
  owner, balance , currency
) VALUES (
  $1, $2 , $3
)RETURNING *;


-- name: GetAccounts :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY name;

-- name: DeleteAccounts :exec
DELETE FROM account
WHERE id = $1;

-- name: UpdateAuthor :exec
UPDATE account
  set balance = $2
WHERE id = $1;