package main

import (
	"log"
	"sdui_app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendData(c *fiber.Ctx) error {
	fields := []models.FieldProperty{}
	options := []models.Option{}

	options = append(options,
		models.Option{
			Optiontitle: "test",
			Optionvalue: 1,
		}, models.Option{
			Optiontitle: "test",
			Optionvalue: 2,
		},
	)

	fields = append(fields,
		models.FieldProperty{
			Fieldname:  "Username",
			Isshown:    true,
			Fieldtype:  "text",
			Options:    nil,
			Widgettype: "textField",
			Design: models.FieldDesign{
				Color:      "1E90FF",
				Fontsize:   1,
				Padding:    2,
				Visibility: "true",
			},
			Isenabled: false,
		}, models.FieldProperty{
			Fieldname:  "Email",
			Isshown:    true,
			Fieldtype:  "text",
			Options:    nil,
			Widgettype: "textField",
			Design: models.FieldDesign{
				Color:      "228B22",
				Fontsize:   1,
				Padding:    2,
				Visibility: "false",
			},
			Isenabled: true,
		}, models.FieldProperty{
			Fieldname:  "This is a Button",
			Isshown:    true,
			Fieldtype:  "text",
			Options:    options,
			Widgettype: "textButton",
			Design: models.FieldDesign{
				Color:      "228B22",
				Fontsize:   1,
				Padding:    2,
				Visibility: "false",
			},
			Isenabled: true,
		},
	)

	// var options map[string]interface{}
	// json.Unmarshal([]byte(`{"sample":"test"}`), &options)

	return c.JSON(
		Result{
			Code:    "00",
			Message: "Welcome to Service Driven UI.",
			Data: models.LoginModel{
				Fields:  fields,
				Version: "v1.4",
			},
		},
	)
}

func Routes(app *fiber.App) {
	app.Get("/loginFields", SendData)
}

func main() {
	// // handleRequest()
	// envRouting.LoadEnv()
	// database.Migration()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // means all user can access the API or you can just specify the user that can use the system (example: facebook.com) only
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", //header that we only accept
	}))
	Routes(app)
	log.Fatal(app.Listen("192.168.1.9:8000"))
}
