-- name: GetDoctor :one
SELECT * FROM doctors
WHERE id = $1 LIMIT 1;

-- name: ListDoctors :many
SELECT * FROM doctors
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateDoctor :one
INSERT INTO doctors (
  name,
  crm,
  average_rate
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteDoctor :exec
DELETE FROM doctors
WHERE id = $1;

-- name: UpdateDoctor :exec
UPDATE doctors
  set name = $2,
  crm = $3,
  average_rate = $4
WHERE id = $1;

-- name: ListDoctorReviews :many
SELECT * FROM reviews
WHERE doctor_id = $1
ORDER BY id;

-- name: CountDoctorReviews :one
SELECT COUNT(id)
FROM reviews
WHERE doctor_id = $1;