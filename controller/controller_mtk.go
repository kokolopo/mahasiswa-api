package controller

import (
	"mahasiswa-api/helper"
	"mahasiswa-api/matakuliah"
	"net/http"

	"github.com/labstack/echo/v4"
)

type mtkController struct {
	service matakuliah.Service
}

func NewMatakuliahController(service matakuliah.Service) *mtkController {
	return &mtkController{service}
}

func (h *mtkController) CreateNew(c echo.Context) error {
	var input matakuliah.InputMatakuliah

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newData, errNew := h.service.CreateNew(input)
	if errNew != nil {
		return c.JSON(http.StatusBadRequest, errNew)
	}

	convData := matakuliah.ConvertResponse(newData)
	response := helper.ResponseFormat("input berhasil", http.StatusCreated, "success", convData)

	return c.JSON(http.StatusCreated, response)
}
