CREATE TABLE certficates (
    created_at timestamp WITH time zone,
    deleted_at timestamp WITH time zone,
    host text NOT NULL,
    id uuid,
    name text,
    notes text,
    port text NOT NULL,
    updated_at timestamp WITH time zone,
    PRIMARY KEY (
        id
)
);

ALTER TABLE certficates
    ADD CONSTRAINT unique_host_port UNIQUE (host, port);

