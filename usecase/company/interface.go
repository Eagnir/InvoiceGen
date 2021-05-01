package company

import "InvoiceGen/entity"

//Company UseCase interface
type CompanyUseCase interface {
	GetEntityById(id int) (*entity.Company, error)
	ListAll() ([]*entity.Company, error)
	SaveObjectFromNew(obj *entity.Company, er error) (int64, error)
	SaveObject(obj *entity.Company) (int64, error)
	DeleteByObject(obj *entity.Company) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.Company) ([]*entity.Company, error)
}
