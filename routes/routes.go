package routes

import (
	"meus-livros/controllers"
	_ "meus-livros/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()

	// Rota para criar um livro
	r.POST("/books", controllers.CreateBook)

	// Rota para a raiz
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API de livros rodando!"})
	})

	r.GET("/books", controllers.GetBooks)

	// Rota para atualizar um livro
	r.PATCH("/books/:id", controllers.UpdateBook)

	// Rota para deletar um livro
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // Inicia o servidor na porta padrão (8080)
}
