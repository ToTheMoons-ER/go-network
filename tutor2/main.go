package main

import (
	"fmt"
	"net/http"
	"strconv" //แปลง String เป็น Int

	"github.com/gorilla/mux"
)

func main() {
	// จัดการรีเควส
	r := mux.NewRouter()

	// เส้นทางสำหรับการบวกเลข
	r.HandleFunc("/cal/{Number1}/plus/{Number2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		Number1Str := vars["Number1"]
		Number2Str := vars["Number2"]

		// แปลงพารามิเตอร์ที่เป็นสตริงห้เป็นจำนวนเต็ม
		Number1, err1 := strconv.Atoi(Number1Str)
		Number2, err2 := strconv.Atoi(Number2Str)

		// ตรวจสอบข้อผิดพลาดในการแปลง
		if err1 != nil || err2 != nil {
			http.Error(w, "ตัวเลขไม่ถูกต้องใน URL", http.StatusBadRequest)
			return
		}

		// ทำการบวกเลข
		total := Number1 + Number2

		// ตอบด้วยผลลัพธ์ที่ตกแต่งด้วย HTML
		fmt.Fprintf(w, "<h1>Calculate: %s plus %s </h1>", Number1Str, Number2Str)

		// ตอบด้วยผลลัพธ์
		fmt.Fprintf(w, "Total = %d\n", total)
	})

	// Handler function for the root ("/") URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// Handler function for the "/about" URL
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My name is Kanya")
	})

	// Serve static files from the "static/" directory
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", r)
}