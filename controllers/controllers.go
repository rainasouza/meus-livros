package controllers

import (
	"context"
	"meus-livros/firebase"
	"meus-livros/models"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

// @Summary      Cria um novo livro
// @Description  Adiciona um novo livro ao Firestore
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body  models.Book  true  "Livro"
// @Success      201   {object}  map[string]any
// @Failure      400   {object}  map[string]any
// @Failure      500   {object}  map[string]any
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	docRef, _, err := firebase.Client.Collection("books").Add(context.Background(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Livro criado com sucesso",
		"id":      docRef.ID, // Retorna o ID gerado automaticamente
	})
}

// @Summary      Retorna livros
// @Description  Busca todos os livros do Firestore
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200   {array}  models.Book
// @Failure      500   {object}  map[string]any
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	ctx := context.Background()
	docs, err := firebase.Client.Collection("books").Documents(ctx).GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var books []models.Book
	for _, doc := range docs {
		var book models.Book
		// Tenta preencher o endereço de memória &book com os dados do documento. Se algo der errado (ex.: tipos incompatíveis, campo ausente, dados corrompidos), retorna esse erro. Se der certo, o erro é nil.
		if err := doc.DataTo(&book); err == nil {
			books = append(books, book)
		}
	}

	c.JSON(http.StatusOK, books)
}

// @Summary      Atualiza um livro
// @Description  Atualiza parcialmente um livro pelo ID no Firestore
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id     path      string           true  "ID do Livro"
// @Param        book   body      object           true  "Campos para atualizar"
// @Success      200    {object}  map[string]any
// @Failure      400    {object}  map[string]any
// @Failure      500    {object}  map[string]any
// @Router       /books/{id} [patch]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	// Recebe um JSON genérico
	var updates map[string]any

	// Verifica se o JSON é válido
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	docRef := firebase.Client.Collection("books").Doc(id)

	// Constrói os updates para Firestore
	var firestoreUpdates []firestore.Update
	for key, value := range updates {
		firestoreUpdates = append(firestoreUpdates, firestore.Update{
			Path:  key,
			Value: value,
		})
	}

	// Faz a atualização parcial
	_, err := docRef.Update(ctx, firestoreUpdates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Livro atualizado com sucesso",
		"id":      id,
	})
}

// @Summary      Deleta um livro
// @Description  Remove um livro do Firestore pelo ID
// @Tags         books
// @Produce      json
// @Param        id   path      string           true  "ID do Livro"
// @Success      200  {object}  map[string]any
// @Failure      500  {object}  map[string]any
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	docRef := firebase.Client.Collection("books").Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso", "id": id})
}
