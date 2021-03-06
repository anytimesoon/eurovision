package routes

import (
	"eurovision/pkg/routes/comments"
	"eurovision/pkg/routes/countries"
	"eurovision/pkg/routes/users"
	"eurovision/pkg/routes/votes"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Start() {
	mime.AddExtensionType(".js", "application/javascript")

	router := mux.NewRouter()
	router.Use(addHeaders, logging)

	homeRouter := router.PathPrefix("/").Subrouter()
	homeRouter.HandleFunc("/", countries.All).Methods(http.MethodGet)

	// Country
	countryRouter := router.PathPrefix("/country").Subrouter()
	countryRouter.HandleFunc("/", countries.Update).Methods(http.MethodPut)
	countryRouter.HandleFunc("/{slug}", countries.FindOne).Methods(http.MethodGet)

	// User
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", users.All).Methods(http.MethodGet)
	userRouter.HandleFunc("/", users.Update).Methods(http.MethodPut)
	userRouter.HandleFunc("/new", users.Create).Methods((http.MethodPost))
	userRouter.HandleFunc("/{slug}", users.FindOne).Methods(http.MethodGet)
	userRouter.HandleFunc("/{slug}/rem", users.RemoveUser).Methods(http.MethodDelete)

	// Comment
	commentRouter := router.PathPrefix("/comment").Subrouter()
	commentRouter.HandleFunc("/", comments.All).Methods(http.MethodGet)
	commentRouter.HandleFunc("/new", comments.Create).Methods((http.MethodPost))
	// commentRouter.HandleFunc("/{uuid}/rem", comments.RemoveComment).Methods(http.MethodDelete)

	// Vote
	voteRouter := router.PathPrefix("/vote").Subrouter()
	voteRouter.HandleFunc("/new", votes.Create).Methods(http.MethodPost)
	voteRouter.HandleFunc("/", votes.Update).Methods(http.MethodPut)

	headersOk := handlers.AllowedHeaders([]string{"Content-type", "Authorization", "Origin", "Access-Control-Allow-Origin", "Accept", "Options", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router))) //keeps the server alive on port 8080
}
