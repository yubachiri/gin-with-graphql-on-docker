package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"

	"m-share/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {

	envErr := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
    if envErr != nil {
				// .env読めなかった場合の処理
				//log.Fatal("Error loading .env file")
				fmt.Printf("%v\n", "読めてないウホ🦍")
		}
	env := os.Getenv("ENV")
	fmt.Println(env)

	driverName := "mysql"
	DsName := os.Getenv("DSNAME")
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
