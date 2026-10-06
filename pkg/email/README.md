# 📦 Package `email`

**Source Path:** `./pkg/email`

## 🚀 Functions

### `GenerateEmailHTML`

GenerateEmailHTML takes a map of key-value pairs and generates an HTML
email string based on the values provided. The key-value pairs should
be in the following format:

  - Title: The title of the email
  - Subject: The subject of the email
  - Message: The body message of the email
  - Button: The text for the call-to-action button
  - Link: The URL for the call-to-action button

The generated HTML will be responsive and have a basic CSS style.

```go
func GenerateEmailHTML(data map[string]string) string
```

