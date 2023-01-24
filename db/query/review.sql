-- name: GetReview :one
SELECT * FROM reviews
WHERE id = $1 LIMIT 1;

-- name: ListReviews :many
SELECT * FROM reviews
ORDER BY id;

-- name: CreateReview :one
INSERT INTO reviews (
  user_id,
  title,
  content,
  classification,
  amount,
  insurance,
  place_id,
  place_rate,
  doctor_id,
  doctor_rate,
  midwife_id,
  midwife_rate,
  doula_id,
  doula_rate,
  team, 
  team_rate,
  image
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
)
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE id = $1;
