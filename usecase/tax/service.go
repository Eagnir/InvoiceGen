package tax

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//Tax Service  interface
type TaxService struct {
	repo repository.DBContext
}

//NewService create new use case
func NewService(r repository.DBContext) *TaxService {
	return &TaxService{
		repo: r,
	}
}

func (s *TaxService) GetEntityById(id int) (*entity.Tax, error) {
	obj := entity.Tax{}
	db := s.repo.Context.First(&obj, id)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.Tax_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *TaxService) ListAll() ([]*entity.Tax, error) {
	var objs []*entity.Tax
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *TaxService) SaveObjectFromNew(obj *entity.Tax, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	if obj.TaxId != 0 {
		return 0, exception.Tax_PrimeryKeyNotZero
	}
	return s.SaveObject(obj)
}

func (s *TaxService) SaveObject(obj *entity.Tax) (int64, error) {
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	if obj.TaxId == 0 {
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

func (s *TaxService) DeleteByObject(obj *entity.Tax) (int64, error) {
	return s.DeleteById(obj.TaxId)
}

func (s *TaxService) DeleteById(id int) (int64, error) {
	obj, err := s.GetEntityById(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.Tax_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *TaxService) Search(queryObject entity.Tax) ([]*entity.Tax, error) {
	var objs []*entity.Tax
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
