BEGIN;

CREATE TABLE IF NOT EXISTS barangmasuk (
    id_barangmasuk UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_produk UUID NOT NULL,
    id_gudang UUID NOT NULL,
    jml_masuk VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;