package matakuliah

type Service interface {
	CreateNew(input InputMatakuliah) (Matakuliah, error)
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
