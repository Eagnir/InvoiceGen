package adminUser

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"InvoiceGen/interface/web/api/setting"
	"errors"

	"gorm.io/gorm"
)

//AdminUser Service  interface
type AdminUserService struct {
	repo *repository.DBContext
}

//NewService create new use case
func NewService(r *repository.DBContext) *AdminUserService {
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

func (s *AdminUserService) VerifyCredential(email string, password string) (*entity.AdminUser, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	obj := entity.AdminUser{}
	db := s.repo.Context.Preload("Company").Where(&entity.AdminUser{Email: email, Password: password}).First(&obj)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.AdminUser_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *AdminUserService) GenerateAuthToken(user *entity.AdminUser) error {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	//obj := entity.AdminUser{}
	db := s.repo.Context.Where(&entity.AdminUser{AdminUserId: user.AdminUserId}).First(user)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return exception.AdminUser_RecordNotFound
		}
		return db.Error
	}
	uuid := entity.NewUUID()
	user.AuthToken = &uuid
	db = s.repo.Context.Save(user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *AdminUserService) VerifyAuthTokenAndEmail(token *entity.UUID, email string) (*entity.AdminUser, error) {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	obj := entity.AdminUser{}
	db := s.repo.Context.Preload("Company").Where(&entity.AdminUser{AuthToken: token, Email: email}).First(&obj)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.AdminUser_InvalidAuthToken
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *AdminUserService) ChangePassword(adminUser *entity.AdminUser, password string) error {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	adminUser.Password = password
	adminUser.AuthToken = nil
	db := s.repo.Context.Save(adminUser)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func (s *AdminUserService) ResetPassword(email string) error {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	obj := entity.AdminUser{}
	db := s.repo.Context.Where(&entity.AdminUser{Email: email}).First(&obj)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return exception.AdminUser_RecordNotFound
		}
		return db.Error
	}
	token, err := entity.StringToUUID(setting.APIResetToken)
	if err != nil {
		return err
	}
	obj.AuthToken = &token
	db = s.repo.Context.Save(&obj)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *AdminUserService) InvalidateAuthToken(adminUser *entity.AdminUser) error {
	s.repo.OpenContext()
	defer s.repo.CloseContext()
	adminUser.AuthToken = nil
	db := s.repo.Context.Save(adminUser)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
