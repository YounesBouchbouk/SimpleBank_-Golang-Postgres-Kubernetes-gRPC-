-- name: Transfers :one
INSERT INTO transfers (
  from_account_id,to_account_id, amount
) VALUES (
  $1, $2 , $3
)RETURNING *;


-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers;


-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1;
