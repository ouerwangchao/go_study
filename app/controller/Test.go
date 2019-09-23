package controller

import (
	"code_jzggxx.com/ouer/admin/app/database"
	"code_jzggxx.com/ouer/admin/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Test(c *gin.Context) {
	for i := 1; i <= 10000; i++ {
		go add(i)
	}
}

func Normal(c *gin.Context) {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
		add(i)
	}
}

func add(i int) {
	var AddData model.Test
	testData := &AddData
	testData.Value = i
	testData.CreatedAt = time.Now().Unix()
	database.MasterDB.Create(testData)
}
