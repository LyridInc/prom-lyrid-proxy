package main

import (
	entry "PerseusNocApp.ExporterModuleNocPromGo/ExportersFunction"
	"github.com/joho/godotenv"
)

// this function is not a part of the serverless and will not be processed by Lyrid,
// this is for user to be able to locally test and build their functions
func main() {

	// Loads env variables
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//	return
	//}

	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
	godotenv.Load()

	router := entry.Initialize()
	router.Run(":3000")
}

/* Old handler is no longer valid
func handler(w http.ResponseWriter, r *http.Request) {

	fn.PreRun()

	decoder := json.NewDecoder(r.Body)
	var input fn.LyFnInputParams
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	fnOutput := fn.Run(input)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fnOutput)

	fn.PostRun()
}

*/
