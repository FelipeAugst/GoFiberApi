package controllers

import (
	"api/models"
	"api/repository"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateCliente(c *fiber.Ctx) error {
	var cliente models.Cliente
	err := c.BodyParser(&cliente)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := cliente.Format(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error() + cliente.CPF})
	}

	r, err := repository.NewClienteRepo()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.Create(cliente); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cliente)
}

func ListAllClientes(c *fiber.Ctx) error {
	var clientes []models.Cliente
	r, err := repository.NewClienteRepo()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	clientes, err = r.ListAll()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(clientes)

}

func ListClientes(c *fiber.Ctx) error {

	param := c.Params("filter")
	if len(param) < 3 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.New("insira ao menos 3 letras na busca").Error()})
	}
	var clientes []models.Cliente
	r, err := repository.NewClienteRepo()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	clientes, err = r.List(param)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(clientes)

}

func ByIdCliente(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})

	}

	r, err := repository.NewClienteRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	cliente, err := r.ById(uint(id))
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	return c.JSON(cliente)

}

func EditCliente(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var cliente models.Cliente
	if err := c.BodyParser(&cliente); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := cliente.Format(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	cliente.ID = uint(id)

	r, err := repository.NewClienteRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := r.Update(cliente); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func DeleteCliente(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var cliente models.Cliente
	cliente.ID = uint(id)

	r, err := repository.NewClienteRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := r.Delete(cliente); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}
