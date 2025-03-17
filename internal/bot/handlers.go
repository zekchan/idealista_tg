package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func (s *Service) registerHandlers() {
	s.bot.Handle("/start", s.handleStart)
	s.bot.Handle("/list", s.handleList)
	s.bot.Handle(tele.OnText, s.handleText)
}

func (s *Service) handleStart(c tele.Context) error {
	return c.Send("Hello! I'm your new Telegram bot.")
}

func (s *Service) handleList(c tele.Context) error {
	ads, err := s.storage.GetAds()
	if err != nil {
		return c.Send("Error retrieving ads: " + err.Error())
	}

	if len(ads) == 0 {
		return c.Send("No ads found in storage.")
	}

	message := "📋 *List of Saved Properties:*\n\n"
	for i, ad := range ads {
		adLink := fmt.Sprintf("https://idealista.com/imovel/%s", ad.Id)
		message += fmt.Sprintf("%d. [Property %s](%s)\n", i+1, ad.Id, adLink)
	}

	return c.Send(message, &tele.SendOptions{
		ParseMode:             tele.ModeMarkdown,
		DisableWebPagePreview: true,
	})
}

func (s *Service) handleText(c tele.Context) error {
	text := c.Message().Text
	matches := s.idealistaAdRegex.FindStringSubmatch(text)
	if len(matches) < 2 {
		return c.Send("Please send me an Idealista property link")
	}
	adID := matches[1]
	if exists, err := s.storage.HasAd(adID); err != nil || exists {
		return c.Reply("Ad already exists")
	}
	// Get the ad details
	ad, err := s.idealistaClient.GetAd(adID)
	if err != nil {
		return c.Reply("Error fetching ad details")
	}
	s.storage.SaveAd(&ad, c.Sender().Username)
	return c.Reply(fmt.Sprintf("Found ad with ID: %s, price: %d, area: %d, rooms: %s", ad.Id, ad.Price, ad.Area, ad.Rooms))
}
