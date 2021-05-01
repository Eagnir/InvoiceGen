package company

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//Company Service  interface
type CompanyService struct {
	repo repository.DBContext
}

//NewService create new use case
func NewService(r repository.DBContext) *CompanyService {
	return &CompanyService{
		repo: r,
	}
}

func (s *CompanyService) GetEntityById(id int) (*entity.Company, error) {
	obj := entity.Company{}
	db := s.repo.Context.First(&obj, id)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.Company_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *CompanyService) ListAll() ([]*entity.Company, error) {
	var objs []*entity.Company
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *CompanyService) SaveObjectFromNew(obj *entity.Company, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	if obj.CompanyId != 0 {
		return 0, exception.Company_PrimeryKeyNotZero
	}
	return s.SaveObject(obj)
}

func (s *CompanyService) SaveObject(obj *entity.Company) (int64, error) {
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	if obj.CompanyId == 0 {
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

func (s *CompanyService) DeleteByObject(obj *entity.Company) (int64, error) {
	return s.DeleteById(obj.CompanyId)
}

func (s *CompanyService) DeleteById(id int) (int64, error) {
	obj, err := s.GetEntityById(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.Company_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *CompanyService) Search(queryObject entity.Company) ([]*entity.Company, error) {
	var objs []*entity.Company
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
