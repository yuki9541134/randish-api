package main

import (
	"github.com/joho/godotenv"
	"github.com/yuki9541134/randish-api/handler"
	"github.com/yuki9541134/randish-api/infra"
	"github.com/yuki9541134/randish-api/infra/repository"
	"github.com/yuki9541134/randish-api/usecase"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB接続
	db := infra.ConnectDB()

	dishRepository := repository.NewDishRepository(db)
	dishUseCase := usecase.NewDishUseCase(dishRepository)
	dishHandler := handler.NewDishHandler(dishUseCase)

	http.HandleFunc("/", dishHandler.HandleGetAll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
