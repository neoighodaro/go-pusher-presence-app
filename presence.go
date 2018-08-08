package main

// Here, we import the required packages (including Pusher)
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pusher "github.com/pusher/pusher-http-go"
)

// Here, we register the Pusher client
var client = pusher.Client{
    AppId:   "PUSHER_APP_ID",
    Key:     "PUSHER_APP_KEY",
    Secret:  "PUSHER_APP_SECRET",
    Cluster: "PUSHER_APP_CLUSTER",
    Secure:  true,
}

// Here, we define a user struct
type user struct {
	Username  string `json:"username" xml:"username" form:"username" query:"username"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// Here, we create a global user variable to hold user details for a session
var loggedInUser user

// Here, we check if a user is logged in
func isUserLoggedIn(rw http.ResponseWriter, req *http.Request){
	if loggedInUser.Username != "" && loggedInUser.Email != "" {
		json.NewEncoder(rw).Encode(loggedInUser)
	} else {
		json.NewEncoder(rw).Encode("false")
	}
}

// -------------------------------------------------------
// Here, we receive a new user's details in a POST request and
// bind it to an instance of the user struct, we further use this
// user instance to check if a user is logged in or not
// -------------------------------------------------------
func NewUser(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &loggedInUser)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(rw).Encode(loggedInUser)
}


// -------------------------------------------------------
// Here, we authorize users so that they can subscribe to the presence channel
// -------------------------------------------------------
func pusherAuth(res http.ResponseWriter, req *http.Request) {

	params, _ := ioutil.ReadAll(req.Body)

	presenceData := pusher.MemberData{
		UserId: loggedInUser.Username,
		UserInfo: map[string]string{
			"email": loggedInUser.Email,
		},
	}
	
	response, err := client.AuthenticatePresenceChannel(params, presenceData)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Fprintf(res, string(response))
}

func main() {

	// Serve the static files and templates from the static directory
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Here, we determine if a user is logged in or not.
	http.HandleFunc("/isLoggedIn", isUserLoggedIn)

	// -------------------------------------------------------
	// Listen on these routes for new user registration and user authorization,
	// thereafter, handle each request using the matching handler function.
	// -------------------------------------------------------
	http.HandleFunc("/new/user", NewUser)
	http.HandleFunc("/pusher/auth", pusherAuth)

	// Start executing the application on port 8090
	log.Fatal(http.ListenAndServe(":8090", nil))
}