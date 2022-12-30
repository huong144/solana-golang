package schema

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// connect db
	ConnectDB()
}
func ConnectDB() *gorm.DB {
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
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(1000000)
	//db.AutoMigrate(&Account{}, &Transaction{}, &BalanceChange{}, &Instruction{}, &InnerInstruction{})
	return db
}

func AutoMigration() any {
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
	if err != nil {
		log.Printf("Cannot connected to database!")
	}
	//sqlDB, err := db.DB()
	//sqlDB.SetMaxIdleConns(1000000)
	db.AutoMigrate(
		&Account{},
		&Transaction{},
		&BalanceChange{},
		&Instruction{},
		&InnerInstruction{},
		&SignatureAddressRel{},
		&UserRegisterWebhook{},
		&CheckBlockFlag{})
	return nil
}
