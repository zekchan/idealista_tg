package bot

import (
	"regexp"
	"time"

	"idealista_tg/internal/config"
	"idealista_tg/pkg/idealista"

	tele "gopkg.in/telebot.v4"
)

type Service struct {
	bot              *tele.Bot
	idealistaClient  idealista.Client
	config           *config.Config
	idealistaAdRegex *regexp.Regexp
}

func NewService(cfg *config.Config, client idealista.Client) *Service {
	return &Service{
		config:           cfg,
		idealistaClient:  client,
		idealistaAdRegex: regexp.MustCompile(`.*/imovel/(\d+).*`),
	}
}

func (s *Service) Start() error {
	pref := tele.Settings{
		Token:  s.config.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return err
	}
	s.bot = bot

	s.registerHandlers()

	s.bot.Start()
	return nil
}
