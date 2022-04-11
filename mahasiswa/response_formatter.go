package mahasiswa

type MhsDataResponse struct {
	NPM         int    `json:"npm"`
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	Jurusan     string `json:"jurusan"`
}

func ConvertResponse(dirtyData Mahasiswa) MhsDataResponse {
	cleanData := MhsDataResponse{
		NPM:         dirtyData.NPM,
		NamaLengkap: dirtyData.NamaLengkap,
		Email:       dirtyData.Email,
		Jurusan:     dirtyData.Jurusan,
	}

	return cleanData
}
