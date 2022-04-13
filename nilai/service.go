package nilai

type Service interface {
	Save(input InputNilai) (Nilai, error)
	UpdateNilai(id int) (Nilai, error)
	GetAll() ([]NilaiDataResponse, error)
}

type service struct {
	repository Repository
}

func NewNilaiService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Save(input InputNilai) (Nilai, error) {
	var nilai Nilai

	mhs, _ := s.repository.FindMhsByNpm(input.Npm)
	mtk, _ := s.repository.FindMtkByMtkName(input.NamaMtk)
	// if err != nil {
	// 	return mhs, err
	// }

	nilai.NPMMahasiswa = mhs.NPM
	nilai.IdMatakuliah = mtk.Id
	nilai.Nilai = input.Nilai

	newNilai, err := s.repository.Save(nilai)
	if err != nil {
		return newNilai, err
	}

	return newNilai, nil
}

func (s *service) GetAll() ([]NilaiDataResponse, error) {
	nilais, err := s.repository.FindAll()

	var sliceNilai []NilaiDataResponse
	var sliceNamaMhs []string
	var slcieNamaMtk []string
	if err != nil {
		return sliceNilai, err
	}

	for i := range nilais {
		mhs, _ := s.repository.FindMhsByNpm(nilais[i].NPMMahasiswa)

		sliceNamaMhs = append(sliceNamaMhs, mhs.NamaLengkap)
	}

	for i := range nilais {
		mhs, _ := s.repository.FindMtkByMtkId(nilais[i].IdMatakuliah)

		slcieNamaMtk = append(slcieNamaMtk, mhs.NamaMatakuliah)
	}

	for i := range nilais {
		sliceNilai = append(sliceNilai, NilaiDataResponse{NamaLengkap: sliceNamaMhs[i], NamaMtk: slcieNamaMtk[i], Nilai: nilais[i].Nilai})
	}

	return sliceNilai, nil
}
