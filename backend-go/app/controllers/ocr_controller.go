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

package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"app/services"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
// 🔶 MAIN FUNCTION
///////////////////////////////////////////////////////////////////////////////////////////////////

func OcrHandler(w http.ResponseWriter, r *http.Request) {
	// 🟢 [GENERAL] CHECK IF REQUEST METHOD IS POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// 🟢 [GENERAL] RETRIEVE FILE FROM FORM DATA
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Falha ao receber arquivo: "+err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao processar upload: %v\n", err)
		return
	}
	defer file.Close()

	// 🟢 [GENERAL] CREATE A BUFFER TO HOLD FORM DATA
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// 🟢 [GENERAL] CREATE FORM FILE PART FOR MULTIPART UPLOAD
	part, err := writer.CreateFormFile("file", "upload.jpg")
	if err != nil {
		http.Error(w, "Erro ao criar form-data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 🟢 [GENERAL] COPY UPLOADED FILE INTO FORM PART
	_, err = io.Copy(part, file)
	if err != nil {
		http.Error(w, "Erro ao copiar arquivo", http.StatusInternalServerError)
		return
	}
	writer.Close()

	// 🟢 [GENERAL] SEND FORM DATA TO OCR SERVICE
	responseData, err := services.SendToOCR(body, writer.FormDataContentType())
	if err != nil {
		http.Error(w, "Erro ao conectar ao OCR", http.StatusInternalServerError)
		return
	}

	// 🟢 [GENERAL] WRITE RESPONSE AS JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
