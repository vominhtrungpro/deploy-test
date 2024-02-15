package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Tạo một struct để đại diện cho dữ liệu
type Data struct {
	Message string `json:"message"`
}

func main() {
	// Khai báo một route và handler cho API
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// Tạo một instance của struct Data
		data := Data{
			Message: "Hello, world!",
		}

		// Chuyển đổi dữ liệu sang định dạng JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Thiết lập Header cho phản hồi HTTP
		w.Header().Set("Content-Type", "application/json")

		// Gửi dữ liệu JSON dưới dạng phản hồi HTTP
		w.Write(jsonData)
	})

	// Khởi động máy chủ và lắng nghe cổng 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
