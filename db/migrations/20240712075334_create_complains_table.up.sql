BEGIN;

CREATE TABLE IF NOT EXISTS complains (
    id_complain UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_user UUID NOT NULL,
    massage TEXT NOT NULL,
    id_transaksi UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;