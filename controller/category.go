package controller

import (
	"codezard-pos/db"
	"codezard-pos/dto"
	"codezard-pos/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct{}

func (c Category) FindAll(ctx *gin.Context) {
	var categories []model.Category
	db.Conn.Find(&categories)

	var result []dto.CategoryResponse
	for _, category := range categories {
		result = append(result, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
			// CreatedAt: category.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Category) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})

}

// POST /categories <<--{"name":"Floower"} type-JSON
func (c Category) Create(ctx *gin.Context) {
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		Name: form.Name,
	}

	if err := db.Conn.Create(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ส่งค่าออกไป คือ ID , Name
	ctx.JSON(http.StatusCreated, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func (c Category) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	category.Name = form.Name
	db.Conn.Save(&category)
	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})

}

func (c Category) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var category model.Category

	if err := db.Conn.First(&category, id).Delete(&model.Category{}, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"sofe_delete": time.Now()})
	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})
}

func (c Category) Selete2(ctx *gin.Context) {
	id := ctx.Param("id")

	marksStr := id
	marks, err := strconv.Atoi(marksStr)
	if err != nil {
		fmt.Println("Error during conversion")
		ctx.JSON(http.StatusOK, gin.H{"Error during conversion": err})

		return
	}

	if marks%2 == 0 {
		ctx.JSON(http.StatusOK, gin.H{"even(เลขคู่)": marks})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"odd(เลขคี่)": marks})
	}

	var category model.Category
	if err := db.Conn.First(&category, id).Unscoped().Delete(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"hard_delete": time.Now()})
	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"ID": id,
	})

}
