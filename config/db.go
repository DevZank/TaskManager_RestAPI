package config

import (
	"database/sql"
	"fmt"
	"os"

	dotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Pega as variaveis de ambiente do arquvivo .env
// Inicializa a conexão com o banco de dados
func SetupDB() *sql.DB {
	err := dotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	dbConn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Print("Sucesso na conexão com o DB...")

	return dbConn
}
