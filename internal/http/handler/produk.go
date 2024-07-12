package handler

import (
	"net/http"
	"strconv"

	"github.com/Kevinmajesta/backendpergudanganmi/internal/entity"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/http/binder"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/service"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProdukHandler struct {
	produkService service.ProdukService
}

func NewProdukHandler(produkService service.ProdukService) ProdukHandler {
	return ProdukHandler{produkService: produkService}
}

func (h *ProdukHandler) CreateProduk(c echo.Context) error {
	input := binder.ProdukCreateRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}
	newProduk := entity.NewProduk(input.Namabarang, input.Harga, input.Stok)
	produk, err := h.produkService.CreateProduk(newProduk)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully input a new products", produk))
}

func (h *ProdukHandler) UpdateProduk(c echo.Context) error {
	var input binder.ProdukUpdateRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}

	id, err := uuid.Parse(input.ProdukId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "invalid user ID"))
	}

	exists, err := h.produkService.CheckProdukExists(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "could not verify produk existence"))
	}
	if !exists {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "produk ID does not exist"))
	}

	inputUser := entity.UpdateProduk(id, input.Namabarang, input.Harga, input.Stok)

	updatedProduk, err := h.produkService.UpdateProduk(inputUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success update produk", updatedProduk))
}

func (h *ProdukHandler) FindAllProduct(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1 // default page if page parameter is invalid
	}

	products, err := h.produkService.FindAllProduct(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success show data products", products))
}

func (h *ProdukHandler) DeleteProduct(c echo.Context) error {
	var input binder.ProdukDeleteRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	id := uuid.MustParse(input.ProdukId)

	isDeleted, err := h.produkService.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses delete product", isDeleted))
}
