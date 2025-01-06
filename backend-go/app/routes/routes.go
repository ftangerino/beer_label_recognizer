package routes

import (
	"net/http"
	"app/controllers"
)

func SetupRoutes() *http.ServeMux {
	// 🟢 [GENERAL] CREATE A NEW HTTP SERVER MULTIPLEXER
	r := http.NewServeMux()
	
	// 🟢 [GENERAL] REGISTER OCR HANDLER FOR UPLOAD ROUTE
	r.HandleFunc("/upload", controllers.OcrHandler)
	
	// 🟢 [GENERAL] RETURN CONFIGURED ROUTER
	return r
}
