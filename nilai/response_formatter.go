package nilai

type NilaiDataResponse struct {
	NamaLengkap string `json:"nama_lengkap"`
	NamaMtk     string `json:"matakuliah"`
	Nilai       int    `json:"nilai"`
}

// func ConvertResponse(dirtyData InputNilai) NilaiDataResponse {
// 	cleanData := NilaiDataResponse{
// 		Npm:     dirtyData.Npm,
// 		NamaMtk: dirtyData.NamaMtk,
// 		Nilai:   dirtyData.Nilai,
// 	}

// 	return cleanData
// }
