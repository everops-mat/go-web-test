package handlers

import (
	"go-web-test/internal/logger"
	"go-web-test/internal/sayings"
	"html/template"
	"net/http"
)

// HTML template for rendering the excuse
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>EverOpin' Style Opportunity Generator</title>
    <style>
        body { font-family: Arial, sans-serif; text-align: center; background-color: #f4f4f4; }
        .container { max-width: 600px; margin: 50px auto; background: white; padding: 20px; border-radius: 10px; box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); }
        .excuse { font-size: 1.5em; font-weight: bold; color: #20B2AA; }
    </style>
</head>
<body>
    <div class="container">
        <h1>EverOpin' Style Opportunity Generator</h1>
        <p>The cause of the problem is:</p>
        <p class="excuse">{{.}}</p>
        <p><small>Reload for a new excuse.</small></p>
    </div>
</body>
</html>
`

// RandomSayingHandler returns a random saying as an HTML response
func RandomSayingHandler(w http.ResponseWriter, r *http.Request) {
	// Get a random saying from internal/sayings package
	saying, err := sayings.GetRandomSaying()
	if err != nil {
		logger.JSONLogger("error", "No sayings available")
		http.Error(w, "No sayings available", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template
	tmpl, err := template.New("excuse").Parse(htmlTemplate)
	if err != nil {
		logger.JSONLogger("error", "Failed to parse HTML template: "+err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response header
	w.Header().Set("Content-Type", "text/html")

	// Execute the template and send response
	if err := tmpl.Execute(w, saying); err != nil {
		logger.JSONLogger("error", "Failed to render HTML: "+err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// Log the served saying
	logger.JSONLogger("info", "Served saying: "+saying)
}
