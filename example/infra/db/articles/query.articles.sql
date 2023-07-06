SET search_path = public;

-- name: Create :one
INSERT INTO articles(user_id, category_id, slug, title, body) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: Update :one
UPDATE articles SET slug = $2, title = $3, body = $4 WHERE id = $1 RETURNING *;

-- name: ListGlobal :many
SELECT * FROM articles ORDER BY created_at DESC;

-- name: ListInCategory :many
SELECT * FROM articles WHERE category_id = $1 ORDER BY created_at DESC;

-- name: ListOfUser :many
SELECT * FROM articles WHERE user_id = $1 ORDER BY created_at DESC;

-- name: ListOfUserInCategory :many
SELECT * FROM articles WHERE user_id = $1 AND category_id = $2 ORDER BY created_at DESC;
