package mahasiswa

import "gorm.io/gorm"

type Repository interface {
	Save(mahasiswa Mahasiswa) (Mahasiswa, error)
}

type repository struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(mahasiswa Mahasiswa) (Mahasiswa, error) {
	err := r.db.Create(&mahasiswa).Error
	if err != nil {
		return mahasiswa, err
	}

	return mahasiswa, nil
}
