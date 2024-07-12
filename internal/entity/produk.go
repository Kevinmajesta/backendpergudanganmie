package entity

import (
	"github.com/google/uuid"
)

type Products struct {
	ProdukId   uuid.UUID `json:"id_produk" gorm:"column:id_produk;primaryKey"`
	Namabarang string    `json:"namabarang"`
	Harga      string    `json:"harga"`
	Stok       string    `json:"stok"`
	Auditable
}

func NewProduk(namabarang, harga, stok string) *Products {
	return &Products{
		ProdukId:   uuid.New(),
		Namabarang: namabarang,
		Harga:      harga,
		Stok:       stok,
		Auditable:  NewAuditable(),
	}
}

func UpdateProduk(id_produk uuid.UUID, namabarang, harga, stok string) *Products {
	return &Products{
		ProdukId:   id_produk,
		Namabarang: namabarang,
		Harga:      harga,
		Stok:       stok,
		Auditable:  UpdateAuditable(),
	}
}
