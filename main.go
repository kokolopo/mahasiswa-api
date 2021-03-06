package main

import (
	"mahasiswa-api/controller"
	"mahasiswa-api/mahasiswa"
	"mahasiswa-api/matakuliah"
	"mahasiswa-api/nilai"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//koneksi ke database
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

	//domain nilai
	nilaiRepository := nilai.NewNilaiRepository(db)
	nilaiService := nilai.NewNilaiService(nilaiRepository)
	nilaiController := controller.NewNilaiController(nilaiService)

	e := echo.New()
	api := e.Group("/api/v1")

	//route mahasiswa
	api.POST("/mahasiswa", mhsController.RegisterMhs)

	//route matakuliah
	api.POST("/matakuliah", mtkController.CreateNew)
	api.GET("/matakuliah", mtkController.GetAll)
	api.GET("/matakuliah/:id", mtkController.GetById)

	//route nilai
	api.GET("/nilai", nilaiController.GetAllNilai)

	e.Logger.Fatal(e.Start(":8080"))

}
