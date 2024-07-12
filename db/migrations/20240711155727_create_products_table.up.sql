BEGIN;

CREATE TABLE IF NOT EXISTS products (
    id_produk UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    namabarang VARCHAR(255) NOT NULL,
    harga VARCHAR(255) NOT NULL,
    stok VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;