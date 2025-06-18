# Bible Server Project

This project is a simple web server built with Go that serves a random chapter and its verses from a Bible stored in an SQLite database. The server displays the book name and chapter at the top, followed by the verses.

## Project Structure

```
.
├── bible
│   └── ARC.sqlite
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
├── static
│   └── style.css
└── templates
    └── index.html
```

- **bible/ARC.sqlite**: The SQLite database containing the Bible.
- **Dockerfile**: The Docker configuration file for building the application image.
- **go.mod** and **go.sum**: Go module files for dependency management.
- **main.go**: The main Go application file.
- **static/style.css**: CSS file for styling the web page.
- **templates/index.html**: HTML template for displaying the Bible verses.

## Prerequisites

- [Go](https://golang.org)
- [Docker](https://www.docker.com)

## Setup and Running

### Running Locally

1. **Clone the repository**:
```bash
   git clone <repository-url>
   cd bible-server
```

2.  **Run the Docker container**:
    
```bash
docker build -t bible-server .
docker run -p 8080:8080 bible-server
```
**OR** 

2.  **Run the server**:
    
```bash
go run main.go
```
    
3.  **Access the server**: Open your web browser and go to  `http://localhost:8080`.

## Features

-   Displays a random Bible chapter and its verses.
-   Modern, responsive web design.