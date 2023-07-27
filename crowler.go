package main

import (
	"fmt"
	// "go/types"
	"log"
	"net/http"

	// "reflect"
	"encoding/json"

	"golang.org/x/net/html"
)

type Product struct {
	Name string `json: "\name"`
}

func main() {
	url := "https://kream.co.kr/products/89369"

	// 웹페이지 다운로드
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to download the page: %s", err)
	}
	defer resp.Body.Close()

	// 응답이 정상적으로 도착했는지 확인
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to retrieve the page: %s", resp.Status)
	}

	// HTML 파싱
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %s", err)
	}

	// 웹페이지의 내용 추출
	content := extractContent(doc)

	// content 값을 JSON 형식으로 변환
	jsonData, err := json.Marshal(content)
	if err != nil {
		log.Fatalf("Failed to convert content to JSON: %s", err)
	}

	jsonString := string(jsonData)

	// JSON 문자열을 구조체로 디코딩
	var product Product
	err = json.Unmarshal([]byte(jsonString), &product)
	if err != nil {
		fmt.Println("JSON 디코딩 에러:", err)
		return
	}

	// 추출된 내용 출력
	fmt.Println(product.Name)

}

func extractContent(n *html.Node) string {
	var content string

	// 텍스트 노드를 만나면 내용을 추출하여 content에 추가
	if n.Type == html.TextNode {
		content += n.Data
	}

	// 자식 노드들을 재귀적으로 방문하며 내용 추출
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content += extractContent(c)
	}

	return content
}
