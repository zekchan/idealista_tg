package bot

import (
	tele "gopkg.in/telebot.v4"
)

func (s *Service) registerHandlers() {
	s.bot.Handle("/start", s.handleStart)
}

func (s *Service) handleStart(c tele.Context) error {
	return c.Send("Hello! I'm your new Telegram bot.")
}
