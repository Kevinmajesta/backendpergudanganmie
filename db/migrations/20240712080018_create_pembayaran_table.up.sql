BEGIN;

CREATE TABLE IF NOT EXISTS pembayaran (
    id_transaksi UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    total_bayar VARCHAR(255) NOT NULL,
    id_produk UUID NOT NULL,
    jumlahbarang VARCHAR(255) NOT NULL,
    id_user UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;