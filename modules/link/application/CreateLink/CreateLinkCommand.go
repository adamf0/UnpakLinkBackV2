package application

type CreateLinkCommand struct {
	ShortUrl string
	LongUrl  string
	Password *string
	Start    *string
	End      *string
	Creator  string
}
