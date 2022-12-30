package Database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func ConnectDatabase(dbName string) (*mongo.Collection, context.Context, *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//if e != nil {
	//	log.Fatalln("error 1", e)
	//}
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	databaseName := os.Getenv("MONGODB_DATABASE")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatalln(err)
	}
	database := client.Database(databaseName)
	repository := database.Collection(dbName)
	return repository, ctx, client
}

func ConnectEVMDB() {

}
