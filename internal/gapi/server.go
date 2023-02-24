package gapi

import (
	"fmt"

	"github.com/Banana-Boat/gin-sqlc-template/internal/db"
	gin_sqlc_template "github.com/Banana-Boat/gin-sqlc-template/internal/pb"
	"github.com/Banana-Boat/gin-sqlc-template/internal/util"
)

type Server struct {
	gin_sqlc_template.UnimplementedTestServer
	config     util.Config
	store      *db.Store
	tokenMaker *util.TokenMaker
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

	return server, nil
}
