package main

import (
	"Gin-MVC/backend/config"
	"Gin-MVC/backend/routes"
)

func main() {
	// ルーターのセットアップ
	r := routes.SetupRouter()

	// サーバー起動
	r.Run(config.Port)
}
