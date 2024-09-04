package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type BetHandlerPort interface {
	CreateBet(c *fiber.Ctx) error
	GetBet(c *fiber.Ctx) error
	UpdateBet(c *fiber.Ctx) error
	DeleteBet(c *fiber.Ctx) error
}

type StatsHandlerPort interface {
}

type ServerAdapter struct {
	port  string
	bet   BetHandlerPort
	stats StatsHandlerPort
}

func NewServer(port string, bet BetHandlerPort, stats StatsHandlerPort) *ServerAdapter {
	return &ServerAdapter{
		port:  port,
		bet:   bet,
		stats: stats,
	}
}

func (s *ServerAdapter) StartServer() {
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	// Define routes
	app.Route("/api/v1/bets", s.BetsRouter)
	app.Route("/api/v1/stats", s.StatsRouter)

}
