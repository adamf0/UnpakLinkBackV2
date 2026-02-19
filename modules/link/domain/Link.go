package domain

import (
	"os"
	"time"

	common "UnpakSiamida/common/domain"
	"UnpakSiamida/common/helper"

	"github.com/google/uuid"
)

type Link struct {
	common.Entity

	ID          uint       `gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID  `gorm:"type:char(36);uniqueIndex"`
	ShortUrl    string     `gorm:"type:longtext;column:short_url;"`
	LongUrl     string     `gorm:"type:longtext;column:long_url;"`
	Creator     string     `gorm:"type:longtext;column:creator;"`
	StartAccess *time.Time `gorm:""`
	EndAccess   *time.Time `gorm:""`
	Password    *string    `gorm:"type:longtext;"`
	Status      *string    `gorm:"type:longtext;column:status;"` //arsip, delete
}

func (Link) TableName() string {
	return "links"
}

// === CREATE ===
func NewLink(shortUrl string, longUrl string, password *string, start *string, end *string, creator string, isUnique bool) common.ResultValue[*Link] {
	if !isUnique {
		return common.FailureValue[*Link](NotUnique())
	}

	var startTime, endTime *time.Time
	var err error

	if start != nil {
		dc := helper.NewDateChain(*start).
			UseParseStrategy(helper.DateLayoutFirstFactory{}.CreateParser()).
			Parse()

		startTime, err = dc.Ptr()
		if err != nil {
			return common.FailureValue[*Link](InvalidFormatStart())
		}
	}
	if end != nil {
		dc := helper.NewDateChain(*end).
			UseParseStrategy(helper.DateLayoutFirstFactory{}.CreateParser()).
			Parse()

		endTime, err = dc.Ptr()
		if err != nil {
			return common.FailureValue[*Link](InvalidFormatEnd())
		}
	}

	jenisfile := &Link{
		UUID:        uuid.New(),
		ShortUrl:    shortUrl,
		LongUrl:     longUrl,
		Password:    password,
		StartAccess: startTime,
		EndAccess:   endTime,
		Creator:     creator,
		Status:      helper.StrPtr("active"),
	}

	return common.SuccessValue(jenisfile)
}

// === UPDATE ===
func Update(
	prev *Link,
	uid uuid.UUID,
	shortUrl string,
	creator string,
	isUnique bool,
) common.ResultValue[*Link] {
	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.ShortUrl = shortUrl
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	if !isUnique {
		return common.FailureValue[*Link](NotUnique())
	}

	return common.SuccessValue(prev)
}

func MoveDelete(
	prev *Link,
	uid uuid.UUID,
	creator string,
) common.ResultValue[*Link] {
	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.Status = helper.StrPtr("delete")
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	return common.SuccessValue(prev)
}

func MoveActive(
	prev *Link,
	uid uuid.UUID,
	creator string,
) common.ResultValue[*Link] {
	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.Status = helper.StrPtr("active")
		prev.StartAccess = nil
		prev.EndAccess = nil
		prev.Password = nil
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	return common.SuccessValue(prev)
}

func MoveArchive(
	prev *Link,
	uid uuid.UUID,
	creator string,
) common.ResultValue[*Link] {
	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.Status = helper.StrPtr("archive")
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	return common.SuccessValue(prev)
}

func GiveTime(
	prev *Link,
	uid uuid.UUID,
	startAccess string,
	endAccess string,
	creator string,
) common.ResultValue[*Link] {
	factory := helper.DateLayoutSecondFactory{}

	tglStart, err := helper.NewDateChain(startAccess).
		UseParseStrategy(factory.CreateParser()).
		Parse().
		Ptr()

	if err != nil {
		return common.FailureValue[*Link](InvalidFormatStart())
	}

	tglEnd, err := helper.NewDateChain(endAccess).
		UseParseStrategy(factory.CreateParser()).
		Parse().
		Ptr()

	if err != nil {
		return common.FailureValue[*Link](InvalidFormatEnd())
	}

	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.StartAccess = tglStart
		prev.EndAccess = tglEnd
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	return common.SuccessValue(prev)
}

func RollbackTime(
	prev *Link,
	uid uuid.UUID,
	creator string,
) common.ResultValue[*Link] {
	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.StartAccess = nil
		prev.EndAccess = nil
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	return common.SuccessValue(prev)
}

func GivePassword(
	prev *Link,
	uid uuid.UUID,
	password string,
	creator string,
) common.ResultValue[*Link] {
	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.Password = &password
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	return common.SuccessValue(prev)
}

func RollbackPassword(
	prev *Link,
	uid uuid.UUID,
	creator string,
) common.ResultValue[*Link] {
	if prev == nil {
		return common.FailureValue[*Link](EmptyData())
	}

	if prev.UUID != uid {
		return common.FailureValue[*Link](InvalidData())
	}

	if creator == os.Getenv("Administrator") || prev.Creator == creator {
		prev.Password = nil
	} else {
		return common.FailureValue[*Link](InvalidData())
	}

	return common.SuccessValue(prev)
}
