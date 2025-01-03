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

package controllers

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "mime/multipart"
    "net/http"
    "os"
    "time"

    "backend-go/app/database"
)

type OCRResponse struct {
    Brand string `json:"brand"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// 🔶 MAIN FUNCTION
///////////////////////////////////////////////////////////////////////////////////////////////////
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Erro ao fazer parse do multipart form", http.StatusBadRequest)
        return
    }
    file, _, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Erro ao receber a imagem no campo 'image'", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // 🟢 [GENERAL] CHECK IMAGE FROM MEMORY BUFFER
    fileBuffer := &bytes.Buffer{}
    if _, err := io.Copy(fileBuffer, file); err != nil {
        http.Error(w, "Erro ao ler o conteúdo da imagem", http.StatusInternalServerError)
        return
    }

    // 🟢 [GENERAL] BUILD FORM-DATA MULTIPART TO SEND FOR OCR
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("image", "beer_can.jpg")
    if err != nil {
        http.Error(w, "Erro ao criar campo 'image' para OCR", http.StatusInternalServerError)
        return
    }
    if _, err := io.Copy(part, bytes.NewReader(fileBuffer.Bytes())); err != nil {
        http.Error(w, "Erro ao anexar imagem no form-data", http.StatusInternalServerError)
        return
    }
    writer.Close()

    // 🟢 [GENERAL] READ URL
    ocrURL := os.Getenv("OCR_SERVICE_URL")
    if ocrURL == "" {
        // Fallback
        log.Println("[AVISO] OCR_SERVICE_URL não definida; usando http://ocr-service-python:5001/process-image")
        ocrURL = "http://ocr-service-python:5001/process-image"
    }
    resp, err := http.Post(ocrURL, writer.FormDataContentType(), body)
    if err != nil {
        http.Error(w, "Erro na comunicação com o OCR Service", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        http.Error(w, "OCR Service retornou erro", resp.StatusCode)
        return
    }

    // 🟢 [GENERAL] READ JSON
    var ocrResp OCRResponse
    if err := json.NewDecoder(resp.Body).Decode(&ocrResp); err != nil {
        http.Error(w, "Erro ao decodificar JSON do OCR Service", http.StatusInternalServerError)
        return
    }

    if ocrResp.Brand == "" {
        http.Error(w, "Nenhuma marca foi identificada pelo OCR", http.StatusNotFound)
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := database.SaveToDatabase(ctx, fileBuffer.Bytes(), ocrResp.Brand); err != nil {
        http.Error(w, "Erro ao salvar no banco de dados", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Marca identificada: %s", ocrResp.Brand)
}
