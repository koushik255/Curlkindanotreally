package funcs 

import (
	"fmt"
	"os"
	"net/http"
	"io"
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

	tea.Print_Nice("json")

	
	fmt.Println(string(body))
}



func main() {
	fmt.Println("HEllo dude ")


}