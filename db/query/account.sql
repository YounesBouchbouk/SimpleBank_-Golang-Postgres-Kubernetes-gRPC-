-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance , currency
) VALUES (
  $1, $2 , $3
)RETURNING *;


-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY owner offset sqlc.arg(offsetnb) limit sqlc.arg(limitnb);

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;


