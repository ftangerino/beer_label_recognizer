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

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"app/database"
	"go.mongodb.org/mongo-driver/bson"
	"mime/multipart"
)

// 🟢 [GENERAL] SEND IMAGE TO OCR SERVICE (FLASK BACKEND)
func SendToOCR(body bytes.Buffer, contentType string) ([]byte, error) {
	// 🟢 [GENERAL] MAKE POST REQUEST TO OCR SERVICE
	resp, err := http.Post("http://ocr-service:5000/ocr", contentType, &body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 🟢 [GENERAL] READ RESPONSE DATA FROM OCR SERVICE
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 🟢 [GENERAL] PARSE JSON RESPONSE
	var result map[string]interface{}
	json.Unmarshal(responseData, &result)

	// 🟢 [GENERAL] CHECK FOR MATCH IN RESPONSE AND SAVE TO MONGO
	if match, ok := result["match"]; ok {
		saveToMongo(match.(string))
	}
	return responseData, nil
}

// 🟢 [GENERAL] SAVE OCR RESULT TO MONGODB
func saveToMongo(brand string) {
	// 🟢 [GENERAL] GET MONGODB COLLECTION REFERENCE
	collection := database.Client.Database("beerdb").Collection("beer_recognition")
	
	// 🟢 [GENERAL] INSERT OCR RESULT INTO MONGODB
	_, err := collection.InsertOne(context.TODO(), bson.M{
		"brand_name": brand,
		"created_at": time.Now(),
	})
	if err != nil {
		// 🔴 [ERROR HANDLING] LOG ERROR IF INSERT FAILS
		fmt.Println("Erro ao salvar no MongoDB:", err)
	}
}
