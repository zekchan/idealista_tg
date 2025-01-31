package idealista

type Ad struct {
	Id    string
	Price int
	Title string
	Area  int
	Rooms int
}

type Client interface {
	GetAd(id string) (Ad, error)
}

type ClientType string

const (
	ScrapeClientType ClientType = "scrape"
)

func NewClient(clientType ClientType) Client {
	switch clientType {
	case ScrapeClientType:
		return &ScrapeClient{}
	default:
		return nil
	}
}
