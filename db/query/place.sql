-- name: GetPlace :one
SELECT * FROM places
WHERE id = $1 LIMIT 1;

-- name: ListPlaces :many
SELECT * FROM places
ORDER BY id;

-- name: CreatePlace :one
INSERT INTO places (
  name,
  address,
  city,
  state,
  country,
  classification, 
  average_rate
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: DeletePlace :exec
DELETE FROM places
WHERE id = $1;

-- name: UpdatePlace :exec
UPDATE places
  set name = $2,
  address = $3,
  city = $4,
  state = $5,
  country = $6,
  classification = $7
WHERE id = $1;

-- name: CountPlaceReviews :one
SELECT COUNT(id)
FROM reviews
WHERE place_id = $1;