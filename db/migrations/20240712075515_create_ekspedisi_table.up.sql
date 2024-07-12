BEGIN;

CREATE TABLE IF NOT EXISTS ekpedisi (
    id_ekspedisi UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_resi UUID NOT NULL,
    id_transaksi UUID NOT NULL,
    id_rute UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;