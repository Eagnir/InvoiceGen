package invoice

import "InvoiceGen/entity"

//Invoice UseCase interface
type InvoiceUseCase interface {
	GetEntityById(id int) (*entity.Invoice, error)
	ListAll() ([]*entity.Invoice, error)
	SaveObjectFromNew(obj *entity.Invoice, er error) (int64, error)
	SaveObject(obj *entity.Invoice) (int64, error)
	DeleteByObject(obj *entity.Invoice) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.Invoice) ([]*entity.Invoice, error)
}
