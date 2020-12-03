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
	CreateMusic(Music) (Music, error)
	ReadAll() ([]*Music, error)
	ReadMusic(int) (*Music, error)
	UpdateMusic(Music) (bool, error)
	RemoveMusic(int) (bool, error)
} 

type service struct {
	db *sqlx.DB
	conf *config.Config
}

func New(db *sqlx.DB, c *config.Config) (Service, error){
	return service{db,c}, nil
}


func (s service) CreateMusic(m Music) (Music, error)  { 
	query := "INSERT INTO music (name, artist, year) VALUES (?,?, ?)" 
	res, err := s.db.Exec(query, m.Name, m.Artist, m.Year)
	if err != nil{
		 return m, err
	}
	m.ID,_ = res.LastInsertId()
	fmt.Println(res.LastInsertId())
	return m, nil
}


func (s service) ReadAll() ([]*Music, error) {
	var arr []*Music
	query := "SELECT * FROM music"
	if err := s.db.Select(&list,query); 
	err != nil {
		return nil, err
	}
		return arr, nil
}


func (s service) ReadMusic(ID int) (*Music, error) {
	var music Music
	query := "SELECT * FROM music WHERE ID = ?"
	if err := s.db.Get(&music, query, ID); 
	err != nil {
		return nil, err
	}
	return &music, nil
}


func (s service) UpdateMusic(m Music) (bool,error)  { 
	query := "UPDATE music SET Name = ?, Artist = ?, Year = ?, WHERE ID = ? " 
	_, err := s.db.Exec(query, m.Name, m.Artist, m.Year, m.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}


func (s service) RemoveMusic(ID int) (bool,error)   {
	query := "DELETE FROM music WHERE ID = ?"
	_,err := s.db.Exec(query, ID) 
	if err != nil {
		return false, err
	}
	return true, nil
}
