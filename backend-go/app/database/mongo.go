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

package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var BeerCollection *mongo.Collection

///////////////////////////////////////////////////////////////////////////////////////////////////
// 🔰 DATABASE FUNCTIONS
///////////////////////////////////////////////////////////////////////////////////////////////////

func ConnectDB(mongoURI string) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // 🟢 [GENERAL] Client Creation and Connection
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("[Mongo] Erro ao conectar:", err)
    }

    // 🟢 [GENERAL] Connection Verifications
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("[Mongo] Erro ao verificar conexão:", err)
    }
    db := client.Database("beerdb")
    collections, _ := db.ListCollectionNames(ctx, bson.M{"name": "beer_recognition"})
    if len(collections) == 0 {
        err := db.CreateCollection(ctx, "beer_recognition")
        if err != nil {
            log.Fatal("[Mongo] Erro ao criar a coleção:", err)
        }
        log.Println("[Mongo] Coleção 'beer_recognition' criada com sucesso.")
    } else {
        log.Println("[Mongo] Coleção 'beer_recognition' já existe.")
    }
    BeerCollection = db.Collection("beer_recognition")
    log.Println("[Mongo] Conectado em:", mongoURI)
}

// 🟢 [GENERAL] BeerRecognitionResult IN collection
type BeerRecognitionResult struct {
    BrandName string    `bson:"brand_name"`
    Image     []byte    `bson:"image"`
    CreatedAt time.Time `bson:"created_at"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// 🔰 MONGO CRUD FUNCTIONS
///////////////////////////////////////////////////////////////////////////////////////////////////

func SaveToDatabase(ctx context.Context, image []byte, brand string) error {
    doc := BeerRecognitionResult{
        BrandName: brand,
        Image:     image,
        CreatedAt: time.Now(),
    }

    _, err := BeerCollection.InsertOne(ctx, doc)
    if err != nil {
        log.Println("[Mongo] Erro ao salvar:", err)
        return err
    }

    log.Println("[Mongo] Documento salvo com sucesso:", brand)
    return nil
}
