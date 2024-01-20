package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		for i == status {
			return true
		}
	}
	return false
}

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		log.Print("begining database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("transaction commit error: ", err)
			} else {
				log.Print("rollback transaction due to status code: ", c.Writer.Status())
				txHandle.Rollback()
			}
		}
	}
}
