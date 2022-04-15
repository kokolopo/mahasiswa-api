package controller

import (
	"mahasiswa-api/helper"
	"mahasiswa-api/nilai"
	"net/http"

	"github.com/labstack/echo/v4"
)

type nilaiController struct {
	service nilai.Service
}

func NewNilaiController(service nilai.Service) *nilaiController {
	return &nilaiController{service}
}

func (s *nilaiController) GetAllNilai(c echo.Context) error {
	nilais, err := s.service.GetAll()
	if err != nil {
		res := helper.ResponseFormat("Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	res := helper.ResponseFormat("All Matakuliah Data", http.StatusOK, "success", nilais)

	return c.JSON(http.StatusOK, res)
}

func (s *nilaiController) UpdateNilai(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "asdasd",
	})
}
