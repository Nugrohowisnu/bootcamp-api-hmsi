package customers

import "bootcamp-api-hmsi/models"

type (
	CustomerRepository interface {
		GetAll() (*[]models.Customers, error)
		Create(c *models.RequestInsertCustomer) error
		Update(id int, c *models.RequestUpdateCustomer) error
		Delete(id int) error
	}

	CustomerUsecase interface {
		FindAll() (*[]models.Customers, error)
		Insert(c *models.RequestInsertCustomer) error
		Update(id int, c *models.RequestUpdateCustomer) error
		Delete(id int) error
	}
)
