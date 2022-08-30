package plan

import (
	"fmt"
	"net/http"
)

func DeletePlan(w http.ResponseWriter, r *http.Request) {
	// Delete plan skal slette planen fra plan db og fra household, det er ikke nødvendigt at slette meals fra planen, da meals kan stadig være relevante.
	fmt.Println("Noget med en plan")
}
