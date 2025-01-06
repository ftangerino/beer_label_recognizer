///////////////////////////////////////////////////////////////////////////////////////////////////
// 📅 IMPORTS | CODING: UTF-8
///////////////////////////////////////////////////////////////////////////////////////////////////
// ✅ → Discussed and realized
// 🟢 → Discussed and not realized (to be done after the meeting)
// 🟡 → Little important and not discussed (unhindered)
// 🔴 → Very important and not discussed (hindered)
// ❌ → Canceled
// ⚪ → Postponed (technical debit)
///////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"log"
	"net/http"
	"app/routes"
	"app/database"
)

func main() {
	// 🟢 [GENERAL] INITIALIZE MONGODB CONNECTION
	database.InitMongo()

	// 🟢 [GENERAL] SET UP API ROUTES
	r := routes.SetupRoutes()

	// 🟢 [GENERAL] DEFINE SERVER PORT
	port := ":5001"
	log.Printf("API rodando na porta %s\n", port)

	// 🟢 [GENERAL] START HTTP SERVER
	err := http.ListenAndServe(port, r)
	if err != nil {
		// 🔴 [ERROR HANDLING] LOG ERROR IF SERVER FAILS TO START
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}
}
