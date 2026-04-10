package main

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseFloat(c *gin.Context, key string) (float64, error) {
	return strconv.ParseFloat(c.Param(key), 64)
}

func main() {
	r := gin.Default()

	// sqrt
	r.GET("/sqrt/:x", func(c *gin.Context) {
		x, err := parseFloat(c, "x")
		if err != nil || x < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "x must be >= 0"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"calculation": "sqrt(" + c.Param("x") + ")",
			"result":      math.Sqrt(x),
		})
	})

	// mod
	r.GET("/mod/:a/:b", func(c *gin.Context) {
		a, err1 := parseFloat(c, "a")
		b, err2 := parseFloat(c, "b")

		if err1 != nil || err2 != nil || b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or division by zero"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"calculation": c.Param("a") + "%" + c.Param("b"),
			"result":      math.Mod(a, b),
		})
	})

	// pow
	r.GET("/pow/:a/:b", func(c *gin.Context) {
		a, err1 := parseFloat(c, "a")
		b, err2 := parseFloat(c, "b")

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"calculation": c.Param("a") + "^" + c.Param("b"),
			"result":      math.Pow(a, b),
		})
	})

	// natural log
	r.GET("/log/:x", func(c *gin.Context) {
		x, err := parseFloat(c, "x")
		if err != nil || x <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "x must be > 0"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"calculation": "log(" + c.Param("x") + ")",
			"result":      math.Log(x),
		})
	})

	r.Run(":5000")
}
