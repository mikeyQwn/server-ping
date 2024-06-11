package delivery

import (
	"net/http"

	"github.com/labstack/echo"
	models "github.com/mikeyQwn/server-ping"
	"github.com/mikeyQwn/server-ping/internal"
	"github.com/mikeyQwn/server-ping/pkg/logger"
)

type Server struct {
	log         logger.Logger
	uc          internal.Usecase
	pingAddress string

	e *echo.Echo
}

func New(log logger.Logger, uc internal.Usecase, pingAddress string) *Server {
	e := echo.New()
	return &Server{
		log:         log,
		uc:          uc,
		pingAddress: pingAddress,

		e: e,
	}
}

func (s *Server) Run(addr string) error {
	return s.e.Start(addr)
}

func (s *Server) MapHandlers() {
	s.e.GET("/ping", s.pingHandler)
	s.e.GET("/healthcheck", s.healthcheckHandler)
}

func (s *Server) pingHandler(ctx echo.Context) error {
	if err := s.uc.CheckIsOnline(s.pingAddress); err != nil {
		return ctx.JSON(http.StatusOK, models.StatusResponse{IsOnline: false, Error: err})
	}
	return ctx.JSON(http.StatusOK, models.StatusResponse{IsOnline: true})
}

func (s *Server) healthcheckHandler(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "OK")
}
