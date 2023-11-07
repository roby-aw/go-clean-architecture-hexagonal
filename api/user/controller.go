package user

import (
	"net/http"

	userBusiness "github.com/roby-aw/go-clean-architecture-hexagonal/business/user"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service userBusiness.Service
}

func NewController(service userBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         account
// @Accept       json
// @Produce      json
// @Param	     AuthLogin body user.AuthLogin true "Login"
// @Success      200 {object} user.AuthLogin
// @Failure      400 {object} user.AuthLogin
// @Router 		/user/login [post]
func (Controller *Controller) Login(c *fiber.Ctx) error {
	var auth userBusiness.AuthLogin
	if err := c.BodyParser(&auth); err != nil {
		return err
	}
	res, err := Controller.service.Login(auth)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    res.Token.AccessToken,
		MaxAge:   res.Token.AccessTokenExpired,
		HTTPOnly: true,
		Secure:   true,
	})
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success login",
		"result":  res,
	})
}

func (Controller *Controller) Register(c *fiber.Ctx) error {
	var data userBusiness.Register
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	err := Controller.service.RegisterUser(data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success register",
	})
}

func (Controller *Controller) Logout(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success logout",
	})
}

func (Controller *Controller) GetMe(c *fiber.Ctx) error {
	id := c.Locals("id").(string)
	res, err := Controller.service.GetUserByID(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid token",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success get data",
		"result":  res,
	})
}
