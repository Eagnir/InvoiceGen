package adminUser

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//AdminUser Service  interface
type AdminUserService struct {
	repo repository.DBContext
}

//NewService create new use case
func NewService(r repository.DBContext) *AdminUserService {
	return &AdminUserService{
		repo: r,
	}
}

func (s *AdminUserService) GetEntityById(id int) (*entity.AdminUser, error) {
	obj := entity.AdminUser{}
	db := s.repo.Context.First(&obj, id)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.AdminUser_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *AdminUserService) ListAll() ([]*entity.AdminUser, error) {
	var objs []*entity.AdminUser
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *AdminUserService) SaveObjectFromNew(obj *entity.AdminUser, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	if obj.AdminUserId != 0 {
		return 0, exception.AdminUser_PrimeryKeyNotZero
	}
	return s.SaveObject(obj)
}

func (s *AdminUserService) SaveObject(obj *entity.AdminUser) (int64, error) {
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	if obj.AdminUserId == 0 {
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

func (s *AdminUserService) DeleteByObject(obj *entity.AdminUser) (int64, error) {
	return s.DeleteById(obj.AdminUserId)
}

func (s *AdminUserService) DeleteById(id int) (int64, error) {
	obj, err := s.GetEntityById(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.AdminUser_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *AdminUserService) Search(queryObject entity.AdminUser) ([]*entity.AdminUser, error) {
	var objs []*entity.AdminUser
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
