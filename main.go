package main

import (
	"meus-livros/firebase"
	"meus-livros/routes"
)

// @title           Meus Livros API
// @version         1.0
// @description     API para gerenciamento de livros
// @host            localhost:8080
// @BasePath        /
func main() {
	firebase.InitializeFirebase()
	routes.HandleRequests()

}
