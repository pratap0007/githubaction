package dashboard

import (
	"fmt"
	"net/http"
)

func Test() {
	fmt.Println("testing the code base")
}
func Details(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Here are the dashboard details")
}
