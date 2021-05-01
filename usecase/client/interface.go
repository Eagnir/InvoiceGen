package client

import "InvoiceGen/entity"

//Client UseCase interface
type ClientUseCase interface {
	GetEntityById(id int) (*entity.Client, error)
	ListAll() ([]*entity.Client, error)
	SaveObjectFromNew(obj *entity.Client, er error) (int64, error)
	SaveObject(obj *entity.Client) (int64, error)
	DeleteByObject(obj *entity.Client) (int64, error)
	DeleteById(id int) (int64, error)
	Search(queryObject entity.Client) ([]*entity.Client, error)
}
