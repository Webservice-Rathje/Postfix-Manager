package Controller

import (
	"github.com/Webservice-Rathje/Postfix-Manager/utils"
	"github.com/gofiber/fiber/v2"
)

type createUserRequest struct {
	Secret       string `json:"secret"`
	MailPrefix   string `json:"mail_prefix"`
	Domain       string `json:"domain"`
	Password     string `json:"password"`
	Enabled      bool   `json:"enabled"`
	Alias        string `json:"alias"`
	AliasEnabled string `json:"alias_enabled"`
}

type createUserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func CreateUserController(c *fiber.Ctx) error {
	obj := createUserRequest{}
	err := c.BodyParser(obj)
	if err != nil {
		return c.JSON(createUserResponse{
			false,
			"Invalid request body",
		})
	}
	if utils.MatchSecret(obj.Secret) {

	} else {
		return c.JSON(createUserResponse{
			false,
			"Invalid Secret",
		})
	}
}
