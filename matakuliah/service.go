package matakuliah

type Service interface {
	CreateNew(input InputMatakuliah) (Matakuliah, error)
	GetAll() ([]Matakuliah, error)
	GetById(id int) (Matakuliah, error)
}

type service struct {
	repository Repository
}

func NewMatkulService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateNew(input InputMatakuliah) (Matakuliah, error) {
	var matkul Matakuliah

	matkul.NamaMatakuliah = input.NamaMatakuliah
	matkul.Sks = input.Sks
	matkul.Dosen = input.Dosen

	newMatkul, err := s.repository.Save(matkul)
	if err != nil {
		return matkul, err
	}

	return newMatkul, nil
}

func (s *service) GetAll() ([]Matakuliah, error) {
	matkuls, err := s.repository.FindAll()
	if err != nil {
		return matkuls, err
	}

	return matkuls, nil
}

func (s *service) GetById(id int) (Matakuliah, error) {
	matkul, err := s.repository.FindById(id)
	if err != nil {
		return matkul, err
	}

	return matkul, nil
}
