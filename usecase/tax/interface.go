package tax

import "InvoiceGen/entity"

//Tax UseCase interface
type TaxUseCase interface {
	GetEntityById(id int) (*entity.Tax, error)
	ListAll() ([]*entity.Tax, error)
	SaveObjectFromNew(obj *entity.Tax, er error) (int64, error)
	SaveObject(obj *entity.Tax) (int64, error)
	DeleteByObject(obj *entity.Tax) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.Tax) ([]*entity.Tax, error)
}
