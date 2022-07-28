package test

import (
	"fmt"
	"testing"

	"getcharzp.cn/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {
	dsn := "root:123456@tcp(localhost:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("1111")
	}

	data := make([]*models.ProblemBasic, 0)
	err = db.Take(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("Problem==> %v \n", v)
	}
}
