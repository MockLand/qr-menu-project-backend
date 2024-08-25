package config

import "fmt"


const (
	DBUser 	   	= "zarokewinda"
	DBPassword 	= "password"
	DBName   	= "dashboard"
	DBPort   	= "5432"
	DBHost   	= "0.0.0.0"
	DBType 		= "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPassword, DBName)
}