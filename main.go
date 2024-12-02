package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func setupClient(ctx context.Context) (*genai.Client, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	fmt.Sprint(err)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getGoogleShopUrl(query string) string {
	baseUrl := "https://www.google.com/search?tbm=shop&q=" + query
	return baseUrl
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the uploaded file from the form data
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Initialize Gemini API client
	client, err := setupClient(ctx)
	if err != nil {
		http.Error(w, "Failed to set up Gemini client", http.StatusInternalServerError)
		return
	}

	model := client.GenerativeModel("gemini-1.5-pro")

	imageBytes, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}

	// Define the prompt
	prompt := `

		**Task:** Analyze the given image and identify multiple objects. For each detected object:

		- Focus only on products; ignore people (e.g., human, man, woman, child).
		- Extract the product's name, company, and a brief description (maximum 50 words).
		- Identify two similar products based on available product details.

		**Output Format:**
		Return the data in a structured JSON format as shown below:

		[
		    {
		        "name": "Product Name",
		        "desc": "This product is used for ... and has use cases such as ...",
		        "similar_products": ["Similar Product 1", "Similar Product 2"]
		    },
		    {
		        "name": "Another Product Name",
		        "desc": "This product is used for ... and has use cases such as ...",
		        "similar_products": ["Similar Product A", "Similar Product B"]
		    }
		]


		**Important Notes:**
		1. Exclude "..." placeholders or incomplete entries in the similar_products field.
		2. Ensure descriptions are concise and relevant.
		`

	// Generate response from Gemini API
	genResp, err := model.GenerateContent(ctx, []genai.Part{
		genai.ImageData("jpeg", imageBytes),
		genai.Text(prompt),
	}...)
	if err != nil {
		http.Error(w, "Failed to generate content", http.StatusInternalServerError)
		return
	}

	// Respond with the result in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(genResp)
	w.Write(b)
}

func main() {
	http.HandleFunc("/analyze-image", imageHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
