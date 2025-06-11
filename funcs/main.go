package funcs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"bytes"
	"log"


	// "github.com/charmbracelet/lipgloss"
	tea "github.com/shikmade/my-cli-app/tea"
)







func SendGet(x string) {
	url := x
	//send
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close() 

	//check response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		os.Exit(1)
	}

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	szt := strconv.Itoa(resp.StatusCode)

	tea.Print_idk("RESPONSE CODE: ",szt)
	tea.Print_Nice("RESPONSE CODE:",szt)

	fmt.Println(string(body))
}



type Post_Request struct {
	Name string `json:"name"`
	Value string 	`json:"value"`
	More  string    `json:"more"`
}


func Get_data(x string,y string,z string)(Post_Request){
	post := Post_Request{
		Name : x,
		Value: y,
		More : z,
	}
	return post
}



func SendPost(x string, data Post_Request) {
	url := x 

	fmt.Println(data)
	
	jsonData,err := json.Marshal(data)
	if err != nil {
		log.Fatalf("error marshalling JSON: %s",err)
		os.Exit(1)
	}

	 body := bytes.NewBuffer(jsonData)


	fmt.Println()
	resp,err := http.Post(url, "application/json",body)
	if err != nil {
		fmt.Println("Errponmg senrdng requestst :",err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("REquest failed with sattus code:", resp.StatusCode)
		os.Exit(1)
	}

	respbody,err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:",err)
		os.Exit(1)
	}

	szt := strconv.Itoa(resp.StatusCode)

	tea.Print_Nice("RESPONSE CODE:",szt)
	fmt.Println(string(respbody))

}



func main() {
	fmt.Println("HEllo dude ")


}