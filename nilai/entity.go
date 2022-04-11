package nilai

type Nilai struct {
	Id           int `json:"id"`
	NPMMahasiswa int `json:"npm_mahasiswa"`
	IdMatakuliah int `json:"id_matakuliah"`
	Nilai        int `json:"nilai"`
}
