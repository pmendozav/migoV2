package main

import (
	"migoV2/infrastructure/server"
	"migoV2/infrastructure/db"
	"migoV2/interfaces/repositories"
	"migoV2/interfaces/services"
	"migoV2/usecases/login"
	"migoV2/usecases/common"
)

func main() {
	serverHandler := server.NewServerHandler()
	dbHandler := db.NewOrmHandler()

	handlers := make(map[string] repositories.DbHandler)
    handlers["DbBongoUserRepo"] = dbHandler
    handlers["DbBongoCommonRepo"] = dbHandler

	//login service
	createUserInteractor := new(login.BongoUserInteractor)
	createUserInteractor.BongoUserRepository = repositories.NewDbBongoUserRepo(handlers)
	loginServiceHandler := services.LoginServiceHandler{}
	loginServiceHandler.BongoUserInteractor = createUserInteractor
	

	//common service
	createCommonInteractor := new(common.BongoCommonInteractor)
	createCommonInteractor.BongoCommonRepository = repositories.NewDbBongoCommonRepo(handlers)
	commonServiceHandler := services.CommonServiceHandler{}
	commonServiceHandler.BongoCommonInteractor = createCommonInteractor


	serverHandler.Server.POST("/signup", loginServiceHandler.Signup)
	serverHandler.Server.POST("/login", loginServiceHandler.SignIn)

	serverHandler.Server_r.GET("/partner/:wildcard", commonServiceHandler.FindPartnersWithWildcard)

	// serverHandler.Server.Logger.Fatal(serverHandler.Server.Start(":8080"))
	serverHandler.Server.Logger.Fatal(serverHandler.Server.StartTLS(":8080", "./config/keys/server.crt", "./config/keys/server.key"))
}
