package plan

import (
	"fmt"
	"net/http"
)

func GetPlan(w *http.ResponseWriter, r *http.Request) {
	fmt.Println("Noget med en plan")
}
