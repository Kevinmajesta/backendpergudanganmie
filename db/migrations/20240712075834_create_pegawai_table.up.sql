BEGIN;

CREATE TABLE IF NOT EXISTS pegawai (
    id_pegawai UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    namapegawai VARCHAR(255) NOT NULL,
    id_gudang UUID NOT NULL,
    alamatpegawai TEXT NOT NULL,
    nohp VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;