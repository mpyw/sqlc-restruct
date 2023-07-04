-- +migrate Up
ALTER TABLE articles
ADD CONSTRAINT articles_user_id_fkey
FOREIGN KEY(user_id) REFERENCES users(id),
ADD CONSTRAINT articles_category_id_fkey
FOREIGN KEY(category_id) REFERENCES categories(id);

-- +migrate Down
ALTER TABLE articles
DROP CONSTRAINT articles_user_id_fkey,
DROP CONSTRAINT articles_category_id_fkey;
