///////////////////////////////////////////////////////////////////////////////////////////////////
// 📥 IMPORTS | CODING: UTF-8
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
    "os"

    "backend-go/app/controllers"
    "backend-go/app/database"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
// 🔶 MAIN FUNCTION
///////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://mongo:27017"
        log.Println("[AVISO] MONGO_URI não definida; usando valor padrão:", mongoURI)
    }
    database.ConnectDB(mongoURI)
    http.HandleFunc("/upload", controllers.UploadHandler)

    log.Println("Servidor rodando na porta 8081...")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
