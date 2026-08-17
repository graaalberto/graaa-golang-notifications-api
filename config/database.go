package config

import (
	"fmt"
	"log"
	"os"

	"frotago-notifications-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "fleet_db"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Africa/Luanda",
		host, user, password, dbname, port, sslmode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("❌ Falha crítica ao conectar ao PostgreSQL (%s) para Notificações: %v", dbname, err)
	}

	log.Printf("✅ Notificações: Conexão com o PostgreSQL [%s:%s/%s] estabelecida!", host, port, dbname)

	// Cria tabela de histórico de notificações automaticamente no fleet_db
	_ = DB.AutoMigrate(&models.NotificationLog{})
}
