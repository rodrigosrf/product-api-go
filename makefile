# Nome do binário
BINARY_NAME=api-product-go-lab

# Diretório para os arquivos gerados
BUILD_DIR=bin

# Variáveis de ambiente para o MongoDB
export $(cat .env | xargs)

.PHONY: all build run clean mod test swagger-ui swagger

all: build

# Compilar o binário
build:
	@echo "Building the binary..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

# Rodar o binário
run: build
	@echo "Running the application..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Limpar os arquivos gerados
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)/*
	rm -rf docs

swagger:
	@echo "Generating Swagger documentation..."
	swag init -g main.go --output docs

# Remover/Adicionar pacotes necessarios
mod:
	@echo "Update mods..."
	go mod tidy

# Rodar testes
test:
	@echo "Running tests..."
	go test ./...

# Rodar o Swagger UI para visualizar a documentação
swagger-ui:
	@echo "Running Swagger UI..."
	docker-compose up -d swagger-ui

