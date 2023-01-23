package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://random-data-api.com/api/v2/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	//converted to json
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
	reqData:=map[string]string{}
	reqDataInt := map[string]int
	json.Unmarshal([]byte(responseData), &reqData, )
	json.Unmarshal([]byte(responseData), &reqDataInt)

	fmt.Println("firstname : " + reqData["first_name"])
	fmt.Println("Lastname : " + reqData["last_name"])
	fmt.Println("Username : " + reqData["Username"])
	fmt.Println("id : " + reqDataInt["id"])
}
