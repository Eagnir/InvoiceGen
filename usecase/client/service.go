package client

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//Client Service  interface
type ClientService struct {
	repo *repository.DBContext
}

//NewService create new use case
func NewService(r *repository.DBContext) *ClientService {
	return &ClientService{
		repo: r,
	}
}

func (s *ClientService) GetEntityById(id int) (*entity.Client, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	obj := entity.Client{}
	db := s.repo.Context.First(&obj, id)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.Client_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *ClientService) ListAll() ([]*entity.Client, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	var objs []*entity.Client
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *ClientService) ListForCompanyId(companyId int, preloads ...string) ([]*entity.Client, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	var objs []*entity.Client
	db := s.repo.Context.Joins("DefaultCurrency").Where(entity.Client{
		CompanyId: companyId,
	})
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	db = db.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *ClientService) SaveObjectFromNew(obj *entity.Client, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	if obj.ClientId != 0 {
		return 0, exception.Client_PrimeryKeyNotZero
	}
	return s.SaveObject(obj)
}

func (s *ClientService) SaveObject(obj *entity.Client) (int64, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	if obj.ClientId == 0 {
		db := s.repo.Context.Create(&obj)
		if db.Error != nil {
			return 0, db.Error
		}
		return db.RowsAffected, nil
	} else {
		db := s.repo.Context.Updates(&obj)
		if db.Error != nil {
			return 0, db.Error
		}
		return db.RowsAffected, nil
	}
}

func (s *ClientService) DeleteByObject(obj *entity.Client) (int64, error) {
	return s.DeleteById(obj.ClientId)
}

func (s *ClientService) DeleteById(id int) (int64, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	obj, err := s.GetEntityById(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.Client_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *ClientService) Search(queryObject entity.Client) ([]*entity.Client, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	var objs []*entity.Client
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
