package nilai

import "gorm.io/gorm"

type Repository interface {
	Save(nilai Nilai) (Nilai, error)
}

type repository struct {
	db *gorm.DB
}

func NewNilaiRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nilai Nilai) (Nilai, error) {
	return nilai, nil
}
