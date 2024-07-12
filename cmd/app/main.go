package main

import (
	"github.com/Kevinmajesta/backendpergudanganmi/configs"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/server"
)

func main() {
	_, err := configs.NewConfig(".env")
	checkError(err)

	srv := server.NewServer("app", nil, nil)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
