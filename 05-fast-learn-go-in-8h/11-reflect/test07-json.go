package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float64  `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "吴孟达"}}
	//结构体 转换为 json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Println(string(jsonStr))
	fmt.Printf("json str is : %s\n", jsonStr)

	// json str >>> struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Println(myMovie)
	fmt.Printf("myMovie is :%v\n", myMovie)
}
