package Database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() (*gorm.DB, error) {
	godotenv.Load(".env")
	var (
		dbHost     = os.Getenv("DATABASE_HOST")
		dbPort     = os.Getenv("DATABASE_PORT")
		dbName     = os.Getenv("DATABASE_NAME")
		dbUserName = os.Getenv("DATABASE_USERNAME")
		dbPassWord = os.Getenv("DATABASE_PASSWORD")
	)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := `host=` + dbHost + ` user=` + dbUserName + ` password=` + dbPassWord + ` dbname=` + dbName + ` port=` + dbPort + ` sslmode=disable TimeZone=Asia/Shanghai`

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(1000000)
	//db.AutoMigrate(&Account{}, &Transaction{}, &BalanceChange{}, &Instruction{}, &InnerInstruction{})
	return db, err
}
