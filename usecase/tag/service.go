package tag

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	"errors"

	"gorm.io/gorm"
)

//Tag Service  interface
type TagService struct {
	repo repository.DBContext
}

//NewService create new use case
func NewService(r repository.DBContext) *TagService {
	return &TagService{
		repo: r,
	}
}

func (s *TagService) GetEntityByName(name string) (*entity.Tag, error) {
	obj := entity.Tag{}
	db := s.repo.Context.First(&obj, name)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return nil, exception.Tag_RecordNotFound
		}
		return nil, db.Error
	}
	return &obj, nil
}

func (s *TagService) ListAll() ([]*entity.Tag, error) {
	var objs []*entity.Tag
	db := s.repo.Context.Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}

func (s *TagService) InsertObjectFromNew(obj *entity.Tag, er error) (int64, error) {
	if er != nil {
		return 0, er
	}
	return s.InsertObject(obj)
}

func (s *TagService) InsertObject(obj *entity.Tag) (int64, error) {
	err := obj.Validate()
	if err != nil {
		return 0, err
	}
	_, err = s.GetEntityByName(obj.Name)
	if !errors.Is(err, exception.Tag_RecordNotFound) {
		return 0, exception.Tag_NameAlreadyExist
	}
	db := s.repo.Context.Create(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *TagService) DeleteByObject(obj *entity.Tag) (int64, error) {
	return s.DeleteByName(obj.Name)
}

func (s *TagService) DeleteByName(id string) (int64, error) {
	obj, err := s.GetEntityByName(id)
	if err != nil {
		return 0, err
	}
	if obj == nil {
		return 0, exception.Tag_RecordNotFound
	}
	db := s.repo.Context.Delete(&obj)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (s *TagService) Search(queryObject entity.Tag) ([]*entity.Tag, error) {
	var objs []*entity.Tag
	db := s.repo.Context.Where(queryObject).Find(&objs)
	if db.Error != nil {
		return nil, db.Error
	}
	return objs, nil
}
