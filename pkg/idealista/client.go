package idealista

type Ad struct {
	Id string
}

type Client interface {
	GetAd(id string) (Ad, error)
}

type ScrapeClient struct{}

func (c *ScrapeClient) GetAd(id string) (Ad, error) {
	return Ad{Id: id}, nil
}

type ClientType string

const (
	ScrapeClientType ClientType = "scrape"
)

// NewClient creates a new Idealista client with optional configurations
func NewClient(clientType ClientType) Client {
	switch clientType {
	case ScrapeClientType:
		return &ScrapeClient{}
	default:
		return nil
	}
}
