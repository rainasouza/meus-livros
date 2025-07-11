
# 📚 Meus Livros - API REST com Go e Firebase

API REST desenvolvida em **Go (Golang)** para gerenciamento de livros, utilizando **Firebase** como banco de dados. A documentação da API é gerada automaticamente com **Swagger**.

---

## 🚀 Funcionalidades

- 🔐 Integração com Firebase
- 📚 CRUD completo de livros
- 📝 Documentação interativa com Swagger

---

## 📦 Estrutura do Projeto

```bash
.
├── controllers/            # Handlers da API
├── firebase/               # Configuração e inicialização do Firebase
├── models/                 # Modelos de dados
├── routes/                 # Definição das rotas da API
├── docs/                   # Documentação gerada pelo Swagger
├── main.go                 # Arquivo principal
├── go.mod / go.sum         # Gerenciamento de dependências
└── serviceAccountKey.json  # Chave de serviço do Firebase
```

---

## ⚙️ Como rodar o projeto

### ✔️ Pré-requisitos

- Go instalado (versão 1.18 ou superior)
- Conta no Firebase com a chave JSON
- Git

### 🚀 Passos para execução

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/meus-livros.git
cd meus-livros

# Instale as dependências
go mod tidy

# Gere a documentação Swagger
swag init

# Execute o servidor
go run main.go
```

---

## 📄 Documentação da API

Acesse a documentação Swagger em:

```
http://localhost:8080/swagger/index.html
```

### 🔍 Exemplo de anotação Swagger:

```go
// @Summary      Lista todos os livros
// @Description  Retorna todos os livros do banco de dados
// @Tags         livros
// @Produce      json
// @Success      200  {array}  models.Livro
// @Router       /livros [get]
```

---

## 👨‍💻 Autor

Feito com 💙 por [Raína Souza](https://github.com/rainasouza)

