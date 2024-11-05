package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/FranklinDevWork/k8s-stock-tracker/api/clients"
	"github.com/FranklinDevWork/k8s-stock-tracker/api/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		avClient := clients.AlphaVantageClient{
			Url:    getEnv("API_URI", "https://www.alphavantage.co/query"),
			ApiKey: getEnv("API_KEY", ""),
		}
		symbol := getEnv("SYMBOL", "MSFT")
		nDays := getEnv("NDAYS", "7")
		response, err := avClient.MakeQueryRequest(symbol, "TIME_SERIES_DAILY")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Issues getting data",
			})
			return
		}
		if nDaysInt, err := strconv.ParseInt(nDays, 0, 32); err == nil {
			c.JSON(http.StatusOK, helpers.LastNDaysFromAV(response, int32(nDaysInt)))
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Error",
		})
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Healthy",
		})
	})

	router.Run(":8080")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
