package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"neco-wallet-center/internal/model"
	"neco-wallet-center/internal/server"
	"neco-wallet-center/internal/utils"
	"net"
	"os"
)

func main() {
	config := &utils.Config{}
	config, err := utils.GetConfig("config.dev.yaml")
	if err != nil {
		return
	}

	db, err := model.InitDB(config)
	if err != nil {
		_ = fmt.Errorf("connect error")
		return
	}
	migration(db)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("start gRPC server")
	grpcServer := server.NewGrpcServer()
	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatal("Launch gRPC server failed.")
	}
}

func migration(db *gorm.DB) {
	_ = db.AutoMigrate(model.ERC20TokenData{})
	_ = db.AutoMigrate(model.ERC1155TokenData{})
	_ = db.AutoMigrate(model.FishingWallet{})
	_ = db.AutoMigrate(model.ERC20WalletLog{})
	_ = db.AutoMigrate(model.ERC1155WalletLog{})
}

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}
