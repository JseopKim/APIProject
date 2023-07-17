package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 루트 핸들러 정의
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "안녕하세요, Go 앱 서버!")
	})

	// 서버 시작
	err := http.ListenAndServe(":2080", nil)
	if err != nil {
		fmt.Println("서버 시작 실패:", err)
	} else {
		fmt.Println("서버 돌아감")
	}
}
