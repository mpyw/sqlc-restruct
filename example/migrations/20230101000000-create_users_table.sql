-- +migrate Up
CREATE TYPE user_status AS ENUM ('active', 'inactive');
CREATE TABLE users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email text NOT NULL UNIQUE,
    name text NOT NULL,
    status user_status NOT NULL default 'active',
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp
);

CREATE TRIGGER refresh_users_updated_at_step1
BEFORE UPDATE ON users FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_users_updated_at_step2
BEFORE UPDATE OF updated_at ON users FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_users_updated_at_step3
BEFORE UPDATE ON users FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step3();

COMMENT ON TABLE users IS 'Users';
COMMENT ON COLUMN users.email IS 'Email';
COMMENT ON COLUMN users.name IS 'Name';

-- +migrate Down
DROP TABLE users;
DROP TYPE user_statuses;
