package matakuliah

import "time"

type Matakuliah struct {
	Id             int       `json:"id"`
	NamaMatakuliah string    `json:"nama_matakuliah"`
	Sks            int       `json:"sks"`
	Dosen          string    `json:"dosen"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
