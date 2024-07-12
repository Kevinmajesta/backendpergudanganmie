BEGIN;

CREATE TABLE IF NOT EXISTS gudang (
    id_gudang UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    alamat TEXT NOT NULL,
    kapasitas VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;