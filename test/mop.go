package test
import (
	"github.com/crmathieu/mop"
//	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func main() {

	mx := http.NewServeMux()
    mx.HandleFunc("/", foo)

    handler := cors.New(cors.Options{
        AllowedHeaders: []string{
            "Authorization",
            "Origin",
            "X-Requested-With",
            "Accept",
            "X-CSRF-Token",
            "Content-Type"},
        AllowCredentials: true,
    }).Handler(mx)
	mop.SetServer(&http.Server{Addr: ":8022", Handler: handler})
}
