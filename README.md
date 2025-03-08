# Museum Exhibition API

This is a basic Go project where I learned how to build a server using the Gin framework. The project serves as an API for managing museum exhibitions.

## Project Structure

- **api/**: Contains the API endpoints for GET and POST requests.
- **data/**: Contains the data models and functions for managing exhibition data.
- **public/**: Contains static assets and HTML files for the project.

## Prerequisites

- Go 1.24.1 or later
- Git

## Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd museum
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Running the Server

To start the server, run:

```bash
go run main.go
```

The server will start on `http://localhost:3333`.

## API Endpoints

- **GET /api/exhibitions**: Retrieve all exhibitions or a specific exhibition by ID.
- **POST /api/exhibitions**: Add a new exhibition.

## Static Assets

Static assets are served from the `/api/assets` endpoint.

## Learning Outcomes

Through this project, I learned how to:

- Set up a basic web server using the Gin framework.
- Create RESTful API endpoints.
- Manage data using Go structs and slices.

## License

This project is licensed under the MIT License.

