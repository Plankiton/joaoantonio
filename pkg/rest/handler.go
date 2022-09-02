package rest

import (
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler struct {
	e       *echo.Echo
	Logger  echo.Logger
	SelfUrl string
}

func (h *Handler) Setup() {
	h.e.Static("/", "static")

	root := h.e.Group("/api")
	root.GET("", h.Get)
}

func (h *Handler) Start(addr string) error {
	return h.e.Start(addr)
}

type HandlerOption func(*Handler)

func New(options ...HandlerOption) *Handler {
	handler := &Handler{}
	handler.e = echo.New()

	for _, o := range options {
		o(handler)
	}

	if handler.e.Logger == nil {
		zapLogger, _ := zap.NewProduction()
		handler.e.Use(echozap.ZapLogger(zapLogger))
	}

	handler.Logger = handler.e.Logger
	handler.Setup()
	return handler
}
