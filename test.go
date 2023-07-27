// ? Go 파일이 패키지 main에 속해 있다는 것을 나타냅니다. 이는 실행 가능한 프로그램을 작성하는 데 필요한 패키지입니다.
package main

//? 필요한 패키지를 가져옵니다. fmt은 포맷된 출력을 제공하고, net/http는 HTTP 프로토콜을 사용하는 웹 서버를 구축하기 위한 패키지
import (
	"fmt"
	"net/http"
)

// ? 프로그램의 진입점인 main 함수입니다. 프로그램 실행 시 처음으로 실행되는 함수
func test_main() {
	//* '/' 경로로 들어오는 HTTP 요청을 처리하기 위한 루트 핸들러를 정의
	//* w http.ResponseWriter는 클라이언트에게 응답을 작성하는 데 사용되는 인터페이스
	//* r *http.Request는 클라이언트의 요청을 나타내는 객체
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "안녕하세요, Go 앱 서버!")
	})

	//?  서버 시작 중에 발생하는 에러를 처리하기 위해 ListenAndServe 함수 호출의 결과를 err 변수에 할당
	err := http.ListenAndServe(":2080", nil)
	if err != nil {
		fmt.Println("서버 시작 실패:", err)
	} else {
		fmt.Println("서버 돌아감")
	}
}
