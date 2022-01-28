package main

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	//"html"
	//"hubuc/config"
	//"hubuc/helper"
)

func main() {

	//get root directory
	//resourcesPath := helper.GetRoot()

	//engine.Layout("content")
	app := fiber.New()

	app.Post("user", User)

	//initialize database

	if err := app.Listen(":7500"); err != nil {
		panic(err)
	}
}

type userHolder struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func userRepository() []userHolder {

	repo := make([]userHolder, 5)

	//insert userholder into repo
	holder := userHolder{
		Username: "John Doe",
		Email:    "jd@gmail.com",
		Password: "123",
	}

	repo[0] = holder

	holder = userHolder{
		Username: "Femi John",
		Email:    "fj@gmail.com",
		Password: "123",
	}

	repo[1] = holder

	holder = userHolder{
		Username: "Femi Ball",
		Email:    "fb@gmail.com",
		Password: "123",
	}

	repo[2] = holder

	holder = userHolder{
		Username: "Temi Ball",
		Email:    "tb@gmail.com",
		Password: "123",
	}

	repo[3] = holder

	holder = userHolder{
		Username: "Kola Ball",
		Email:    "kb@gmail.com",
		Password: "123",
	}

	repo[4] = holder

	return repo
}

func checkRepository(email string) bool {

	repo := userRepository()

	if len(repo) > 0 {
		for _, item := range repo {
			if email == item.Email {
				return true
			}
		}

		return false
	}

	return false
}

func (s userHolder) Validate() error {
	valid := validation.ValidateStruct(&s,
		validation.Field(&s.Username, validation.Required),
		validation.Field(&s.Email, validation.Required),
		validation.Field(&s.Password, validation.Required),
	)

	//extra validation for password
	passLen := len(s.Password)
	if (passLen < 6) || (passLen > 120) {
		return errors.New("password length must be between 6 to 120 characters")
	}

	//validate email against repo
	emailFound := checkRepository(s.Email)
	if emailFound {
		return errors.New("email exist, please use a new email address")
	}

	return valid
}
func User(c *fiber.Ctx) error {

	//create a fake repositor

	var input userHolder
	if err := c.BodyParser(&input); err != nil {
		return response(c, err, "one or more of the fields not formatted properly", false, 400)
	}

	if err := input.Validate(); err != nil {
		return response(c, err, err.Error(), false, 400)
	}

	//code accepts valid date and processes it
	//returns a randomy generated data
	userID, err := uuid.NewV4()
	if err != nil {
		return response(c, "", "unable to generate user id, please try again", false, 500)
	}

	output := fiber.Map{
		"user_id": userID.String(),
	}

	return response(c, output, "success", true, 200)

}

type respHolder struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func response(c *fiber.Ctx, data interface{}, message string, status bool, code int) error {
	var response respHolder
	response.Status = status
	response.Data = data
	response.Message = message

	return c.Status(code).JSON(response)
}
