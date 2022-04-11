package mahasiswa

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegistrasiMahasiswa(input MhsInputRegister) (Mahasiswa, error)
}

type service struct {
	repository Repository
}

func NewMahasiswaService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegistrasiMahasiswa(input MhsInputRegister) (Mahasiswa, error) {
	var mahasiswa Mahasiswa

	//enkripsi password
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.MinCost)
	if errHash != nil {
		return mahasiswa, errHash
	}

	//tangkap nilai dari inputan
	mahasiswa.NPM = input.NPM
	mahasiswa.NamaLengkap = input.NamaLengkap
	mahasiswa.Email = input.Email
	mahasiswa.Alamat = input.Alamat
	mahasiswa.Jurusan = input.Jurusan
	mahasiswa.Kelas = input.Kelas
	mahasiswa.Semester = input.Semester
	mahasiswa.PasswordHash = string(passwordHash)

	//save data yang sudah dimapping kedalam struct Mahasiswa
	mhs, err := s.repository.Save(mahasiswa)
	if err != nil {
		return mhs, err
	}

	return mhs, nil
}
