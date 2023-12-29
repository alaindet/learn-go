package app

// TODO: Move to .env
func initConfig() map[string]interface{} {
	return map[string]interface{}{

		// App
		"APP_NAME": "Gin Books",
		"APP_PORT": "4000",

		// Database
		"DATABASE_USER":     "demo",
		"DATABASE_PASSWORD": "demo",
		"DATABASE_HOST":     "localhost",
		"DATABASE_PORT":     "3306",
		"DATABASE_NAME":     "demo",
	}
}
