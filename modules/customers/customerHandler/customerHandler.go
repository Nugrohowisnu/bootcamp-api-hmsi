package customerHandler

import (
	"bootcamp-api-hmsi/models"
	"bootcamp-api-hmsi/modules/customers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	UC customers.CustomerUsecase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomerUsecase) {
	handler := customerHandler{UC}

	r.GET("/customer", handler.GetAll)
	r.POST("/customer", handler.Insert)
	r.PUT("/customer/:id", handler.Update)    // Add this line
	r.DELETE("/customer/:id", handler.Delete) // Add this line
}

func (h *customerHandler) GetAll(c *gin.Context) {
	result, err := h.UC.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    result,
	})
}

func (h *customerHandler) Insert(c *gin.Context) {
	var request models.RequestInsertCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	err := h.UC.Insert(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Inserted successfully",
		"data":    []string{},
	})
}

func (h *customerHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID",
			"data":    []string{},
		})
		return
	}

	var request models.RequestUpdateCustomer
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	err = h.UC.Update(id, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated successfully",
		"data":    []string{},
	})
}

func (h *customerHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID",
			"data":    []string{},
		})
		return
	}

	err = h.UC.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted successfully",
		"data":    []string{},
	})
}
