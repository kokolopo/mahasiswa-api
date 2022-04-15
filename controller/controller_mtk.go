package controller

import (
	"mahasiswa-api/helper"
	"mahasiswa-api/matakuliah"
	"net/http"
	"strconv"

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

func (h *mtkController) GetAll(c echo.Context) error {
	matkuls, err := h.service.GetAll()
	if err != nil {
		res := helper.ResponseFormat("Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	// wadah response formatter
	var formatResponse []interface{}

	for _, v := range matkuls {
		formatResponse = append(formatResponse, matakuliah.ConvertResponse(v))
	}

	res := helper.ResponseFormat("All Matakuliah Data", http.StatusOK, "success", formatResponse)

	return c.JSON(http.StatusCreated, res)
}

func (h *mtkController) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	data, errData := h.service.GetById(id)
	if errData != nil {
		res := helper.ResponseFormat("Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	if data.Id == 0 {
		res := helper.ResponseFormat("Not Found", http.StatusBadRequest, "failed", err)

		return c.JSON(http.StatusBadRequest, res)
	}

	convData := matakuliah.ConvertResponse(data)
	res := helper.ResponseFormat("Matakuliah Data", http.StatusOK, "success", convData)

	return c.JSON(http.StatusOK, res)
}
