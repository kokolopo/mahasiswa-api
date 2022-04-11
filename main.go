package main

import (
	"mahasiswa-api/controller"
	"mahasiswa-api/mahasiswa"
	"mahasiswa-api/matakuliah"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/kuliah?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("tidak ada koneksi DB")
	}
	// fmt.Println("berhasil koneksi ke database")

	// domain mahasiswa
	mhsRepository := mahasiswa.NewMahasiswaRepository(db)
	mhsService := mahasiswa.NewMahasiswaService(mhsRepository)
	mhsController := controller.NewMahasiswaController(mhsService)

	// domain matakuliah
	mtkRepository := matakuliah.NewMatkulRepository(db)
	mtkService := matakuliah.NewMatkulService(mtkRepository)
	mtkController := controller.NewMatakuliahController(mtkService)

	e := echo.New()
	api := e.Group("/api/v1")
	api.POST("/mahasiswa", mhsController.RegisterMhs)
	api.POST("/matakuliah", mtkController.CreateNew)

	e.Logger.Fatal(e.Start(":8080"))
}
