# Weather App (WAPP) Project DockerFile

# Parent Image -- 1.23.3 Bookworm (Debian 12)
FROM golang:1.23.3-bookworm

# Create Working Directory
WORKDIR /app

# Copy Source to Working Directory
COPY . .

# Install Dependencies
RUN go mod download

