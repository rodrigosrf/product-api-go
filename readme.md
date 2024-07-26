# API de Gerenciamento de Produtos em Go

Este repositório contém uma API RESTful para gerenciamento de produtos, construída usando Go (Golang). A API permite a criação, recuperação, atualização e exclusão de registros de produtos.

## Funcionalidades

- Criar, ler, atualizar e deletar produtos
- Endpoints RESTful
- Integração com MongoDB para persistência de dados

## Primeiros Passos

Estas instruções ajudarão você a configurar e executar o projeto em sua máquina local para desenvolvimento e testes.

### Pré-requisitos

- [Go](https://golang.org/doc/install) 1.16 ou superior
- [MongoDB](https://www.mongodb.com/try/download/community) (rodando localmente ou acessível remotamente)
- [Git](https://git-scm.com/)

### Instalação

1. Clone o repositório:

    ```bash
    git clone https://github.com/seu-usuario/product-management-api-go.git
    cd product-management-api-go
    ```

2. Instale os pacotes Go necessários:

    ```bash
    go mod tidy
    ```

3. Configure as variáveis de ambiente. Crie um arquivo `.env` na raiz do projeto e adicione as seguintes variáveis:

    ```plaintext
    MONGODB_URI=mongodb://localhost:27017
    MONGODB_DB=productdb
    PORT=8080
    ```

4. Inicie o servidor da API:

    ```bash
    go run main.go
    ```

A API deve estar rodando em `http://localhost:8080`.

## Uso

A API fornece os seguintes endpoints:

- `GET /products` - Recupera todos os produtos
- `GET /products/{id}` - Recupera um produto específico por ID
- `POST /products` - Cria um novo produto
- `PUT /products/{id}` - Atualiza um produto existente por ID
- `DELETE /products/{id}` - Deleta um produto por ID

### Exemplos de Requisições

- **Recuperar todos os produtos**

    ```bash
    curl -X GET http://localhost:8080/products
    ```

- **Recuperar um produto por ID**

    ```bash
    curl -X GET http://localhost:8080/products/{id}
    ```

- **Criar um novo produto**

    ```bash
    curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{"name": "Novo Produto", "price": 99.99}'
    ```

- **Atualizar um produto por ID**

    ```bash
    curl -X PUT http://localhost:8080/products/{id} -H "Content-Type: application/json" -d '{"name": "Produto Atualizado", "price": 79.99}'
    ```

- **Deletar um produto por ID**

    ```bash
    curl -X DELETE http://localhost:8080/products/{id}
    ```

## Estrutura do Projeto

