package taxGroup

import "InvoiceGen/entity"

//Tax Group UseCase interface
type TaxGroupUseCase interface {
	GetEntityById(id int) (*entity.TaxGroup, error)
	ListAll() ([]*entity.TaxGroup, error)
	SaveObjectFromNew(obj *entity.TaxGroup, er error) (int64, error)
	SaveObject(obj *entity.TaxGroup) (int64, error)
	DeleteByObject(obj *entity.TaxGroup) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.TaxGroup) ([]*entity.TaxGroup, error)
}
