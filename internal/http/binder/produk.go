package binder

type ProdukCreateRequest struct {
	Namabarang string `json:"namabarang" validate:"required,email"`
	Harga      string `json:"harga" validate:"required"`
	Stok       string `json:"stok" validate:"required"`
}

type ProdukUpdateRequest struct {
	ProdukId   string `param:"id_produk" validate:"required"`
	Namabarang string `json:"namabarang" validate:"required,email"`
	Harga      string `json:"harga" validate:"required"`
	Stok       string `json:"stok" validate:"required"`
}

type ProdukDeleteRequest struct {
	ProdukId string `json:"id_produk" validate:"required"`
}
