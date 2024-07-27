package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rodrigosrf/api-product-go-lab/database"
	"github.com/rodrigosrf/api-product-go-lab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	repository *database.ProductRepository
)

func InitializeHandler(db *mongo.Database) {
	repository = database.NewProductRepository(db)
}

// ShowAccount godoc
// @Summary      Create product
// @Description  Create product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param request body models.Product true "query params"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/product/{id} [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.Code = uuid.New().String()

	ctx := context.Background()
	err := repository.Create(ctx, &newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": newProduct.Code})
}

// ShowAccount godoc
// @Summary      List products
// @Description  Get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200 {object} []models.Product
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/product [get]
func GetProducts(c *gin.Context) {
	ctx := context.Background()

	products, err := repository.FindAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// ShowAccount godoc
// @Summary      Get product
// @Description  Get product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200 {object} models.Product
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/product/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	product, err := repository.FindByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// ShowAccount godoc
// @Summary      Update product
// @Description  Update product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Param request body models.Product true "query params"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/product/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updateProduct models.Product

	if err := c.BindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	update := bson.M{
		"name":        updateProduct.Name,
		"description": updateProduct.Description,
		"price":       updateProduct.Price,
	}
	err := repository.Update(ctx, id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// ShowAccount godoc
// @Summary      Delete product
// @Description  Delete product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	err := repository.Delete(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
