# Beef Counter API

An API application for counting different types of beef mentioned in text from the BaconIpsum API.

## Overview

This application is designed to count the occurrences of various beef types in text fetched from the BaconIpsum API, which provides sample text with various meat-related terms. The API filters out only beef-related terms and returns the count of each beef type found.

## Features

- Fetches sample text from the BaconIpsum API
- Filters and counts each type of beef mentioned in the text
- Case-insensitive matching
- Ignores punctuation marks in the text
- Result caching for improved performance
- HTTP API using the Gin framework

## Project Structure

This project follows a standard structure according to modern Go development practices:

```
pie-fire-dire/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── app/
│   │   └── app.go               # Application controller
│   ├── handler/
│   │   └── beef_handler.go      # HTTP handlers
│   ├── model/
│   │   └── beef.go              # Data models
│   └── service/
│       ├── beef_service.go      # Beef management service
│       └── meat_service.go      # Meat data fetching service
├── pkg/
│   └── util/
│       └── text_processor.go    # Text processing utility functions
└── test/
    └── integration_test.go      # Integration tests
```

## Getting Started

### Prerequisites

- Go version 1.18 or newer
- Git

### Installation

1. Clone the project:
```bash
git clone https://github.com/rattanun00000/pie-fire-dire.git
cd pie-fire-dire
```

2. Install dependencies:
```bash
go mod tidy
```

### Running the Application

```bash
go run cmd/server/main.go
```

Or build an executable:
```bash
go build -o pie-fire-dire cmd/server/main.go
./pie-fire-dire
```

The application will start on port 8080 by default.

### Testing

Run all tests:
```bash
go test ./...
```

Run integration tests:
```bash
go test -v ./test/
```

Run tests with coverage:
```bash
go test ./... -cover
```

## Using the API

### Available Endpoints

- **GET** `/beef/summary`: Returns the count of each beef type found in the sample text

### Example Request

```bash
curl -X GET http://localhost:8080/beef/summary
```

### Example Response

```json
{
  "beef": {
    "beef": 12,
    "t-bone": 4,
    "ribeye": 7,
    "sirloin": 5,
    "brisket": 3,
    "chuck": 2,
    "tenderloin": 1,
    "filet mignon": 3,
    "short ribs": 2,
    "ground round": 1,
    "strip steak": 3,
    "flank": 2,
    "tri-tip": 4,
    "corned beef": 1
  }
}
```

## Design and Decisions

### Code Segregation
The code is divided into different components following Clean Architecture and Separation of Concerns principles, providing a clear structure that is easy to understand and maintain.

### Performance
- Results from API calls are cached for 5 minutes to reduce unnecessary external API calls
- Efficient Regular Expressions for word searching
- Beef types are sorted by length (longest to shortest) to prevent overlapping matches

### Error Handling
- Proper handling of potential errors during external API calls
- Appropriate status and messages returned in case of errors