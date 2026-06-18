module broker

go 1.26.1

require (
	common v0.0.0
	github.com/go-chi/chi/v5 v5.3.0
	github.com/go-chi/cors v1.2.2
)

replace common => ../common
