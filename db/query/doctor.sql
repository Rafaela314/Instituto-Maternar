-- name: GetDoctor :one
SELECT * FROM doctors
WHERE id = $1 LIMIT 1;

-- name: ListDoctors :many
SELECT * FROM doctors
ORDER BY id;

-- name: CreateDoctor :one
INSERT INTO doctors (
  name,
  crm
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteDoctor :exec
DELETE FROM doctors
WHERE id = $1;