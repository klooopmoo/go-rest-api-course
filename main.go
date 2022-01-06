package main

import (
	"fmt"
	transportHTTP "github.com/klooopmoo/go-rest-api-course/internal/transport/http"
	"net/http"
	//"net/http"
)

//struct which contains things like pointers
// to data connections
type App struct{}

//Run sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our API ")
	handler := transportHTTP.NewHandler()
	handler.SetUpRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed to set up server ")
		return err
	}
	return nil
}
func main() {
	fmt.Println("Go Rest API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our Rest API")
		fmt.Println(err)
	}
}
