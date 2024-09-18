# Weather API Demo

## Purpose
The Weather API application provides weather information based on latitude and longitude. It is designed to demo a weather forecasts API and related data in a simple, RESTful manner.

## Features
- Retrieve weather data using geographical coordinates.
- Provides temperature, weather conditions, and a description of how the temperature feels.
- Supports CRUD operations for weather data management.

## Design
### Architecture
- **Language**: Go
- **Framework**: Gorilla Mux for routing
- **Middleware**: Custom middleware for JSON responses, logging, and metrics

### Endpoints
- **GET /api/v1/weather/{latitude},{longitude}**: Retrieve weather information for a specific location.
- **PUT /api/v1/weather**: Update weather data (not fully implemented here).
- **POST /api/v1/weather**: Create new weather data (not fully implemented here).
- **DELETE /api/v1/weather**: Delete weather data (not fully implemented here).

### Environment Variables
| Name           | Default Value | Description                                                      |
| -------------- | ------------- | ---------------------------------------------------------------- |
| SERVER_ADDRESS | :8080         | Specifies the server address and port to bind the application to |

## Usage
### Prerequisites
- Go 1.22 installed on your machine.
- Docker installed for containerization.

### Running Locally
1. **Clone the repository**:

```bash
git clone https://github.com/sooonmitch/weather-api-demo
cd weather-api-demo
```
   
2. **Run the application**:
```bash
go run ./
```

3. **Access the API**: Use `curl` or a tool like Postman to interact with the API
```bash
curl "http://localhost:8080/api/v1/weather/40,74"
```

### Running with Docker
1. **Build the Docker image**:
```bash
docker build -t app .
```

2. **Run the Docker container**:
```
docker run -p 8080:8080 -e SERVER_ADDRESS=0.0.0.0:8080 app
```

3. **Test the API**: Again, use curl or Postman:
```bash
curl "http://localhost:8080/api/v1/weather/40,74"
```

## Releasing
Set the release tag:
```bash
git tag -a v1.0.0 -m "<Release Message>"
```

Push the release tag
```bash
git push origin v1.0.0
```

Run gorelease
```bash
goreleaser release
```

## Testing
Run unit tests to ensure the application works as expected:
```bash
go test ./...
```

## Notes
- This is a demo, it will lack features and comforts
- Modify the SERVER_ADDRESS environment variable as needed.
- Ensure that ports used are not conflicting with other applications.