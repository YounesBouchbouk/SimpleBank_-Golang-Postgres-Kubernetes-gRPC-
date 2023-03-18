-- name: accounts :one
INSERT INTO accounts (
  owner, balance , currency
) VALUES (
  $1, $2 , $3
)RETURNING *;


-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY name;

-- name: DeleteAccounts :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAuthor :exec
UPDATE accounts
  set balance = $2
WHERE id = $1;