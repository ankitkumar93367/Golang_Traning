import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase(){
	dsn := "root:root@tcp(127.0.0.1:3306)/MyDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Faild to connect to db")
	}

	db.AutoMigrate(Book{})

	DB = db
}