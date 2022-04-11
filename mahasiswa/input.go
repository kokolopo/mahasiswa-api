package mahasiswa

type MhsInputRegister struct {
	NPM          int    `json:"npm"`
	NamaLengkap  string `json:"nama_lengkap"`
	Email        string `json:"email"`
	Alamat       string `json:"alamat"`
	Jurusan      string `json:"jurusan"`
	Kelas        string `json:"kelas"`
	Semester     int    `json:"semester"`
	PasswordHash string `json:"password_hash"`
}
