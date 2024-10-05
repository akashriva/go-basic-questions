package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // Corrected field tag
}

var (
	users []User
	mu    sync.Mutex
)

type CreateUserResponse struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

// Define the homeHandler function
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	var response CreateUserResponse
	// if user exits do not created
	for _, value := range users {
		if value.Email == user.Email {
			response = CreateUserResponse{
				Status: "User Already Exits",
				User:   User{},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
			}
			return
		}
	}
	users = append(users, user)

	response = CreateUserResponse{
		Status: "User created",
		User:   user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid userId", http.StatusBadRequest)
		return
	}
	fmt.Println("UserID ------------->", userId)
	mu.Lock()
	defer mu.Unlock()

	var response CreateUserResponse
	var user User
	for _, user = range users {
		if user.UserId == userId {
			response = CreateUserResponse{
				Status: "User Found",
				User:   user,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
			}
			return
		} else {
			response = CreateUserResponse{
				Status: "User not found",
				User:   user,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound) // Use 404 Not Found for user not found
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
			}
		}
	}

	response = CreateUserResponse{
		Status: "User not found",
		User:   user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound) // Use 404 Not Found for user not found
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// login middleware
// login middleware
func loginUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("email")
		password := r.Header.Get("password")

		// Log the headers for debugging
		fmt.Printf("Received Email: %v, Password: %v\n", email, password)

		// Check if headers exist
		if email == "" || password == "" {
			http.Error(w, "Missing email or password headers", http.StatusBadRequest)
			return
		}

		mu.Lock()
		defer mu.Unlock()

		var authentication bool
		for _, user := range users {
			if user.Email == email && user.Password == password {
				fmt.Printf("Authenticated user: %v\n", user.Email)
				authentication = true
				break
			}
		}

		if !authentication {
			fmt.Println("Authentication failed")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authentication successful, proceeding to next handler")
		next.ServeHTTP(w, r)
	})
}

func upDateUser(w http.ResponseWriter, r *http.Request) {
	// Extract userId from query parameter
	userIdStr := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid userId", http.StatusBadRequest)
		return
	}

	// Decode request body to get updateUserData
	var updateUserData User
	if err := json.NewDecoder(r.Body).Decode(&updateUserData); err != nil {
		http.Error(w, "Error decoding body", http.StatusBadRequest)
		return
	}

	// Lock the mutex before modifying the users slice
	mu.Lock()
	defer mu.Unlock()

	var response CreateUserResponse
	userUpdated := false

	// Iterate through users to find the user by ID
	for i, user := range users {
		if user.UserId == userId {
			// Update the user details
			users[i].Name = updateUserData.Name
			users[i].Email = updateUserData.Email
			users[i].Password = updateUserData.Password // Use updateUserData.Password

			// Prepare successful update response
			response = CreateUserResponse{
				Status: "Update successful",
				User:   users[i],
			}
			userUpdated = true
			break
		}
	}

	// Handle case where user is not found
	if !userUpdated {
		response = CreateUserResponse{
			Status: "User not found",
			User:   User{},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
		return
	}

	// Respond with the updated user details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func main() {
	// Register handlers with routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/create-user", createUser)
	http.HandleFunc("/users", getAllUser)
	http.HandleFunc("/get-user-by-id", getUserById)
	//middleware Router
	http.Handle("/update-user", loginUser(http.HandlerFunc(upDateUser)))
	// Start the server
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
