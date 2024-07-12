BEGIN;

CREATE TABLE IF NOT EXISTS outlet (
    id_outlet UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_pegawai UUID NOT NULL,
    alamatoutlet VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;