# DataDrill
Data Drill – Website Content Scraper (Golang)

Data Drill is a simple yet effective web scraping tool built using Go. It’s designed to fetch and structure key content from any publicly available webpage — including titles, meta descriptions, headings, paragraphs, and links — and return the data as clean, readable JSON.

This project is ideal for developers, content analysts, and SEO researchers who want a lightweight backend service for gathering web content without relying on heavy frameworks or third-party services.

What It Does

Data Drill takes a URL as input and extracts:

Page title

Meta description (if available)

All headings (H1 through H6)

Visible paragraph text

All hyperlinks present on the page

The data is returned via a simple HTTP API in a structured JSON format.

Tech Stack

Language: Go (Golang)

HTTP Server: net/http

HTML Parsing: goquery (jQuery-like syntax)

CORS Support: rs/cors

Project Structure

data-drill/
├── main.go         # Main file with the HTTP server and scraping logic
├── go.mod          # Go module definition
└── README.md       # You're reading it!

Getting Started

1. Clone the Repository

git clone https://github.com/your-username/data-drill.git
cd data-drill

2. Install Dependencies

go mod tidy

3. Run the Server

go run main.go

The API will be live at: http://localhost:8080

API Usage

Endpoint

GET /scrape?url=https://example.com

Example Request

curl "http://localhost:8080/scrape?url=https://example.com"

Example Response

{
  "title": "Example Domain",
  "meta_description": "This domain is for use in illustrative examples in documents.",
  "headings": ["Example Domain"],
  "paragraphs": ["This domain is for use in illustrative examples..."],
  "links": ["https://www.iana.org/domains/example"]
}

Notes & Limitations

Only works on websites that serve static HTML (no JS rendering).

Make sure the website you’re scraping allows it (check robots.txt).

Handles basic errors and invalid input URLs gracefully.

Not designed for bulk scraping or crawling at this stage.

Possible Future Enhancements

Add support for scraping JavaScript-rendered content using headless browsers.

Rate-limiting and retry logic.

Option to export data to CSV or database.

CLI version for local usage.
