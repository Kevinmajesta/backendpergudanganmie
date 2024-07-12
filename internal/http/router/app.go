package router

import (
	"net/http"

	"github.com/Kevinmajesta/backendpergudanganmi/internal/http/handler"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/route"
)

const (
	Admin = "admin"
	User  = "user"
)

var (
	allRoles  = []string{Admin, User}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{User}
)

func PublicRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.LoginUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
	}
}

func PrivateRoutes(userHandler handler.UserHandler, produkHandler handler.ProdukHandler) []*route.Route {
	return []*route.Route{
		//produk
		{
			Method:  http.MethodPost,
			Path:    "/produk/add",
			Handler: produkHandler.CreateProduk,
			Roles:   onlyAdmin,
		},
		{
			Method:  http.MethodPut,
			Path:    "/produk/edit/:id_produk",
			Handler: produkHandler.UpdateProduk,
			Roles:   onlyAdmin,
		},
		{
			Method:  http.MethodGet,
			Path:    "/produk/all",
			Handler: produkHandler.FindAllProduct,
			Roles:   onlyAdmin,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/produk/delete",
			Handler: produkHandler.DeleteProduct,
			Roles:   onlyAdmin,
		},
	}
}
