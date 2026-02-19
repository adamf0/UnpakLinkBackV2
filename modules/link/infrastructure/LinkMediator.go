package infrastructure

import (
	commoninfra "UnpakSiamida/common/infrastructure"

	create "UnpakSiamida/modules/link/application/CreateLink"
	delete "UnpakSiamida/modules/link/application/DeleteLink"
	getAll "UnpakSiamida/modules/link/application/GetAllLinks"
	get "UnpakSiamida/modules/link/application/GetLink"
	getdefault "UnpakSiamida/modules/link/application/GetLinkDefault"
	setupUuid "UnpakSiamida/modules/link/application/SetupUuidLink"

	givePassword "UnpakSiamida/modules/link/application/GivePassword"
	moveLink "UnpakSiamida/modules/link/application/MoveLink"
	rollbackPassword "UnpakSiamida/modules/link/application/RollbackPassword"
	rollbackTime "UnpakSiamida/modules/link/application/RollbackTime"
	timeLink "UnpakSiamida/modules/link/application/TimeLink"
	update "UnpakSiamida/modules/link/application/UpdateLink"
	domainLink "UnpakSiamida/modules/link/domain"
	eventLink "UnpakSiamida/modules/link/event"

	"github.com/mehdihadeli/go-mediatr"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "fmt"
)

func RegisterModuleLink(db *gorm.DB) error {
	// dsn := "root:@tcp(127.0.0.1:3306)/unpak_link?charset=utf8mb4&parseTime=true&loc=Local"

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	return fmt.Errorf("Indikator Renstra DB connection failed: %w", err)
	// 	// panic(err)
	// }

	repoLink := NewLinkRepository(db)
	// if err := db.AutoMigrate(&domainLink.Link{}); err != nil {
	// 	panic(err)
	// }

	// Pipeline behavior
	// mediatr.RegisterRequestPipelineBehaviors(NewValidationBehaviorLink())

	// Register request handler
	mediatr.RegisterRequestHandler[
		create.CreateLinkCommand,
		string,
	](&create.CreateLinkCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		update.UpdateLinkCommand,
		string,
	](&update.UpdateLinkCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		moveLink.MoveLinkCommand,
		string,
	](&moveLink.MoveLinkCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		givePassword.GivePasswordCommand,
		string,
	](&givePassword.GivePasswordCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		timeLink.TimeLinkCommand,
		string,
	](&timeLink.TimeLinkCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		rollbackTime.RollbackTimeCommand,
		string,
	](&rollbackTime.RollbackTimeCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		rollbackPassword.RollbackPasswordCommand,
		string,
	](&rollbackPassword.RollbackPasswordCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		moveLink.MoveLinkCommand,
		string,
	](&moveLink.MoveLinkCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		delete.DeleteLinkCommand,
		string,
	](&delete.DeleteLinkCommandHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		getdefault.GetLinkDefaultByShortQuery,
		*domainLink.LinkDefault,
	](&getdefault.GetLinkDefaultByShortQueryHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		get.GetLinkByUuidQuery,
		*domainLink.Link,
	](&get.GetLinkByUuidQueryHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		getAll.GetAllLinksQuery,
		domainLink.PagedLinks,
	](&getAll.GetAllLinksQueryHandler{
		Repo: repoLink,
	})

	mediatr.RegisterRequestHandler[
		setupUuid.SetupUuidLinkCommand,
		string,
	](&setupUuid.SetupUuidLinkCommandHandler{
		Repo: repoLink,
	})

	commoninfra.RegisterDomainEvent(&eventLink.LinkCountEvent{})

	mediatr.RegisterNotificationHandler[eventLink.LinkCountEvent](
		eventLink.NewLinkCountEventHandler(db),
	)

	return nil
}
