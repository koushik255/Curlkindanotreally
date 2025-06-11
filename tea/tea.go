package tea

import (
	"fmt"

	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/spf13/cobra"
	lip "github.com/charmbracelet/lipgloss"

)

func Print_Nice(x string,y string) {
	var style = lip.NewStyle().
    Bold(true).
    Underline(true).
    Foreground(lip.Color("#FAFAFA")).
    Background(lip.Color("#7D56F4")).
    PaddingTop(0).
    PaddingLeft(2).
    Width(2)

	fmt.Println(style.Render("","",x,y))
}
func Print_idk(x string, y string){
	var style = lip.NewStyle().
  	  BorderStyle(lip.NormalBorder()).
    	BorderForeground(lip.Color("63"))

    fmt.Println(style.Render("","",x,y))

// // Set a rounded, yellow-on-purple border to the top and left
// 	var anotherStyle = lipgloss.NewStyle().
//    		BorderStyle(lipgloss.RoundedBorder()).
//     	BorderForeground(lipgloss.Color("228")).
//     	BorderBackground(lipgloss.Color("63")).
//     	BorderTop(true).
//     	BorderLeft(true)

}




func main() {
	fmt.Println("sup")
}