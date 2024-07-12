BEGIN;

CREATE TABLE IF NOT EXISTS supir (
    id_supir UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    kendaraan VARCHAR(255) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    nohp VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    tanggallahir VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;