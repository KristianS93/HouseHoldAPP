package assistants

import "net/http"

//EnableCors take the pointer to the ResponseWriter and sets the CORS
// Input has to be an ADDRESS to a responsewriter.
func EnableCors(w *http.ResponseWriter) {
	//Setting the cors for the pointer to the responsewriter.
	(*w).Header().Set("Access-Control-Allow-Origin", ConstAllowedCORS)
}

//SetHeader sets the content type to json, to a pointer to the responsewriter.
func SetHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}
