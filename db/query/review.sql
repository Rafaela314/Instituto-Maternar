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
  date,
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
  overall_rate,
  image
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
)
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE id = $1;

-- name: UpdateReview :exec
UPDATE reviews
  set title = $2,
  content = $3,
  classification = $4, 
  amount = $5, 
  insurance = $6, 
  place_id = $7, 
  place_rate = $8,
  doctor_id = $9,
  doctor_rate = $10,
  midwife_id= $11,
  midwife_rate = $12,
  doula_id = $13,
  doula_rate = $14,
  team = $15, 
  team_rate = $16,
  image = $17 
WHERE id = $1;