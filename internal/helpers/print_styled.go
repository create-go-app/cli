package helpers

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// PrintStyled prints styled output for the given string by a lipgloss.Style template.
func PrintStyled(s, state, style string) {
	// Set lipgloss colors.
	successColor := lipgloss.AdaptiveColor{Light: "#16a34a", Dark: "#4ade80"}
	warningColor := lipgloss.AdaptiveColor{Light: "#ca8a04", Dark: "#facc15"}
	errorColor := lipgloss.AdaptiveColor{Light: "#dc2626", Dark: "#f87171"}
	infoColor := lipgloss.AdaptiveColor{Light: "#4b5563", Dark: "#9ca3af"}

	// Create a new lipgloss style instance.
	lg := lipgloss.NewStyle()

	// Switch between states.
	switch state {
	case "info":
		state = lg.Foreground(infoColor).SetString("– ").String()
	case "success":
		state = lg.Foreground(successColor).SetString("✓ ").String()
	case "error":
		state = lg.Foreground(errorColor).SetString("✕ ").String()
	case "warning":
		state = lg.Foreground(warningColor).SetString("‼ ").String()
	}

	// Concat state with the given string.
	concatStrings := strings.Join([]string{state, s}, "")

	// Switch between styles.
	switch style {
	case "margin-top-bottom":
		s = renderStyled(concatStrings, lg.MarginTop(1).MarginBottom(1))
	case "margin-top":
		s = renderStyled(concatStrings, lg.MarginTop(1))
	case "margin-bottom":
		s = renderStyled(concatStrings, lg.MarginBottom(1))
	case "margin-left":
		s = renderStyled(concatStrings, lg.MarginLeft(1))
	case "margin-left-2":
		s = renderStyled(concatStrings, lg.MarginLeft(2))
	default:
		s = concatStrings
	}

	// Print styled output.
	fmt.Println(s)
}

// RenderStyled render a styled string with a given lipgloss.Style template
// using "charmbracelet/lipgloss" package.
func renderStyled(str string, template lipgloss.Style) string {
	return template.Render(str)
}
