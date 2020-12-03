package database

import (
	"errors"
	"github.com/GiancarloCristiano/SeminarioGo/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase (conf *config.Config) (*sqlx.DB, error){
	switch conf.DB.Type{
		case "sqlite3":
			db, err := sqlx.Open(conf.Db.Driver, conf.Db.Conn)
			if err != nil{
				return nil, err
			}

			err = db.Ping()
			if err != nil{
				return nil, err
			}

			return db, nil
		default:
			return nil, errors.New("Invalid DB type")
	}
}