package delivery

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/mikeyQwn/server-ping/internal"
	"github.com/mikeyQwn/server-ping/internal/delivery/templates"
	"github.com/mikeyQwn/server-ping/internal/models"
	"github.com/mikeyQwn/server-ping/pkg/logger"
)

type Server struct {
	log         logger.Logger
	uc          internal.Usecase
	pingAddress string
	hidePort    bool

	e *echo.Echo
}

func New(log logger.Logger, uc internal.Usecase, pingAddress string, hidePort bool) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	return &Server{
		log:         log,
		uc:          uc,
		pingAddress: pingAddress,
		hidePort:    hidePort,

		e: e,
	}
}

func (s *Server) Run(addr string) error {
	return s.e.Start(addr)
}

func (s *Server) MapHandlers() {
	s.e.GET("/", s.handleIndex)
	s.e.GET("/ping", s.handlePing)
	s.e.GET("/healthcheck", s.handleHealthcheck)
	s.e.Static("/static", "internal/delivery/static")
}

func (s *Server) handleIndex(ctx echo.Context) error {
	err := s.uc.CheckIsOnline(s.pingAddress)
	addr := s.pingAddress
	if s.hidePort {
		split := strings.Split(s.pingAddress, ":")
		if len(split) > 0 {
			addr = split[0]
		}
	}
	statusMsg := statusUPMsg
	statusDescription := upDescriptionMsg
	color := colorGreen

	if err != nil {
		statusMsg = statusDOWNMsg
		statusDescription = fmt.Sprintf(connectionErrorFormatMsg, err.Error())
		color = colorRed
	}

	index := templates.Index(serviceName, statusMsg, statusDescription, addr, color)
	layout := templates.Layout(serviceName, index)
	return Render(ctx, http.StatusOK, layout)
}

func (s *Server) handlePing(ctx echo.Context) error {
	if err := s.uc.CheckIsOnline(s.pingAddress); err != nil {
		return ctx.JSON(http.StatusOK, models.StatusResponse{IsOnline: false, Error: err})
	}
	return ctx.JSON(http.StatusOK, models.StatusResponse{IsOnline: true})
}

func (s *Server) handleHealthcheck(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "OK")
}
