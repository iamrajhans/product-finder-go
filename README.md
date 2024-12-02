# Image Analysis API Backend

This is a Go-based backend service that utilizes the Google Gemini Vision API to analyze images. It identifies objects in an image and provides detailed information about each object, such as its name, description, and similar products.

---

## Features

- **Upload and Analyze Images**: Users can upload images via an API endpoint.
- **Object Recognition**: Detects objects in images and excludes humans.
- **Detailed Object Information**: Provides:
  - Product name
  - Product description (50 words max)
  - Up to 2 similar products
- **Structured JSON Response**: Returns data in a clean and structured format.

---

## Technologies Used

- **Programming Language**: Go
- **API**: [Google Gemini Vision API](https://ai.google.dev/gemini-api/docs/vision?lang=go)
- **Frameworks**: Standard Go HTTP library

---

## Prerequisites

1. **Go Installation**: Ensure you have Go installed. Follow the [official guide](https://golang.org/doc/install) for installation.
2. **Google Gemini API Key**:
   - Visit [Google AI Studio](https://aistudio.google.com/).
   - Create and secure your API key.

---

## Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/iamrajhans/product-finder-go.git
   cd product-finder-go
   ```

2. Initialize the Go module:
   ```bash
   go mod init product-finder-go
   ```

3. Install the Google Generative AI Go SDK:
   ```bash
   go get github.com/google/generative-ai-go
   ```

4. Set up your environment variable:
   ```bash
   export GEMINI_API_KEY=your_api_key_here
   ```

5. Build and run the application:
   ```bash
   go run main.go
   ```

---

## API Endpoints

### POST `/analyze-image`

**Description**: Accepts an image file, processes it using the Google Gemini Vision API, and returns object details.

**Request**:
- **Method**: `POST`
- **Content-Type**: `multipart/form-data`
- **Body**: Form data with an `image` field containing the image file.

**Response**:
- **Content-Type**: `application/json`
- **Body**:
  ```json
  [
      {
          "name": "Product Name",
          "description": "Product description in 50 words max.",
          "similar_products": ["Similar Product 1", "Similar Product 2"]
      },
      ...
  ]
  ```

**Example**:
```bash
curl -X POST -F "image=@product_image.jpg" http://localhost:8080/analyze-image
```

---

## Example Response

```json
[
    {
        "name": "Smartphone XYZ",
        "description": "A high-performance smartphone designed for seamless multitasking and superior camera quality.",
        "similar_products": ["Phone ABC", "Phone DEF"]
    },
    {
        "name": "Wireless Earbuds 123",
        "description": "Compact and stylish wireless earbuds offering noise cancellation and excellent sound clarity.",
        "similar_products": ["Earbuds MNO", "Earbuds PQR"]
    }
]
```

---

## Error Handling

- **400 Bad Request**: Invalid or missing image file.
- **405 Method Not Allowed**: Unsupported HTTP methods.
- **500 Internal Server Error**: Issues with the Gemini API or server.

---

## Future Enhancements

- Add authentication for API security.
- Support additional image file formats.
- Implement caching for faster repeated analysis of the same image.

---

## Contributing

Feel free to open issues or submit pull requests. Contributions are welcome!
