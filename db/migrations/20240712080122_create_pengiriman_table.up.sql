BEGIN;

CREATE TABLE IF NOT EXISTS pengiriman (
    id_pengiriman UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_resi UUID NOT NULL,
    id_supir UUID NOT NULL,
    namasupir VARCHAR(255) NOT NULL,
    id_transaksi UUID NOT NULL,
    kendaraan VARCHAR(255) NOT NULL,
    statuspengiriman BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;