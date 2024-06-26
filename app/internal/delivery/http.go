package delivery

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
	"github.com/mikeyQwn/server-ping/config"
	"github.com/mikeyQwn/server-ping/internal"
	"github.com/mikeyQwn/server-ping/internal/delivery/templates"
	"github.com/mikeyQwn/server-ping/internal/models"
	"github.com/mikeyQwn/server-ping/pkg/logger"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	log logger.Logger
	uc  internal.Usecase
	cfg config.Config

	e *echo.Echo
}

func New(log logger.Logger, uc internal.Usecase, cfg config.Config) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	if cfg.TLS.Enabled {
		e.AutoTLSManager.Cache = autocert.DirCache(cfg.TLS.CacheDirectory)
	}

	return &Server{
		log: log,
		uc:  uc,
		cfg: cfg,

		e: e,
	}
}

func (s *Server) Run(addr string) error {
	if s.cfg.TLS.Enabled {
		return s.e.StartAutoTLS(addr)
	}

	return s.e.Start(addr)
}

func (s *Server) MapHandlers() {
	s.e.GET("/", s.handleIndex)
	s.e.GET("/ping", s.handlePing)
	s.e.GET("/healthcheck", s.handleHealthcheck)
	s.e.GET("/poll", s.handlePoll)
	static := s.e.Group("/static", s.cacheMiddleware)
	static.Static("", "internal/delivery/static")
}

func (s *Server) handleIndex(ctx echo.Context) error {
	index := s.getIndex()
	pollable := templates.Pollable(index, "/poll", 5)
	layout := templates.Layout(serviceName, pollable)
	return Render(ctx, http.StatusOK, layout)
}

func (s *Server) handlePing(ctx echo.Context) error {
	if err := s.uc.CheckIsOnline(s.cfg.PingAddress); err != nil {
		return ctx.JSON(http.StatusOK, models.StatusResponse{IsOnline: false, Error: err})
	}
	return ctx.JSON(http.StatusOK, models.StatusResponse{IsOnline: true})
}

func (s *Server) handleHealthcheck(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "OK")
}

func (s *Server) handlePoll(ctx echo.Context) error {
	pollable := templates.Pollable(s.getIndex(), "/poll", 5)
	return Render(ctx, http.StatusOK, pollable)
}

func (s *Server) getIndex() templ.Component {

	err := s.uc.CheckIsOnline(s.cfg.PingAddress)
	addr := s.cfg.PingAddress
	if s.cfg.HidePort {
		split := strings.Split(s.cfg.PingAddress, ":")
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
	return index
}

func (s *Server) cacheMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		const MAX_AGE = time.Hour * 96
		HEADER_TEMPLATE := fmt.Sprintf("max-age=%d", int(MAX_AGE.Seconds()))
		c.Response().Header().Set("Cache-Control", HEADER_TEMPLATE)
		return next(c)
	}
}
