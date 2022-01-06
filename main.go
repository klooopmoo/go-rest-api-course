package main

import "fmt"

//struct which contains things like pointers
// to data connections
type App struct{}

//Run sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our API ")
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
