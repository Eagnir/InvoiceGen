package adminUser

import "InvoiceGen/entity"

//AdminUser UseCase interface
type AdminUserUseCase interface {
	GetEntityById(id int) (*entity.AdminUser, error)
	ListAll() ([]*entity.AdminUser, error)
	SaveObjectFromNew(obj *entity.AdminUser, er error) (int64, error)
	SaveObject(obj *entity.AdminUser) (int64, error)
	DeleteByObject(obj *entity.AdminUser) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.AdminUser) ([]*entity.AdminUser, error)
}
