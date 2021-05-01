package invoiceItem

import "InvoiceGen/entity"

//Invoice Item UseCase interface
type InvoiceItemUseCase interface {
	GetEntityById(id int) (*entity.InvoiceItem, error)
	ListAll() ([]*entity.InvoiceItem, error)
	SaveObjectFromNew(obj *entity.InvoiceItem, er error) (int64, error)
	SaveObject(obj *entity.InvoiceItem) (int64, error)
	DeleteByObject(obj *entity.InvoiceItem) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.InvoiceItem) ([]*entity.InvoiceItem, error)
}
