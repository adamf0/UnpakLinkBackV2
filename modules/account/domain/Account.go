package domain

type Account struct {
	ID        string      `json:"UserId" gorm:"column:userid"`
	UUID      *string     `json:"UUID"`
	Username  string      `json:"Username"`
	Password  string      `json:"-"`
	Level     string      `json:"Level"`
	Name      string      `json:"Name" gorm:"column:nama"`
	Email     string      `json:"-"`
	ExtraRole []ExtraRole `gorm:"-"; json:"ExtraRole,omitempty"`
}

func (Account) TableName() string {
	return "user"
}
