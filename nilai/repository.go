package nilai

import (
	"mahasiswa-api/mahasiswa"
	"mahasiswa-api/matakuliah"

	"gorm.io/gorm"
)

type Repository interface {
	Save(nilai Nilai) (Nilai, error)
	UpdateNilai(nilai Nilai) (Nilai, error)
	FindMhsByNpm(npm int) (mahasiswa.Mahasiswa, error)
	FindMtkByMtkName(mtkName string) (matakuliah.Matakuliah, error)
	FindMtkByMtkId(id int) (matakuliah.Matakuliah, error)
	FindAll() ([]Nilai, error)
}

type repository struct {
	db *gorm.DB
}

func NewNilaiRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nilai Nilai) (Nilai, error) {
	err := r.db.Create(&nilai).Error
	if err != nil {
		return nilai, err
	}
	return nilai, nil
}

func (r *repository) UpdateNilai(nilai Nilai) (Nilai, error) {
	err := r.db.Save(&nilai).Error
	if err != nil {
		return nilai, err
	}

	return nilai, nil
}

func (r *repository) FindMhsByNpm(npm int) (mahasiswa.Mahasiswa, error) {
	var mahasiswa mahasiswa.Mahasiswa
	err := r.db.Where("npm = ?", npm).Find(&mahasiswa).Error
	if err != nil {
		return mahasiswa, err
	}

	return mahasiswa, nil
}

func (r *repository) FindMtkByMtkName(mtkName string) (matakuliah.Matakuliah, error) {
	var matakuliah matakuliah.Matakuliah

	err := r.db.Where("nama_matakuliah = ?", mtkName).Find(&matakuliah).Error
	if err != nil {
		return matakuliah, err
	}

	return matakuliah, nil
}

func (r *repository) FindMtkByMtkId(id int) (matakuliah.Matakuliah, error) {
	var matakuliah matakuliah.Matakuliah

	err := r.db.Where("id = ?", id).Find(&matakuliah).Error
	if err != nil {
		return matakuliah, err
	}

	return matakuliah, nil
}

func (r *repository) FindAll() ([]Nilai, error) {
	var nilai []Nilai

	err := r.db.Find(&nilai).Error
	if err != nil {
		return nilai, err
	}

	return nilai, nil
}
