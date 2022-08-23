package meal

import (
	"fmt"
	"net/http"
)

func DeleteMeal(w *http.ResponseWriter, r *http.Request) {
	fmt.Println("Du har tilføjet et meal.")
}
