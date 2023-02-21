package api

import (
	"fmt"

	"github.com/Banana-Boat/gin-sqlc-template/internal/db"
	"github.com/Banana-Boat/gin-sqlc-template/internal/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker *util.TokenMaker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := util.NewTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user/createUser", server.createUser)
	router.GET("/user/getUser/:id", server.getUser)
	router.GET("/user/listUsers", server.listUsers)
	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
