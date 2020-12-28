package main

import (
	"github.com/marius/moduleProject/login"
)

func main() {

	loginRepository := login.Repository{}
	loginService := login.Service{
		Repository: &loginRepository,
	}

	setupRouter(
		loginService,
	).Run()
}
