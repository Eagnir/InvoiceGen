package taxGroup

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//TaxGroup Service  interface
type TaxGroupService struct {
	repo repository.DBContext
}

//NewService create new use case
func NewService(r repository.DBContext) *TaxGroupService {
	return &TaxGroupService{
		repo: r,
	}
}

func (s *TaxGroupService) GetEntityById(id int) (*entity.TaxGroup, error) {
	obj := entity.TaxGroup{}
	db := s.repo.Context.First(&obj, id)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.TaxGroup_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *TaxGroupService) ListAll() ([]*entity.TaxGroup, error) {
	var objs []*entity.TaxGroup
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *TaxGroupService) SaveObjectFromNew(obj *entity.TaxGroup, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	if obj.TaxGroupId != 0 {
		return 0, exception.TaxGroup_PrimeryKeyNotZero
	}
	return s.SaveObject(obj)
}

func (s *TaxGroupService) SaveObject(obj *entity.TaxGroup) (int64, error) {
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	if obj.TaxGroupId == 0 {
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

func (s *TaxGroupService) DeleteByObject(obj *entity.TaxGroup) (int64, error) {
	return s.DeleteById(obj.TaxGroupId)
}

func (s *TaxGroupService) DeleteById(id int) (int64, error) {
	obj, err := s.GetEntityById(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.TaxGroup_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *TaxGroupService) Search(queryObject entity.TaxGroup) ([]*entity.TaxGroup, error) {
	var objs []*entity.TaxGroup
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
