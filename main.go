package main
import (
	"keuangan-pribadi/route"
	"net/http"
)

func main() {
	e := route.New()

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
}