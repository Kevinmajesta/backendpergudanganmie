package service

import (
	"errors"
	"log"

	"github.com/Kevinmajesta/backendpergudanganmi/internal/entity"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/repository"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/encrypt"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/token"
	"github.com/google/uuid"
)

type ProdukService interface {
	CreateProduk(produk *entity.Products) (*entity.Products, error)
	CheckProdukExists(id uuid.UUID) (bool, error)
	UpdateProduk(produk *entity.Products) (*entity.Products, error)
	FindAllProduct(page int) ([]entity.Products, error)
	DeleteProduct(id_produk uuid.UUID) (bool, error)
}

type produkService struct {
	produkRepository repository.ProdukRepository
	tokenUseCase     token.TokenUseCase
	encryptTool      encrypt.EncryptTool
}

func NewProdukService(produkRepository repository.ProdukRepository, tokenUseCase token.TokenUseCase,
	encryptTool encrypt.EncryptTool) *produkService {
	return &produkService{
		produkRepository: produkRepository,
		tokenUseCase:     tokenUseCase,
		encryptTool:      encryptTool,
	}
}

func (s *produkService) CreateProduk(produk *entity.Products) (*entity.Products, error) {
	if produk.Namabarang == "" {
		return nil, errors.New("Namabarang cannot be empty")
	}
	if produk.Harga == "" {
		return nil, errors.New("Harga cannot be empty")
	}
	if produk.Stok == "" {
		return nil, errors.New("Stok cannot be empty")
	}

	newUser, err := s.produkRepository.CreateProduk(produk)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *produkService) CheckProdukExists(id uuid.UUID) (bool, error) {
	return s.produkRepository.CheckProdukExists(id)
}

func (s *produkService) UpdateProduk(produk *entity.Products) (*entity.Products, error) {
	if produk.Namabarang == "" {
		return nil, errors.New("email cannot be empty")
	}
	if produk.Harga == "" {
		return nil, errors.New("password cannot be empty")
	}
	if produk.Stok == "" {
		return nil, errors.New("fullname cannot be empty")
	}

	updatedProduk, err := s.produkRepository.UpdateProduk(produk)
	if err != nil {
		return nil, err
	}

	return updatedProduk, nil
}

func (s *produkService) FindAllProduct(page int) ([]entity.Products, error) {
	return s.produkRepository.FindAllProduct(page)
}

func (s *produkService) DeleteProduct(id_produk uuid.UUID) (bool, error) {
	product, err := s.produkRepository.FindProductByID(id_produk)
	if err != nil {
		return false, err
	}

	log.Printf("Product to be deleted: %v", product)
	return s.produkRepository.DeleteProduct(product)
}
