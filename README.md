GitHub Gist API Proxy
Go-based API built with the Gin framework. This service acts as a proxy to fetch a user's publicly available Gists from GitHub.

### Prerequisites
To run this application locally without Docker, you will need:

Go: Version 1.25 or higher.

Internet Connection: Required to fetch data from the GitHub API.

Docker: (Optional) For containerized execution.

### How to Run the App
Initialize & Install Dependencies:

Bash
go mod init gists-api
go mod tidy
Start the Server:

Bash
go run main.go
The server will start on http://localhost:8080.

### How to Test
The project includes automated tests to validate the integration with GitHub.

Run all tests:

Bash
go test -v .
Note: The tests require internet access as they validate the connection using the real octocat user profile.

### API Usage (curl)
To fetch Gists for a specific user, use the following command in your terminal:

Bash
curl http://localhost:8080/octocat
Example Response:

JSON
[
  {
    "html_url": "https://gist.github.com/octocat/6a72c1...",
    "description": "Hello World Examples"
  }
]
### Docker Implementation
#### Build the Image
This project uses a multi-stage build for maximum security. It compiles the app in an Alpine environment and then moves the binary to a scratch (empty) image to eliminate the attack surface.

Bash
docker build -t gist-api .
#### Run the Container
The API is configured to listen on port 8080.

Bash
docker run -p 8080:8080 gist-api
### Security Features Included
Multi-Stage Build: Reduces image size and prevents source code leakage.

Distroless (Scratch): The final image contains no shell or OS utilities, preventing attackers from using tools like sh or ls.

Non-Root User: The application runs as a restricted appuser.

Static Binary: Compiled with CGO_ENABLED=0 to prevent dependency injection.

Stripped Binary: Uses -ldflags="-w -s" to remove debug symbols.

### Technical Choices
Gin Gonic: Chosen for its fast routing and built-in JSON handling.

HTTP Timeouts: Configured with a 5-second timeout to prevent the server from hanging on slow upstream responses.

Struct Mapping: Uses Go structs with JSON tags to provide a clean, filtered contract to the end user.