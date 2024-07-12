package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Kevinmajesta/backendpergudanganmi/internal/entity"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProdukRepository interface {
	CreateProduk(produk *entity.Products) (*entity.Products, error)
	CheckProdukExists(id uuid.UUID) (bool, error)
	UpdateProduk(produk *entity.Products) (*entity.Products, error)
	FindAllProduct(page int) ([]entity.Products, error)
	DeleteProduct(products *entity.Products) (bool, error)
	FindProductByID(id_produk uuid.UUID) (*entity.Products, error)
}

type produkRepository struct {
	db        *gorm.DB
	cacheable cache.Cacheable
}

func NewProdukRepository(db *gorm.DB, cacheable cache.Cacheable) *produkRepository {
	return &produkRepository{db: db, cacheable: cacheable}
}

func (r *produkRepository) CreateProduk(produk *entity.Products) (*entity.Products, error) {
	if err := r.db.Create(&produk).Error; err != nil {
		return produk, err
	}
	return produk, nil
}

func (r *produkRepository) CheckProdukExists(id uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.Model(&entity.Products{}).Where("id_produk = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *produkRepository) UpdateProduk(produk *entity.Products) (*entity.Products, error) {
	fields := make(map[string]interface{})

	if produk.Namabarang != "" {
		fields["namabarang"] = produk.Namabarang
	}
	if produk.Harga != "" {
		fields["harga"] = produk.Harga
	}
	if produk.Stok != "" {
		fields["stok"] = produk.Stok
	}

	if err := r.db.Model(produk).Where("id_produk = ?", produk.ProdukId).Updates(fields).Error; err != nil {
		return produk, err
	}

	return produk, nil
}

func (r *produkRepository) FindAllProduct(page int) ([]entity.Products, error) {
	var products []entity.Products
	key := fmt.Sprintf("FindAllProducts_page_%d", page)
	const pageSize = 10

	data, _ := r.cacheable.Get(key)
	if data == "" {
		offset := (page - 1) * pageSize
		if err := r.db.Limit(pageSize).Offset(offset).Find(&products).Error; err != nil {
			return products, err
		}
		marshalledproducts, _ := json.Marshal(products)
		err := r.cacheable.Set(key, marshalledproducts, 5*time.Minute)
		if err != nil {
			return products, err
		}
	} else {
		err := json.Unmarshal([]byte(data), &products)
		if err != nil {
			return products, err
		}
	}
	return products, nil
}

func (r *produkRepository) FindProductByID(id_produk uuid.UUID) (*entity.Products, error) {
	product := new(entity.Products)
	if err := r.db.Where("id_produk = ?", id_produk).Take(product).Error; err != nil {
		log.Printf("Error finding product by ID: %v", err)
		return product, err
	}
	log.Printf("Product found: %v", product)
	return product, nil
}

func (r *produkRepository) DeleteProduct(product *entity.Products) (bool, error) {
	log.Printf("Deleting product: %v", product)
	if err := r.db.Delete(product).Error; err != nil {
		log.Printf("Error deleting product: %v", err)
		return false, err
	}
	log.Println("Product deleted successfully")
	return true, nil
}
