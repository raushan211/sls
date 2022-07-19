package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

type Link struct {
	ID        int    `json:"id"`
	ShortLink string `json:"short_link"`
	LongLink  string `json:"long_link"`
}

var Data map[int]Link
var lastID int

func main() {
	createDBConnection()
	defer DB.Close()
	//Data = make(map[string]User)
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupRoutes(r *gin.Engine) {

	r.POST("/short_link/create", createHandler)
	r.GET("/:id", redirectHandler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

}

// POST
func SaveLongLink(c *gin.Context) {
	reqBody := Link{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": "invalid request body",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	lastID++
	reqBody.ID = lastID
	reqBody.ShortLink = "http://localhost:8080/srl/" + fmt.Sprint(lastID)
	Data[lastID] = reqBody
	c.JSON(http.StatusOK, reqBody)
	c.Writer.Header().Set("Content-Type", "application/json")
	return
}

// GET

func GetLongLink(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		res := gin.H{
			"error": "invalid request body",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		res := gin.H{
			"error": "invalid request body",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	if _, ok := Data[idInt]; !ok {
		res := gin.H{
			"error": "link not found",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Writer.Header().Set("Content-Type", "application/json")
		return
	}
	c.Redirect(http.StatusMovedPermanently, Data[idInt].LongLink)
	return
}
