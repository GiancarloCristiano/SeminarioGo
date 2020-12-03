package music

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/GiancarloCristiano/SeminarioGo/internal/config"
)

type Music struct {
	ID int64
	Name string
	Artist string
	Year int64
}

type Service interface {
	AddMusic(Music) (Music, error)
	FindAll() ([]*Music, error)
	FindByID(int) (*Music, error)
	updateMusic(Music) (bool, error)
	RemoveByID(int) (bool, error)
} 

type service struct {
	db *sqlx.DB
	conf *config.Config
}

func New(db *sqlx.DB,c *config.Config) (Service, error){
	return service{db,c}, nil
}

func (s service) AddMusic(m Music) (Music,error)  { 
	query := "INSERT INTO music (name, artist, year) VALUES (?,?, ?)" 
	res, err := s.db.Exec(query, m.Name, m.Artist, m.Year)
	if err != nil{
		 return m, err
	}
	m.ID,_ = res.LastInsertId()
	fmt.Println(res.LastInsertId())
	return m, nil
}