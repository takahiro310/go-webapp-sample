package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

// DBGet is xo database access sample
func DBGet(c *gin.Context) {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM test ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tests := []model.Test{}

	// iterate result
	for result.Next() {
		test := model.Test{}
		var id int64
		var name sql.NullString

		err = result.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}

		test.ID = id
		test.Name = name
		tests = append(tests, test)
	}
	fmt.Println(tests)
	c.JSON(http.StatusOK, gin.H{"tests": tests})
}
