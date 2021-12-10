-- Create houses table
CREATE TABLE houses (
    id          UUID DEFAULT uuid_generate_v4() not null unique,
    description varchar(255) not null,
    address     varchar(255) not null,
    owner_id    UUID not null,
    created_at  timestamp with time zone not null default now()
);

ALTER TABLE houses ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");
CREATE INDEX ON houses ("owner_id");
