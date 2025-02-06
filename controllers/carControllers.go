package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID int `json:"car_id"`
	Brand string `json:"brand"`
	Price string `json:"price"`
}

var DataCar = []Car {
	{CarID: 1, Brand: "Toyota", Price: "100.000.000"},
}

// POST
func CreateCar(ctx *gin.Context) {
	var addCar Car

	// Casting terlebih dahulu dan disimpan didalam variabel dan menggunakan pointer
	// pada variabel yang menggunakan pointer akan mengubah pada asal value-nya
	err := ctx.ShouldBindJSON(&addCar)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	addCar.CarID = len(DataCar) + 1
	DataCar = append(DataCar, addCar)

	// mengirim data yang telah di request, lalu mengirimkan response ke client
	ctx.JSON(http.StatusCreated, gin.H{
		"car" : addCar,
	})
}


// GET

func GetCar(ctx *gin.Context) {
	// membuat param untuk bisa diambil data sesuai dengan ID (param sesuai dengan atribut struct)
	CarID := ctx.Param("CarID")
	IDConvCar, err := strconv.Atoi(CarID) // mengubah ke int

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error_Status":  "Bad Request",
			"Error_Message": "Invalid CarID format",
		})
		return
	}
	

	condition := false
	var Data Car

	// Membuat perulangan range dari data yang sudah ditambahkan atau get data
	for i, car := range DataCar {
		if IDConvCar == car.CarID {
			condition = true
			Data = DataCar[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Error_Status" : "",
			"Error_Message" : fmt.Sprintf("Car With ID %v Not Found", CarID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": Data,
	})
}

// Update
func UpdateCar(ctx *gin.Context) {
	GetCarID := ctx.Param("CarID")
	carID, err := strconv.Atoi(GetCarID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var CarUpdated Car
	condition := false

	if err := ctx.ShouldBindBodyWithJSON(&CarUpdated); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_message": "Failed",
			"error_message":  err.Error(),
		})
		return
	}

	for i, car := range DataCar {
		if carID == car.CarID {
			condition = true
			DataCar[i].Brand = CarUpdated.Brand
			DataCar[i].Price = CarUpdated.Price
			DataCar[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"status_message" : "",
			"error_message" : "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("Successfuly id %v : ", carID),
	})
}


func DeleteCar(ctx *gin.Context) {
	GetCarID := ctx.Param("CarID")
	carID, err := strconv.Atoi(GetCarID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	condition := false
	var Index int

	for i, car := range DataCar {
		if carID == car.CarID {
			condition = true
			Index = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_message" : "Not Found",
			"error_message" : "Error",
		})
	}

	copy(DataCar[Index:], DataCar[Index+1:])
	DataCar[len(DataCar)-1] = Car{}
	DataCar = DataCar[:len(DataCar)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("Successfully id %v ", carID),
	})

}