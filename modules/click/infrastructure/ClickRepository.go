package infrastructure

import (
	commondomain "UnpakSiamida/common/domain"
	domainClick "UnpakSiamida/modules/click/domain"
	"context"
	"strings"

	"gorm.io/gorm"
)

type ClickRepository struct {
	db *gorm.DB
}

func NewClickRepository(db *gorm.DB) domainClick.IClickRepository {
	return &ClickRepository{db: db}
}

func (r *ClickRepository) Create(
	ctx context.Context,
	click *domainClick.Click,
) error {

	err := r.db.WithContext(ctx).
		Create(click).Error

	if err != nil {
		return err
	}

	return nil
}

var allowedSearchColumns = map[string]string{
	// key:param -> db column
	"short_url": "l.short_url",
	"long_url":  "l.long_url",
	"creator":   "l.creator",
}

// ------------------------
// GET ALL
// ------------------------
func (r *ClickRepository) GetAll(
	ctx context.Context,
	search string,
	searchFilters []commondomain.SearchFilter,
	page, limit *int,
) ([]domainClick.ClickDefault, int64, error) {

	var rows = make([]domainClick.ClickDefault, 0)
	var total int64

	db := r.db.WithContext(ctx).
		Table("clicks c").
		Select(`
			c.id AS ID,
			l.uuid AS LinkUUID,
			l.id AS LinkId,
			l.short_url as LinkShort,
			l.long_url as LinkLong,

			c.ip as IP,
			c.ip_client as IPClient,
			c.iso_code as ISO,
			c.country as Country,
			c.user_agent as UserAgent,
			c.created_at as Created
		`).
		Joins(`
			JOIN links l 
				ON c.link_id = l.id
		`)

	// -----------------------------------
	// ADVANCED FILTERS
	// -----------------------------------
	for _, f := range searchFilters {
		col, ok := allowedSearchColumns[strings.ToLower(f.Field)]
		if !ok {
			continue
		}

		val := ""
		if f.Value != nil {
			val = strings.TrimSpace(*f.Value)
		}

		switch strings.ToLower(f.Operator) {
		case "eq":
			db = db.Where(col+" = ?", val)
		case "neq":
			db = db.Where(col+" <> ?", val)
		case "like":
			db = db.Where(col+" LIKE ?", "%"+val+"%")
		case "in":
			db = db.Where(col+" IN ?", strings.Split(val, ","))
		}
	}

	// -----------------------------------
	// GLOBAL SEARCH
	// -----------------------------------
	if strings.TrimSpace(search) != "" {
		like := "%" + search + "%"
		var or []string
		var args []interface{}

		for _, col := range allowedSearchColumns {
			or = append(or, col+" LIKE ?")
			args = append(args, like)
		}

		db = db.Where("("+strings.Join(or, " OR ")+")", args...)
	}

	// -----------------------------------
	// COUNT
	// -----------------------------------
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// -----------------------------------
	// ORDER + PAGINATION
	// -----------------------------------
	db = db.Order("c.id DESC")

	if page != nil && limit != nil && *limit > 0 {
		offset := (*page - 1) * (*limit)
		db = db.Offset(offset).Limit(*limit)
	}

	// -----------------------------------
	// EXECUTE
	// -----------------------------------
	if err := db.Find(&rows).Error; err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}
