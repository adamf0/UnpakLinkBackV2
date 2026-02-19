package application

type CreateClickCommand struct {
	LinkId    uint
	IP        string
	IPClient  string
	ISO       string
	Country   string
	Referer   string
	UserAgent string
}
