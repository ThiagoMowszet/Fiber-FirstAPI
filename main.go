package main

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)




type User struct {
    Id string
    Firstname string
    Lastname string
}    


func handleUser(c *fiber.Ctx) error {
    user := User {
        Firstname: "Joe",
        Lastname: "Doe",
    }
    return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
    user := User{}
    if error := c.BodyParser(&user); error != nil {
        return error
    }

    user.Id = uuid.NewString()
    return c.Status(fiber.StatusOK).JSON(user)
}


func main() {

    // Default config
    app := fiber.New(fiber.Config{
        AppName: "First API with Fiber",
        CaseSensitive: true,
    })

    // Middlewares
    app.Use(logger.New())

    // Basic Hello World in Fiber -> http://localhost:3000
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })



    // Parameters -> http://localhost:3000/abc
    // app.Get("/:value", func(c *fiber.Ctx) error {
    //     return c.SendString("value: " + c.Params("value"))
    // })


    // Optional Parameter ->  http://localhost:3000/thiago
    // app.Get("/:name?", func(c *fiber.Ctx) error {
    //     if c.Params("name") != "" {
    //         return c.SendString("Hello " + c.Params("name"))
    //     }
    //     return c.SendString("Where is John?")
    // })




    // Wildcarts -> http://localhost:3000/api/a/b/c
    app.Get("/api/*", func(c *fiber.Ctx) error {
        return c.SendString("API Path: " + c.Params("*"))
    })

    
    // Basic Routes
    // app.Get("/users", handleUser)
    // app.Post("/users", handleCreateUs)


    app.Use(requestid.New())

    // Group Routes
    userGroup := app.Group("/users")

    userGroup.Post("", handleCreateUser)
    userGroup.Get("", handleUser)



    // Port listen 
    app.Listen(":3000")

}
