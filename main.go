package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ReceiptDB = make(map[string]receipt)

type item struct {
	Description string  `json:"shortDescription"`
	Price       float64 `json:"price,string"`
}

type receipt struct {
	Retailer string  `json:"retailer"`
	Date     string  `json:"purchaseDate"`
	Time     string  `json:"purchaseTime"`
	Total    float64 `json:"total,string"`
	Items    []item  `json:"items"`
}

func getRecepit(c *gin.Context) {
	receipt := "asds"
	c.IndentedJSON(http.StatusOK, receipt)
}

func processReceipts(c *gin.Context) {
	var newReceipt receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//Marshal newReceipt so we can encode in SHA256
	out, err := json.Marshal(newReceipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	h := sha1.New()
	h.Write([]byte(out))
	ReceiptID := hex.EncodeToString(h.Sum(nil))
	fmt.Print(ReceiptID)
	ReceiptDB[ReceiptID] = newReceipt

	c.JSON(http.StatusOK, ReceiptID)
}
func main() {
	router := gin.Default()
	router.GET("/receipt/:id", getRecepit)
	router.POST("/receipts/process", processReceipts)
	router.Run("localhost:8080")
}
