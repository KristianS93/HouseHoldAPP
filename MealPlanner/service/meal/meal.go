package meal

import (
	"fmt"
)

func CreateMeal(w *http.ResponseWriter, r *http.Request) {
	fmt.Println("Du har tilføjet et meal.")
}