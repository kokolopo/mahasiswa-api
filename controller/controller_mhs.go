package controller

import (
	"mahasiswa-api/helper"
	"mahasiswa-api/mahasiswa"
	"net/http"

	"github.com/labstack/echo/v4"
)

type mhsController struct {
	service mahasiswa.Service
}

func NewMahasiswaController(service mahasiswa.Service) *mhsController {
	return &mhsController{service}
}

func (h *mhsController) RegisterMhs(c echo.Context) error {
	var inpMhs mahasiswa.MhsInputRegister

	err := c.Bind(&inpMhs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newMhs, errNew := h.service.RegistrasiMahasiswa(inpMhs)
	if errNew != nil {
		return c.JSON(http.StatusUnprocessableEntity, errNew)
	}

	convData := mahasiswa.ConvertResponse(newMhs)

	response := helper.ResponseFormat("registrasi berhasil", http.StatusCreated, "success", convData)

	return c.JSON(http.StatusCreated, response)
}
