package customerUsecase

import (
	"bootcamp-api-hmsi/models"
	"bootcamp-api-hmsi/modules/customers"
)

type customerRepository struct {
	Repo customers.CustomerRepository
}

func NewCustomerUsecase(Repo customers.CustomerRepository) customers.CustomerUsecase {
	return &customerRepository{Repo}
}

func (r *customerRepository) FindAll() (*[]models.Customers, error) {
	result, err := r.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	// TODO : logic

	return result, nil
}

func (r *customerRepository) Insert(c *models.RequestInsertCustomer) error {
	err := r.Repo.Create(c)

	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Update(id int, c *models.RequestUpdateCustomer) error {
	err := r.Repo.Update(id, c)

	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Delete(id int) error {
	err := r.Repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
