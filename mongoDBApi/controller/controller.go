package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/aashutoshchaudhary/mongapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://AashuChaudhary:AshuDips26@cluster0.os8fhfm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// Most imortant as a connection
var collection *mongo.Collection

// Connect with mongodb

func init() { //specialised method in gollang which used to initialised

	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Db connection Sucessfully established")

	collection = client.Database(dbName).Collection(colName)

	// collection Instance /colection reference is ready
	fmt.Println("Collection instance is Ready")

	//Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

	// func Background() Context
	// Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

	// func TODO() Context
	// TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available

}

// MONGODB HELPERS - FILE

// INSERT 1 RECORD :takes data and add to mongodb

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie with id: ", inserted.InsertedID)
}

// UPDATE DATABASE 1 RECORD

func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}} // we can use both bson.M or bson.D

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count : ", result.ModifiedCount)

}

// Delete record

func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	delete, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete count with delete count : ", delete.DeletedCount)
}

// delete all records

func deleteAllMovies() int64 {
	// bson.D{} is used to delete all documents in the collection
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All movies got deleted count with delete count : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// want to find all the records from mongodb:
// create read operation using filter function

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M //we are going to create one movie ,this movies is whole list of movies .

		err := cursor.Decode(&movie) //cursor will try to decode the value.
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// working out and bring controllerswhich are going to take help of mongodb helperand trow some data  , and remember all the controllers and db helpers created all first letter is lowercase ,becoz  we dont want to send or export them .

// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	// at any router this method is being call  and i am going to give you all the movies
	w.Header().Set("Content-Type", "application/x-www-form-urlencode") // we can use application/json isntead of :application/x-www-form-urlencode.

	//setting more header and values  and this all done through htttp method so we dont need to use mux or anything.

	// GRAB ALLL MOVIES : here all the helper method  comes in
	allMovies := getAllMovies()

	// send as a response
	json.NewEncoder(w).Encode(allMovies)

}

// CREATING A MOVIE IN DATABASE,THE GOAL IS THAT WE ARE CREATING A VARIABLE  OF TYPE MODEL.NETFLIX AND A STRUCTURE I ALREADY HAS BEEN DEFINED A JSON WILL COME UP FROM FROENTED AND WE CAN DECODE OUR JSON ACCORDING TO THEIR STRUCTURE  OR THE DATATYPE WE HAVE DEFINED AND TAKE HELP FROM THE MONGO DB HELPER  FILE AND PROVIDE A DATA TO IT AND IT DOES JOB TO IT.

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") //one MORE INTERESTING THINGALTHOUGH WE ARE SETTING OUR HEADER HERE ,THESE ROUTER THESE PACKAGES  (GORILAMUX) ALSO ALLOWS US TO SET THE HEADER  WHILE WE ARE HANDLING THE ROUTER

	var movie model.Netflix                    //structure is ready
	_ = json.NewDecoder(r.Body).Decode(&movie) // calling json to decode the structure which require body and the decode it for that it pass reference of strucute
	insertOneMovie(movie)
	// craft response
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r) // we are using mux package here to get movie id from URL
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
