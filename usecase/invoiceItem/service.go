package invoiceItem

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//InvoiceItem Service  interface
type InvoiceItemService struct {
	repo repository.DBContext
}

//NewService create new use case
func NewService(r repository.DBContext) *InvoiceItemService {
	return &InvoiceItemService{
		repo: r,
	}
}

func (s *InvoiceItemService) GetEntityById(id int) (*entity.InvoiceItem, error) {
	obj := entity.InvoiceItem{}
	db := s.repo.Context.First(&obj, id)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.InvoiceItem_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *InvoiceItemService) ListAll() ([]*entity.InvoiceItem, error) {
	var objs []*entity.InvoiceItem
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *InvoiceItemService) SaveObjectFromNew(obj *entity.InvoiceItem, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	if obj.InvoiceItemId != 0 {
		return 0, exception.InvoiceItem_PrimeryKeyNotZero
	}
	return s.SaveObject(obj)
}

func (s *InvoiceItemService) SaveObject(obj *entity.InvoiceItem) (int64, error) {
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	if obj.InvoiceItemId == 0 {
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

func (s *InvoiceItemService) DeleteByObject(obj *entity.InvoiceItem) (int64, error) {
	return s.DeleteById(obj.InvoiceItemId)
}

func (s *InvoiceItemService) DeleteById(id int) (int64, error) {
	obj, err := s.GetEntityById(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.InvoiceItem_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *InvoiceItemService) Search(queryObject entity.InvoiceItem) ([]*entity.InvoiceItem, error) {
	var objs []*entity.InvoiceItem
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
