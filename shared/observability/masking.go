package observability

// MaskIdentifier keeps only the last four characters of a sensitive identifier.
// It is intended for logs and traces, never for authentication or persistence.
func MaskIdentifier(value string) string {
	if value == "" {
		return ""
	}
	if len(value) <= 4 {
		return "****"
	}
	masked := make([]byte, len(value)-4)
	for i := range masked {
		masked[i] = '*'
	}
	return string(masked) + value[len(value)-4:]
}
