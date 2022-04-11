package matakuliah

import "gorm.io/gorm"

type Repository interface {
	Save(matkul Matakuliah) (Matakuliah, error)
	FindAll() ([]Matakuliah, error)
	FindById(id int) (Matakuliah, error)
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

func (r *repository) FindAll() ([]Matakuliah, error) {
	var mtk []Matakuliah

	err := r.db.Find(&mtk).Error
	if err != nil {
		return mtk, err
	}

	return mtk, nil
}

func (r *repository) FindById(id int) (Matakuliah, error) {
	var mtk Matakuliah

	err := r.db.Where("id = ?", id).Find(&mtk).Error
	if err != nil {
		return mtk, err
	}

	return mtk, nil
}
