package tea

import (
	"fmt"

	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/spf13/cobra"
	lip "github.com/charmbracelet/lipgloss"
)

func Print_Nice(x string) {
	var style = lip.NewStyle().
    Bold(true).
    Foreground(lip.Color("#FAFAFA")).
    Background(lip.Color("#7D56F4")).
    PaddingTop(2).
    PaddingLeft(2).
    Width(2)

	fmt.Println(style.Render("",x))
}




func main() {
	fmt.Println("sup")
}