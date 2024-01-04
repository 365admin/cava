package cmd

import (
	"encoding/json"
	"fmt"
	"log"
)

func printData(data any, err error) {
	if err != nil {
		log.Fatalln(err)
		return
	}
	printJSON(data)
}
func printJSON(v any) {
	j, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	s := string(j)
	fmt.Print(s)

}
