-- In the interest of time, I'm choosing to directly use the internal
-- integer ID as an external ID. In practice, IDs that are shared outside
-- the database for use by users would be a separately indexed value
-- like a UUID.

-- I'm skipping foreign key cascade details in interest of time.

CREATE TABLE organizations (
    org_id SERIAL PRIMARY KEY
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    org_id integer NOT NULL REFERENCES organizations(org_id)
);

CREATE TABLE services (
    service_id SERIAL PRIMARY KEY,
    title text NOT NULL,
    summary text NOT NULL,
    org_id integer NOT NULL REFERENCES organizations(org_id),
    UNIQUE (title, org_id)
);

CREATE TABLE versions (
    version_id SERIAL NOT NULL,
    service_id integer NOT NULL REFERENCES services(service_id),
    summary text NOT NULL,
    PRIMARY KEY (version_id, service_id)
);
