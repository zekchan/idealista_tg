package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func (s *Service) registerHandlers() {
	s.bot.Handle("/start", s.handleStart)
	s.bot.Handle(tele.OnText, s.handleText)
}

func (s *Service) handleStart(c tele.Context) error {
	return c.Send("Hello! I'm your new Telegram bot.")
}

func (s *Service) handleText(c tele.Context) error {
	text := c.Message().Text()

	matches := s.idealistaAdRegex.FindStringSubmatch(text)
	if len(matches) < 2 {
		return c.Send("Please send me an Idealista property link")
	}
	adID := matches[1]
	fmt.Println(matches)

	// Get the ad details
	ad, err := s.idealistaClient.GetAd(adID)
	if err != nil {
		return c.Send("Error fetching ad details")
	}
	return c.Send(fmt.Sprintf("Found ad with ID: %s", ad.Id))
}
