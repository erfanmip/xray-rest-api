package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// GET endpoint at the root path
	router.GET("/", func(c *gin.Context) {

		userEmail := c.Query("userEmail")
		var (
			xrayCtl *XrayController
			cfg     = &BaseConfig{
				APIAddress: "127.0.0.1",
				APIPort:    10085,
			}
		)
		xrayCtl = new(XrayController)
		err := xrayCtl.Init(cfg)
		defer xrayCtl.CmdConn.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}

		ptn := "user>>>" + userEmail + ">>>traffic>>>downlink"
		trafficData, err := queryTraffic(xrayCtl.SsClient, ptn, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}
		fmt.Println(trafficData)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})
	router.Run(":8080")
}
