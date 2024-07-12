BEGIN;

CREATE TABLE IF NOT EXISTS rute (
    id_rute UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_supir UUID NOT NULL,
    waktutempuh VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;