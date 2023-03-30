SELECT 'CREATE DATABASE aprixia' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'aprixia')\gexec
CREATE TABLE IF NOT EXISTS urldata (
    real_url varchar(256),
    short_url varchar(7) UNIQUE PRIMARY KEY,
    create_at timestamp,
    redirect_count int8
);

CREATE INDEX idx_short_url on urldata (short_url);
