package tag

import "InvoiceGen/entity"

//Tag UseCase interface
type TagUseCase interface {
	GetEntityById(id int) (*entity.Tag, error)
	ListAll() ([]*entity.Tag, error)
	SaveObjectFromNew(obj *entity.Tag, er error) (int64, error)
	SaveObject(obj *entity.Tag) (int64, error)
	DeleteByObject(obj *entity.Tag) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.Tag) ([]*entity.Tag, error)
}
