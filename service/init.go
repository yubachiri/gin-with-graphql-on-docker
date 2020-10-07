package service

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"app/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		// .env読めなかった場合の処理
		//log.Fatal("Error loading .env file")
		fmt.Printf("%v\n", "DB connection error")
	}
	env := os.Getenv("ENV")
	fmt.Println(env)

	driverName := "mysql"

	user := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	connectMethod := os.Getenv("DB_CONNECT_METHOD")
	containerName := os.Getenv("DB_CONTAINER_NAME")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	DsName := user + ":" + password + "@" + connectMethod + "(" + containerName + ":" + port + ")/" + name

	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
