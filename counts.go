// Servidor web que nos permite contar las visitas a nuestra URL 
// Mostrar las veces que nuestra URL es visitada
package main
import (
        "fmt"
        "log"
        "net/http"
        "sync"
)
var     p sync.Mutex
var     count int

func counter(w http.ResponseWriter, r *http.Request){
        p.Lock()
        count++
        fmt.Fprintf(w, "Hello World! I have been seen %d times", count)
        p.Unlock()
}
func main() {
        http.HandleFunc("/", counter)                // We can change de URL
        log.Fatal(http.ListenAndServe(":8081", nil)) // We can change de port
}