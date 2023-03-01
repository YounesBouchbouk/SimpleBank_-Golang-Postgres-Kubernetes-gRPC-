-- name: entries :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)RETURNING *;


-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries;


-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = $1;

