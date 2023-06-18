package handlers

import (
	"crud-go/pkg/cars/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type CreateHandler struct {
// 	service *services.CreateService
// }

// func NewCreateHandler(service *services.CreateService) *CreateHandler {
// 	return &CreateHandler{service: service,}
// }

// CreateItem CarsCreate
// @Summary      Create a Car
// @Description  Create a new car
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        car body models.Car true "Car"
// @Success 200 {string} string "Car created successfully"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /cars [post]
func (h *Handler) CarsCreate(c *gin.Context) {

	req := models.CreateCarRequest{}
	err := c.Bind(&req)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	car := req.ToCar()
	err = h.Service.CarsCreate(car)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"car": "Car Create successfully",
	})

}