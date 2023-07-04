-- +migrate Up
CREATE TABLE categories(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    slug text NOT NULL UNIQUE,
    title text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp
);

CREATE TRIGGER refresh_categories_updated_at_step1
BEFORE UPDATE ON categories FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_categories_updated_at_step2
BEFORE UPDATE OF updated_at ON categories FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_categories_updated_at_step3
BEFORE UPDATE ON categories FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step3();

COMMENT ON TABLE categories IS 'Categories';
COMMENT ON COLUMN categories.slug IS 'Globally unique slug';
COMMENT ON COLUMN categories.title IS 'Title';

-- +migrate Down
DROP TABLE categories;
