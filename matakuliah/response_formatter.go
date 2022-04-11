package matakuliah

type MtkDataResponse struct {
	NamaMatakuliah string `json:"nama_matakuliah"`
	Sks            int    `json:"sks"`
	Dosen          string `json:"dosen"`
}

func ConvertResponse(dirtyData Matakuliah) MtkDataResponse {
	cleanData := MtkDataResponse{
		NamaMatakuliah: dirtyData.NamaMatakuliah,
		Sks:            dirtyData.Sks,
		Dosen:          dirtyData.Dosen,
	}

	return cleanData
}
