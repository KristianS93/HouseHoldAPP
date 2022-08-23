package meal

import (
	"fmt"
	"net/http"
)

func GetMeal(w *http.ResponseWriter, r *http.Request) {
	fmt.Println("Du har tilføjet et meal.")
}
