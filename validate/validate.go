package validate

import (
	"regexp"
	"strings"
)

var (
	fenRe   = regexp.MustCompile(`^([rnbqkbnrRNBQKBNR1-8]{1,8}/){7}[rnbqkbnrRNBQKBNR1-8]{1,8}\s+[wb]\s+([KQkqA-Ha-h]{1,4}|-)\s+([a-h][36]|-)\s+\d+\s+\d+$`)
	moveRe  = regexp.MustCompile(`^[a-h][1-8][a-h][1-8][qrbn]?$`)
	alphaRe = regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)
)

func IsValidFEN(fen string) bool {
	return fenRe.MatchString(strings.TrimSpace(fen))
}

func IsValidMove(move string) bool {
	return moveRe.MatchString(strings.TrimSpace(move))
}

func IsAlphaNumeric(s string) bool {
	return alphaRe.MatchString(s)
}

func SanitizeInput(s string, maxLen int) string {
	t := strings.TrimSpace(s)
	t = strings.ReplaceAll(t, "\x00", "")
	t = strings.ReplaceAll(t, "\r", "")
	if len(t) > maxLen {
		return t[:maxLen]
	}
	return t
}

func EscapeTemplateString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "`", "\\`")
	s = strings.ReplaceAll(s, "${", "\\${")
	return s
}
