db = db.getSiblingDB('beerdb');

db.createCollection('beer_recognition');

db.beer_recognition.insertOne({
  brand_name: "Example Beer",
  image: null,
  created_at: new Date()
});
