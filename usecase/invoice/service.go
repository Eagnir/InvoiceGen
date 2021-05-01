package invoice

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//Invoice Service  interface
type InvoiceService struct {
	repo repository.DBContext
}

//NewService create new use case
func NewService(r repository.DBContext) *InvoiceService {
	return &InvoiceService{
		repo: r,
	}
}

func (s *InvoiceService) GetEntityById(id int) (*entity.Invoice, error) {
	obj := entity.Invoice{}
	db := s.repo.Context.First(&obj, id)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.Invoice_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *InvoiceService) ListAll() ([]*entity.Invoice, error) {
	var objs []*entity.Invoice
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *InvoiceService) SaveObjectFromNew(obj *entity.Invoice, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	if obj.InvoiceId != 0 {
		return 0, exception.Invoice_PrimeryKeyNotZero
	}
	return s.SaveObject(obj)
}

func (s *InvoiceService) SaveObject(obj *entity.Invoice) (int64, error) {
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	if obj.InvoiceId == 0 {
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

func (s *InvoiceService) DeleteByObject(obj *entity.Invoice) (int64, error) {
	return s.DeleteById(obj.InvoiceId)
}

func (s *InvoiceService) DeleteById(id int) (int64, error) {
	obj, err := s.GetEntityById(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.Invoice_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *InvoiceService) Search(queryObject entity.Invoice) ([]*entity.Invoice, error) {
	var objs []*entity.Invoice
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
