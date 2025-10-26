package formatter

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/ty-cs/udict/internal/api"
)

const (
	// ANSI colors 0-8
	colorBlack   = "0"
	colorRed     = "1"
	colorGreen   = "2"
	colorYellow  = "3"
	colorBlue    = "4"
	colorMagenta = "5"
	colorCyan    = "6"
	colorWhite   = "15"
)

func StyleDefinition(response api.Response, limit int) string {
	var s strings.Builder
	re := regexp.MustCompile(`[\r\n]+`)

	// Base style for all indented sections
	contentSharedStyle := lipgloss.NewStyle().PaddingLeft(4)
	titleStyle := lipgloss.NewStyle().PaddingLeft(2)

	// Styles using ANSI 1-16 colors, inheriting from the base style
	wordStyle := lipgloss.NewStyle().
		Bold(true).
		PaddingLeft(1).
		PaddingRight(1).
		Foreground(lipgloss.Color(colorWhite)).
		Background(lipgloss.Color(colorYellow))
	sectionTitleStyle := titleStyle.Foreground(lipgloss.Color(colorYellow))
	contentStyle := contentSharedStyle.Foreground(lipgloss.Color(colorWhite))
	exampleStyle := contentSharedStyle.Foreground(lipgloss.Color(colorBlue))
	authorStyle := contentSharedStyle.Foreground(lipgloss.Color(colorWhite))
	upvoteStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colorGreen))
	downvoteStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colorRed))

	for i, d := range response.List {
		if i >= limit {
			break
		}

		s.WriteString(wordStyle.Render(d.Word))
		s.WriteString("\n")

		// Definition
		cleanDefinition := strings.TrimSpace(re.ReplaceAllString(d.Definition, "\n"))
		cleanDefinition = strings.ReplaceAll(cleanDefinition, "[", "")
		cleanDefinition = strings.ReplaceAll(cleanDefinition, "]", "")
		s.WriteString(sectionTitleStyle.Render("Definition:"))
		s.WriteString("\n")
		s.WriteString(contentStyle.Render(cleanDefinition))
		s.WriteString("\n\n")

		// Example
		cleanExample := strings.TrimSpace(re.ReplaceAllString(d.Example, "\n"))
		cleanExample = strings.ReplaceAll(cleanExample, "[", "")
		cleanExample = strings.ReplaceAll(cleanExample, "]", "")
		s.WriteString(sectionTitleStyle.Render("Example:"))
		s.WriteString("\n")
		s.WriteString(exampleStyle.Render(cleanExample))
		s.WriteString("\n\n")

		// Votes
		s.WriteString(sectionTitleStyle.Render("Votes:"))
		s.WriteString("\n")
		votes := fmt.Sprintf("▲ %s | ▼ %s",
			upvoteStyle.Render(fmt.Sprintf("%d", d.ThumbsUp)),
			downvoteStyle.Render(fmt.Sprintf("%d", d.ThumbsDown)),
		)
		s.WriteString(contentSharedStyle.Render(votes))
		s.WriteString("\n\n")

		// Author and Date
		date, err := time.Parse(time.RFC3339, d.WrittenOn)
		formattedDate := ""
		if err == nil {
			formattedDate = date.Format("Jan 2, 2006")
		}
		s.WriteString(sectionTitleStyle.Render("Author:"))
		s.WriteString("\n")
		authorInfo := fmt.Sprintf("%s (%s)", d.Author, formattedDate)
		s.WriteString(authorStyle.Render(authorInfo))
		s.WriteString("\n")

		if i < len(response.List)-1 && i < limit-1 {
			s.WriteString("\n")
		}
	}

	return s.String()
}
