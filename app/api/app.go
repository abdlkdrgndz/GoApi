package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"realApi/domain"
)

var customers = []domain.Customer{
	{Name: "Ahmet Bayrak", Username: "ahmetbayrak", Email: "ahmet@hotmail.com"},
	{Name: "Cemil Öz", Username: "cemiloz", Email: "cemiloz@mynet.com"},
	{Name: "Elif Boz", Username: "elifboz", Email: "elifboz@gmail.com"},
}

// Welcome /**
func Welcome(c *fiber.Ctx) error {
	return c.Status(http.StatusAccepted).JSON("Welcome to restApi")
}

// GetAllCustomer /**
func GetAllCustomer(c *fiber.Ctx) error {
	var message domain.Messages

	if len(customers) <= 0 {
		return c.Status(http.StatusNotFound).JSON(message.NoCostumerFound)
	}
	return c.Status(http.StatusOK).JSON(customers)
}

// GetByUsernameWith /**
func GetByUsernameWith(c *fiber.Ctx) error {
	var message domain.Messages
	email := c.Params("email")

	for _, elem := range customers {
		if email == elem.Email {
			return c.Status(http.StatusOK).JSON(elem)
		}
	}
	return c.Status(http.StatusNotFound).JSON(message.NotFound)
}

// CreateCustomer /**
func CreateCustomer(c *fiber.Ctx) error {
	var newCustomers domain.Customer
	var message domain.Messages

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Oops! Bad request.")
	}

	customers = append(customers, newCustomers)
	return c.Status(http.StatusOK).JSON(message.Created)

}

// UpdateCustomer /**
func UpdateCustomer(c *fiber.Ctx) error {
	var newCustomers domain.Customer
	var message domain.Messages

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Message: Bad Request")
	}

	for i, elem := range customers {
		if elem.Username == newCustomers.Username {
			elem.Name = newCustomers.Name
			elem.Email = newCustomers.Email
			customers = append(customers[:i], customers[i+1:]...)
			customers = append(customers, elem)
		}
	}

	return c.Status(http.StatusOK).JSON(message.Updated)

}

// DeleteCustomer /**
func DeleteCustomer(c *fiber.Ctx) error {
	username := c.Params("username")
	var message domain.Messages

	log.Println(username)

	for index, elem := range customers {
		if elem.Username == username {
			customers = append(customers[:index], customers[index+1:]...)
			fmt.Println(customers)
		}
	}

	return c.Status(http.StatusOK).JSON(message.Deleted)
}
