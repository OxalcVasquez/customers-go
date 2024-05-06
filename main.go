package main

import (
	"customers-api/config"
	_ "customers-api/docs"
	"customers-api/router"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

// @title           Customers Compartamos API
// @version         1.0
// @description     Api RestFul for manage customers
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://customers-go.onrender.com
// @BasePath  /api/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
