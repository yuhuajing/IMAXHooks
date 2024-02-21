package main

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/login", "./form.html")
	app.Post("/", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))
		//return c.SendString("POST request")
		return c.SendStatus(200)
	})

	log.Fatal(app.Listen(":3000"))
}

//func GetNFTData(c *fiber.Ctx) error {
//	chain := c.Query("chain")
//	contract := c.Query("contract")
//	collection := c.Query("collection")
//	err, res := mongodb.QueryNFTdataDB(config.Dbcollectionnftdata, chain, contract, collection)
//	if err != nil {
//		return c.Status(401).JSON(types.ErrorResponse{
//			Error:   err.Error(),
//			Success: false,
//			Data:    "",
//		})
//	}
//
//	return c.Status(200).JSON(types.ResResponse{
//		Error:   "",
//		Success: true,
//		Data:    res,
//	})
//	//return handler.SendResponse(c, res)
//}
