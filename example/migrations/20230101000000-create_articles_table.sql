-- +migrate Up
CREATE TABLE articles(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    category_id uuid NOT NULL,
    slug text NOT NULL,
    title text NOT NULL,
    body text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp,
    UNIQUE(user_id, slug)
);

CREATE TRIGGER refresh_articles_updated_at_step1
BEFORE UPDATE ON articles FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_articles_updated_at_step2
BEFORE UPDATE OF updated_at ON articles FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_articles_updated_at_step3
BEFORE UPDATE ON articles FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE INDEX articles_category_id_idx ON articles(category_id);
CREATE INDEX articles_created_at_idx ON articles(created_at);

COMMENT ON TABLE articles IS 'Articles';
COMMENT ON COLUMN articles.user_id IS 'Author user ID';
COMMENT ON COLUMN articles.slug IS 'Unique article slug per user';
COMMENT ON COLUMN articles.title IS 'Title';
COMMENT ON COLUMN articles.body IS 'Body';

-- +migrate Down
DROP TABLE articles;
