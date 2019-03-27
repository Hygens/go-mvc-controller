package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to
// a controller.
type Action func(rw http.ResponseWriter, r *http.Request) error

// This is our Base Controller
type AppController struct{}

// The action function helps with error handling in a controller
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

// This is the main controller type
type MyController struct {
	AppController
	*render.Render
}

// Represent the main path for REST API with JSON result
func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

// Represent other sample with REST API Data result that can be used for binary or raw data too
func (c *MyController) Home(rw http.ResponseWriter, r *http.Request) error {
	c.Data(rw, 200, []byte("Pagina Home!!!"))
	return nil
}

// Represent one sample for HTML template render result
func (c *MyController) Example(rw http.ResponseWriter, r *http.Request) error {
	c.HTML(rw, http.StatusOK, "example", nil)
	return nil
}

func main() {
	// Instance for main controller
	c := &MyController{Render: render.New(render.Options{})}

	// Server mux for simple handle routes
	mux := http.NewServeMux()

	// Various sample routes implemented
	mux.Handle("/", c.Action(c.Index))
	mux.Handle("/home", c.Action(c.Home))
	mux.Handle("/example", c.Action(c.Example))
	http.ListenAndServe(":8080", mux)
}
