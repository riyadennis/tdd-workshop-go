package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

type DBConnection struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}
type Config struct {
	DB DBConnection
}

func CreateConnection(c *Config) {
	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		c.DB.User, c.DB.Pass, c.DB.Name)
	dbConn, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Println("failed opening database:", err)
	}
	fmt.Println(dbConn)
}

func AnotherFunc(c *Config) {
	fmt.Println(c.DB)
}

func NewConfig(options ...func(*Config)) {
	c := &Config{
		DBConnection{"localhost", "8080", "dbName", "User", "Pass"},
	}
	for _, option := range options {
		option(c)
	}

}

func main() {
	NewConfig(CreateConnection)
	NewConfig(AnotherFunc)
}
