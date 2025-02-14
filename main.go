package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {

	items := []models.Item{
		{Id: 1, Name: "item1", Price: 1000, Description: "description1", Soldout: false},
		{Id: 2, Name: "item2", Price: 2000, Description: "description2", Soldout: true},
		{Id: 3, Name: "item3", Price: 3000, Description: "description3", Soldout: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
