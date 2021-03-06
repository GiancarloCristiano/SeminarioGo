package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/GiancarloCristiano/SeminarioGo/internal/config"
	"github.com/GiancarloCristiano/SeminarioGo/internal/database"
	"github.com/GiancarloCristiano/SeminarioGo/internal/service/music"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx" 
)


func main() {
	cfg := readConfig()
	
	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
 	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := music.New(db, cfg)
	httpService := music.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config{
	configFile := flag.String("config","./config.yaml","This is the Service configuration")
	flag.Parse() 


	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
	
		fmt.Println(err.Error())
		os.Exit(1) 
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS music (
		id integer primary key autoincrement,
		name varchar,
		artist varchar,
		year integer
		);`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	
	return nil
} 