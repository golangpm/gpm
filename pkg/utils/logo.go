package utils

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

// Logotype output
func Logo() {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	// options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
	}
	renderStr, _ := ascii.RenderOpts("GPM", options)
	fmt.Print(renderStr)
}
