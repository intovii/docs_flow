package server

import (
	"context"
	"fmt"
	"net/http"
	"zavad/config"
	"zavad/internal/delivery/http/middleware"
	// "zavad/internal/entities"
	"zavad/internal/usecase"
	protos "zavad/pkg/proto/gen/go"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const statusOK = `OK`

type Server struct {
	logger  *zap.Logger
	cfg     *config.ConfigModel
	serv    *gin.Engine
	Usecase *usecase.Usecase
	mdlware *middleware.Middleware
}

func NewServer(logger *zap.Logger, cfg *config.ConfigModel, uc *usecase.Usecase, mdlware *middleware.Middleware) (*Server, error) {
	return &Server{
		logger:  logger,
		cfg:     cfg,
		serv:    gin.Default(),
		Usecase: uc,
		mdlware: mdlware,
	}, nil
}

func (s *Server) OnStart(_ context.Context) error {
	s.createController()
	go func() {
		s.logger.Debug("serv started")
		if err := s.serv.Run(s.cfg.Server.Host + ":" + s.cfg.Server.Port); err != nil {
			s.logger.Error("failed to serve: " + err.Error())
		}
		return
	}()
	return nil
}

func (s *Server) OnStop(_ context.Context) error {
	s.logger.Debug("stop grps")
	//s.serv.GracefulStop()
	return nil
}

func (s *Server) GetUserToken(ctx *gin.Context) {
	request := protos.GetUserTokenRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("failed to unmarshar request: %v", err)})
		return
	}

	token, err := s.Usecase.GetUserToken(ctx, &request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to get token: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, &token)
	return
}

func (s *Server) CreateUser(ctx *gin.Context) {
	// request := YClients.ClientRegistration{}
	// err := ctx.ShouldBindJSON(&request)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("failed to unmarshar request: %v", err)})
	// 	return
	// }
	// user, err := s.Usecase.CreateUser(ctx, &request)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to create user: %v", err)})
	// 	return
	// }

	// ctx.JSON(http.StatusCreated, user)
}

