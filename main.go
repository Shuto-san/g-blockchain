package main

import (
	"net/http"
	"os"

	"github.com/Shuto-san/g-blockchain/blockchain"
	"github.com/labstack/echo/v4"
)

var bc = blockchain.NewBlockchain()

func main() {
	HTTPPort := "3001"
	if len(os.Args) > 1 && os.Args[1] != "" {
		HTTPPort = os.Args[1]
	}

	e := echo.New()
	e.GET("/blocks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, bc.Chain)
	})
	e.POST("/mine", func(c echo.Context) (err error) {
		b := new(blockchain.Block)
		if err = c.Bind(b); err != nil {
			return
		}
		block := bc.AddBlock(b.Data)
		return c.JSON(http.StatusOK, block)
	})
	e.Logger.Fatal(e.Start(":" + HTTPPort))
}
