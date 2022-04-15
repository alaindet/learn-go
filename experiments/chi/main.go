package main

func main() {
	app := NewApp()

	app.SetRouterMiddleware(registerMiddleware)
	app.SetRoutes(registerRoutes)
	app.SetRouteNotFound(notFoundHandler)
	app.SetMethodNotAllowed(methodNotAllowedHandler)

	app.Start()
}
