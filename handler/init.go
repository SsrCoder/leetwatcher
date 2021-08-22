// Package handler is used to register handlers for router
package handler

import (
	"github.com/SsrCoder/leetwatcher/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	g *gin.Engine
	s *service.Service
}

func New(g *gin.Engine, s *service.Service) *Handler {
	return &Handler{g: g, s: s}
}
