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

func PrivateRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		//user
		{
			Method:  http.MethodPut,
			Path:    "/users/:id_user",
			Handler: userHandler.UpdateUser,
			Roles:   onlyAdmin,
		},
	}
}
