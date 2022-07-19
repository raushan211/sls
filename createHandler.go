package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func createHandler(c *gin.Context) {
	reqBody := Link{}
	err := c.Bind(&reqBody)

	if err != nil {
		if err != nil {
			ress := gin.H{
				"error": err.Error(),
			}

			c.JSON(http.StatusBadRequest, ress)
			return
		}
	}

	bool_result, count, id := isLongUrlExist(reqBody.LongLink)

	if bool_result == false {
		res := gin.H{
			"error": "something went wrong",
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if count > 0 {
		idd := strconv.Itoa(id)
		//short_link:=`localhost:8082:"+id"`
		res := gin.H{
			"success":    true,
			"long_link":  reqBody.LongLink,
			"short_link": "localhost:8080:" + idd,
		}
		c.JSON(http.StatusOK, res)
		return
	}
	if bool_result == true && (count == -1 || count == 0) {
		_, err_result, result_id := CreateLink(reqBody)
		result_idd := strconv.Itoa(result_id)

		if err_result != "" {
			res := gin.H{
				"error": err_result,
			}
			//c.Writer.Header().Set("Content-Type", "application/json")

			c.JSON(http.StatusBadRequest, res)
			return
		}

		res := gin.H{
			"success":    true,
			"long_link":  reqBody.LongLink,
			"short_link": "localhost:8080/" + result_idd,
		}
		c.JSON(http.StatusOK, res)
	}
}

func CreateLink(reqbody Link) (bool, string, int) {
	var result = true
	var err_responce = ""
	var id = 0
	var currentTime = time.Now().UTC()
	expire_at := currentTime.Add(time.Minute * 5)

	sqlStatement := `
INSERT INTO link(long_link,short_link,expire_at)
VALUES ($1,$2,$3) RETURNING id`
	err2 := DB.QueryRow(sqlStatement, reqbody.LongLink, "test", expire_at.Format(time.RFC3339Nano)).Scan(&id)

	fmt.Println(err2)

	if err2 != nil {
		err_responce = "Something went wrong"
		return false, err_responce, id
	}

	fmt.Println(id)

	result = false
	return result, err_responce, id

}
