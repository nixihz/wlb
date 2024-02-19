package clients

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"wlb/internal/conf"
)

func NewSqlite3Client(data *conf.Data) *gorm.DB {
	db, err := gorm.Open(data.Database.Driver, data.Database.Source)
	db.SingularTable(true)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
