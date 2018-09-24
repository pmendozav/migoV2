package main

import (
	"migoV2/infrastructure/server"
	"migoV2/infrastructure/db"
	"migoV2/interfaces/repositories"
	"migoV2/interfaces/services"
	"migoV2/usecases/login"

)

func main() {
	serverHandler := server.NewServerHandler()
	dbHandler := db.NewOrmHandler()

	handlers := make(map[string] repositories.DbHandler)
    handlers["DbBongoUserRepo"] = dbHandler
    // handlers["DbCustomerRepo"] = dbHandler

	createUserInteractor := new(login.BongoCreateUserInteractor)
	createUserInteractor.BongoUserRepository = repositories.NewDbBongoUserRepo(handlers)

	loginServiceHandler := services.LoginServiceHandler{}
	loginServiceHandler.BongoCreateUserInteractor = createUserInteractor
	
	serverHandler.Server.POST("/signup", loginServiceHandler.Signup)

	serverHandler.Server.Logger.Fatal(serverHandler.Server.Start(":8080"))
	//app.Server.Logger.Fatal(app.Server.StartTLS(":8080", "./config/keys/server.crt", "./config/keys/server.key"))
}
