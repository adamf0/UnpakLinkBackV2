package infrastructure

import (
	create "UnpakSiamida/modules/click/application/CreateClick"
	getAll "UnpakSiamida/modules/click/application/GetAllClicks"
	domainClick "UnpakSiamida/modules/click/domain"

	"github.com/mehdihadeli/go-mediatr"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "fmt"
)

func RegisterModuleClick(db *gorm.DB) error {
	repoClick := NewClickRepository(db)

	// Register request handler
	mediatr.RegisterRequestHandler[
		create.CreateClickCommand,
		uint,
	](&create.CreateClickCommandHandler{
		Repo: repoClick,
	})

	mediatr.RegisterRequestHandler[
		getAll.GetAllClicksQuery,
		domainClick.PagedClicks,
	](&getAll.GetAllClicksQueryHandler{
		Repo: repoClick,
	})

	return nil
}
