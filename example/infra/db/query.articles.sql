SET search_path = public;

-- name: ArticleCreate :one
INSERT INTO articles(user_id, category_id, slug, title, body) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: ArticleUpdate :one
UPDATE articles SET slug = $2, title = $3, body = $4 WHERE id = $1 RETURNING *;

-- name: ArticleListGlobal :many
SELECT * FROM articles ORDER BY created_at DESC;

-- name: ArticleListInCategory :many
SELECT * FROM articles WHERE category_id = $1 ORDER BY created_at DESC;

-- name: ArticleListOfUser :many
SELECT * FROM articles WHERE user_id = $1 ORDER BY created_at DESC;

-- name: ArticleListOfUserInCategory :many
SELECT * FROM articles WHERE user_id = $1 AND category_id = $2 ORDER BY created_at DESC;
