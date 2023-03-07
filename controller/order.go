package controller

import (
	"codezard-pos/db"
	"codezard-pos/dto"
	"codezard-pos/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct{}

func (o Order) FindAll(ctx *gin.Context) {

}

func (o Order) FindOne(ctx *gin.Context) {

}

func (o Order) Create(ctx *gin.Context) {
	var form dto.OrderRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order model.Order
	var orderItems []model.OrderItem
	for _, product := range form.Products {
		orderItems = append(orderItems, model.OrderItem{
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}
	order.Name = form.Name
	order.Tel = form.Tel
	order.Email = form.Email
	order.Products = orderItems
	db.Conn.Create(&order)

	result := dto.OrderResponse{
		ID:    order.ID,
		Name:  order.Name,
		Tel:   order.Tel,
		Email: order.Email,
	}
	var products []dto.OrderProductResponse
	for _, product := range order.Products {
		products = append(products, dto.OrderProductResponse{
			ID:       product.ID,
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
			Image:    product.Image,
		})
	}
	result.Products = products

	ctx.JSON(http.StatusCreated, result)
}
