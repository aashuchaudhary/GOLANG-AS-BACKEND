// ACTION PLAN WE ARE CREATING A LEARNCODE ONLINE SIMILAR KIND OF LOOKING BACKEND API WILL   WE WORKING FOR COURCESAND WANT TOGIVE A FEATURE THAT A USER CAN ACTULAAY GET ALL THE COURSES THEY ARE ABLE TO CREATE NEW COURCES ,DELETE COURSES AND UODATE THE EXISTING NEW COURCES AND ALSO WE WANT TO SEE THAT IF THERE IS NO UNIQUE ID OR NO TITLE NAME BEING GIVEN TO A COURCE AND THE COURSE WILL NOT ADDED,SO WE WANT TO CREATE A HELPER METHOD, FOR HAVING ALL THE DATA WE ARE USING SLICE AS OUR FAKE DATABASE SINCEE WE DONT HAVE DTABASE .

// AND FURTHER MORE WE ARE GOING TO INJECT GORILLA MUX,SO THAT EVERY SINGLE ROUTE WILL BE HANDEL IN SEPERATE METGOD FOR ALLCRUD OPERATION, AS CREATE ,UPDATE AND DELETE OPERATION

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Creating a database that how our course is going to look-like: that'swhy we are creating a model.

// model for cource-file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// course having Author

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB - slice of type course
var courses []Course

// middleware/ helper - file  :its help us to perform some task as encypy the password,not allow user to pull in database without a unique course id or course name.

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.courseName == ""
	return c.CourseName == "" // this is why we dont want to check for course id ,we want user to move further if course id is not empty,why we do so becoz we want to crease course id manually as i dont want to relly on user you go an dpass me a unique id unlikely the case the pass me that
}

func main() {
	fmt.Println("All about api creating and operations")
	r := mux.NewRouter()

	// Seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 299, Author: &Author{Fullname: "Ashutosh", Website: "Lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Nextjs", CoursePrice: 999, Author: &Author{Fullname: "Ashu", Website: "nextjs.dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Django", CoursePrice: 1099, Author: &Author{Fullname: "Dipu", Website: "py.dev"}})
	courses = append(courses, Course{CourseId: "6", CourseName: "Golang", CoursePrice: 1999, Author: &Author{Fullname: "Appu", Website: "go.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("Get") //take the router and call the  routing function and define a route by salsh
	r.HandleFunc("/courses", getAllCourses).Methods("Get")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("Get")
	r.HandleFunc("/course", createOneCourse).Methods("POST") //POST AND PUT USED TO GET THE DATA
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// defining how to handel the situation is call as controler that is a brain of a operation

// COntrollers-files
// serve home route : the reason we are cresting the serve home route as we open our Application  i dont want a blank and empty page  or error in a case ,if anyone visit local host/ our application  an dwant to send some message  for this and we want to create simple serve home

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

// creating another route : the goal of this route is simple  that we want to throw all the  values from database

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")

	// throwup all the things we are in the fake db as a json
	json.NewEncoder(w).Encode(courses)

	// SEEDING: WHEN WE ARE TESTING OR DOING SUCH KIND OF OPERATIONS WE SEED OUR DATABASE TO FAKE VALUES.

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// here we take help of the request  whenever somebody is to require just one course he is going to provide uniqueid ,if someone is passing the uniqueid then the first thing we do to grab our id is task no 1 and  then since i have all courses as an array and just i can go and look into array and compare this id to somebody that the id is available or not ,jsut i have to run through a loop ,compare the value and just return that

	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	// params is a holder of key value pair
	params := mux.Vars(r)
	fmt.Println(params)

	// loop through courses,find matching id and return the response
	for _, course := range courses {
		// for find match id
		if course.CourseId == params["id"] {
			// return response
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")

}

// Injecting Something tO Database

// There are lots of precaution that you can create alots of checker method  to make scure that user is not sending malliscious data

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// what if miode: body is empty
	if r.Body == nil {
		// crafting json response
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about whwn data is sen like-->{}

	var course Course
	// destructing  json
	_ = json.NewDecoder(r.Body).Decode(&course)
	// checking wheter its empty or not givin boolean value
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json you are sending")
		return
	}
	
	// Todo check if title is duplicate
	// loop, title,matches with course.coursename,JSON
	for _, c := range courses {
        if c.CourseName == course.CourseName {
            json.NewEncoder(w).Encode("Course with same name already exists")
            return
        }
    }

	// generate unique id,and convert the id into string , and then after we append course into courses

	// for generating unique id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// update Operation

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	// in case of update having a couples of question that what is the one course should i update so that can be easily identified by a unique course id, the second thing is i need all those value that should be updated  during the update operation what happens resend all the values regardless what the value is filledin or not this is a regular operation ....
	// now since we are not in the database we dont have libberty of going on and saying find by id and update smilar operation in mysql we need to go through the way that how we are going to acomplish first.we are going to look all the values an donce we hit the value as an id  the we are going to take the values  slice of that becoz remeber this is an array or slice of having differnet values we are just going to remove that values an dsince we are having alreadt that id in params so we are simply going to inject the values....

	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop through the values and once we get the  id  so once we hit the id we need to go and remove that value from index and the add the value againin the courses but this time the add is going to happen  add with my id : this is the exactly id we are getting in params

	for index, course := range courses {
		// get id
		if course.CourseId == params["id"] {
			// remove the value from index
			courses = append(courses[:index], courses[index+1:]...)
			// add with my id
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) //json data is coming to me and i want to grab it
			course.CourseId = params["id"]              // update operation so the id will be same
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is notfound
	json.NewEncoder(w).Encode("No course found with given ID")
}

// Delete Operation
// func  DeleteAllCourse -- assignment  todo by myself

// Delete Operation
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// having unique id
	params := mux.Vars(r)

	// loop through courses, find matching ID, and remove the course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("This course has been removed")
			return
		}
	}

	// send a response if the course is not found
	json.NewEncoder(w).Encode("No course found with the given ID")
}

// seeding the data as well as making some of the routesthat can use these controllers
