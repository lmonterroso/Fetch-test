package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	Description string  `json:"shortDescription`
	Price       float64 `json:"price"`
}
type receipt struct {
	Retailer string  `json:"retailer"`
	Date     string  `json:"purchaseDate"`
	Time     string  `json:"purchaseTime"`
	Total    float64 `json:"total"`
	Items    []item  `json:"items"`
}

func getRecepit(c *gin.Context) {
	receipt := "asds"
	c.IndentedJSON(http.StatusOK, receipt)
}

func processReceipts(c *gin.Context) {
	var newReceipt receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

}
func main() {
	router := gin.Default()
	router.GET("/receipt", getRecepit)
	router.Run("localhost:8080")
}
