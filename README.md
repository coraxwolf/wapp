# Weather App

Weather App is a simple web-based application that allows users to check the current weather in their location and look up weather in other areas. This project is build on Go 1.23.3 using HTMX for the client frontend.

## Features
* Local Weather Display: Automaticly shows current conditions to users for their "home" area
* Location-Based Weather Search: Users can search for the weather conditions and forcasts for areas by either city name, zip code, or geolocation (Lat & Lon coordinates)
* Live Updating: Site will update to refresh the current display to show updates to the current conditions.
* Smooth & Responsive: HTMX Technology is used to all the page the refresh data/content without requiring a full page reload.

## Technology Stack
* Go 1.23.3
* HTMX 2.0.3
* OpenWeatherMaps.org API

## Getting Started
### Prerequisites
* Go 1.23.3
* API Key from [OpenWeatherMap.org](https://home.openweathermap.org)

### Install Repo
1. Clone Repositiory
```
git clone https://github.com/coraxwolf/wapp.git
cd wapp
```

2. Environment Variables (in .env file in root)
```
ADDR=8888 # Address for Site
STATIC_PATH=path/to/static/files # (relative to Data Path)
HTTP_LOG=true                    # (log http requests true or false)
HTTP_LOG_NAME=httpfile.log       # (log file for http requests)
OW_TOKEN=yourapikey              # (OpenWeatherMap API Key)
LOG_LEVEL=info                   # (use: info/error/warn/debug/none)
LOG_FILENAME=logfile.log         # (Name of Log File)
LOG_PATH=path/to/logs            # (relative to Data Path)
LOG_CONSOLE=true                 # (log to console or not (false))
LOG_FMT_JSON=false               # (log format in json (true) or text (false))
DATA_PATH=path/to/data/files
```

### Build Commands
* **make clean**: Clean up Build Files
* **make css**: NPM Script to generate tailwindcss css files
* **make templates**: Generate TEMPL template files
* **make build-dev**: Build Dev Binary
* **make run**: Build and Run Dev Binary
* **make test**: Run Tests
* **make cover**: Generate Test Coverages by Functions

### File Structure
* **cmd**: Commands
* **cmd/app**: Main App
* **cmd/httpd**: Web (httpd) Server -- using net/http library
* **handlers**: Route Handlers
* **middleware**: Middleware
* **models**: Data/Repository Models
* **viewmodels**: View Models for the web views
* **views**: TEMPL templates
* **internal**: Internal Dependencies
* **internal/config**: App Configuration/Setup
* **internal/database**: Database Service
* **internal/owapi**: OpenWeatherMaps API Service
* **internal/logger**: Logger Service
* **src**: CSS & JavaScript source files
* **data**: Local data store. Files not synced to repo, but used by running process
* **data/assets**: Static Folder for web assets (css, js, images, fonts)
* **logs**: Log Files

## Development
### TDD
This project is developed using Test Driven Development practices

## License
This project is licensed under the [MIT License](LICENSE.md)