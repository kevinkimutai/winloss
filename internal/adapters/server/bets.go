package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) BetsRouter(api fiber.Router) {
	api.Post("/", s.bet.CreateBet)
	api.Get("/:betID", s.bet.GetBet)
	api.Patch("/:betID", s.bet.UpdateBet)
	api.Delete(":betID", s.bet.DeleteBet)
}
