package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1)
		// m.Unlock()

		time.Sleep(time.Millisecond * 300)

		w.Write([]byte(fmt.Sprintf("Number: %d", number)))
	})

	http.ListenAndServe(":8080", nil)
}
