package matakuliah

import "gorm.io/gorm"

type Repository interface {
	Save(matkul Matakuliah) (Matakuliah, error)
}

type repository struct {
	db *gorm.DB
}

func NewMatkulRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(matkul Matakuliah) (Matakuliah, error) {
	err := r.db.Create(&matkul).Error
	if err != nil {
		return matkul, err
	}
	return matkul, nil
}
