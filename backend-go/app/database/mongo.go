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

package database

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

///////////////////////////////////////////////////////////////////////////////////////////////////
// 🔰 DATABASE FUNCTIONS
///////////////////////////////////////////////////////////////////////////////////////////////////

func InitMongo() {
	// 🟢 [GENERAL] CREATE CONTEXT WITH TIMEOUT FOR CONNECTION
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 🟢 [GENERAL] SET CLIENT OPTIONS AND APPLY CONNECTION URI
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	
	// 🟢 [GENERAL] CONNECT TO MONGO DATABASE
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// 🔴 [ERROR HANDLING] PANIC IF CONNECTION FAILS
		panic(err)
	}

	Client = client
}
