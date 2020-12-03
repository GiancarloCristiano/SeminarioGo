package music

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method string
	path string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func (s httpService) Register(r *gin.Engine){
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

func makeEndpoints(s Service) []*endpoint {
	arr := []*endpoint{}
	arr = append(list, 
		&endpoint{
			method: "GET",
			path: "/music",
			function: getAll(s),
		},
		&endpoint{
			method: "GET",
			path: "/music/:id",
			function: getMusic(s),
		},
		&endpoint{
			method: "POST",
			path: "/music",
			function: postMusic(s),
		},
		&endpoint{
			method: "DELETE",
			path: "/music/:id",
			function: deleteMusic(s),
		},
		&endpoint{
			method: "PUT",
			path: "/music",
			function: putMusic(s),
		},
	)
	return arr
}





func getAll(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		result, err := s.ReadAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"music": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"music": result,
			})
		}
	}
}

func getMusic(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.ReadMusic(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"music": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"music": result,
			})
		}
	}	
}		


func postMusic(s Service) gin.HandlerFunc {
	var music Music
	return func(c *gin.Context) {
		c.BindJSON(&music)
		result, err := s.CreateMusic(music)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"music": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"music": result,
			})
		}
	}
}

func deleteMusic(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.RemoveMusic(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"music": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"music": result,
			})
		}
	}
}




func putMusic(s Service) gin.HandlerFunc {
	var music Music
	return func(c *gin.Context) {
		c.BindJSON(&music)
		result, err := s.UpdateMusic(music)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"music": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"music": result,
			})
		}
	}
}