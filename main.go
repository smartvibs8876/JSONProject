package main

import "JSONProject.com/controllers"

func main() {
	// router := mux.NewRouter()
	// initaliseHandlers(router)
	// log.Fatal(http.ListenAndServe(":8090", setHeaders(router)))
	controllers.GetNewJSON()
}

// func initaliseHandlers(router *mux.Router) {
// 	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintln(res, "Welcome to the project")
// 	}).Methods("GET")
// 	router.HandleFunc("/oldJSON", controllers.GetOldJSON).Methods("GET")
// 	router.HandleFunc("/newJSON", func(res http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintln(res, "Welcome to the project")
// 	}).Methods("GET")
// }

// func setHeaders(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
// 		if r.Method == "OPTIONS" {
// 			return
// 		}
// 		h.ServeHTTP(w, r)
// 	})
// }
