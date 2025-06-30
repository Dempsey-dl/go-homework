package main

import (
	"blog/internal/config"
	"blog/internal/router"
	"blog/pkg/database"

	"log"
)

func main() {
	//加载配置
	yamlpath := "configs/config.yaml"
	cfg, err := config.Load(yamlpath)
	if err != nil {
		log.Fatalln("failed to load yaml")
	}
	//初始化数据库
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalln("failed to init database")
	}
	//初始化路由
	r := router.SetupRouter(cfg, db)
	//启动服务器
	log.Printf("Server is running on %s in %s mode", cfg.Server.Address, cfg.Server.Mode)
	if err := r.Run(cfg.Server.Address); err != nil {
		log.Fatalf("failed to start serviec %v", cfg.Server.Address)
	}
}
