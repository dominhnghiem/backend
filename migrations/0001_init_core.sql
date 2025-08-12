BEGIN;
DROP TABLE IF EXISTS follows CASCADE;
DROP TABLE IF EXISTS likes CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS verification_tokens CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TYPE  IF EXISTS token_purpose;

CREATE TYPE token_purpose AS ENUM ('email_verify', 'password_reset');

CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       email TEXT NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       name TEXT,
                       email_verified_at TIMESTAMPTZ,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       deleted_at TIMESTAMPTZ
);

CREATE TABLE sessions (
                          id BIGSERIAL PRIMARY KEY,
                          user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          refresh_token_hash TEXT NOT NULL UNIQUE,
                          user_agent TEXT,
                          ip INET,
                          expires_at TIMESTAMPTZ NOT NULL,
                          revoked_at TIMESTAMPTZ,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX sessions_user_idx ON sessions (user_id);

CREATE TABLE verification_tokens (
                                     id BIGSERIAL PRIMARY KEY,
                                     user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                     purpose token_purpose NOT NULL,
                                     token_hash TEXT NOT NULL,
                                     expires_at TIMESTAMPTZ NOT NULL,
                                     consumed_at TIMESTAMPTZ,
                                     created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX verification_tokens_unique_active
    ON verification_tokens (user_id, purpose, token_hash)
    WHERE consumed_at IS NULL;

CREATE TABLE posts (
                       id BIGSERIAL PRIMARY KEY,
                       author_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       body TEXT NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       deleted_at TIMESTAMPTZ
);
CREATE INDEX posts_author_created_at_idx ON posts (author_id, created_at DESC);

CREATE TABLE comments (
                          id BIGSERIAL PRIMARY KEY,
                          post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
                          author_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          body TEXT NOT NULL,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                          deleted_at TIMESTAMPTZ
);
CREATE INDEX comments_post_created_at_idx ON comments (post_id, created_at);
CREATE INDEX comments_author_created_at_idx ON comments (author_id, created_at);

CREATE TABLE likes (
                       user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       PRIMARY KEY (user_id, post_id)
);
CREATE INDEX likes_post_idx ON likes (post_id);

CREATE TABLE follows (
                         follower_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                         followee_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                         PRIMARY KEY (follower_id, followee_id),
                         CHECK (follower_id <> followee_id)
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
RETURN NEW;
END; $$ LANGUAGE plpgsql;

CREATE TRIGGER set_users_updated_at   BEFORE UPDATE ON users   FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER set_posts_updated_at   BEFORE UPDATE ON posts   FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER set_comments_updated_at BEFORE UPDATE ON comments FOR EACH ROW EXECUTE FUNCTION set_updated_at();
COMMIT;
