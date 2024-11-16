# GoWeatherLookup
"GoWeatherLookup" is a Go-based service that receives a valid ZIP code, identifies the city, and returns current weather in Celsius, Fahrenheit, and Kelvin. 

## Requirements

- The system must receive a valid zip code of 8 digits.
- The system should look up the location name using the provided zip code and return the temperatures formatted as:
  - Celsius (°C)
  - Fahrenheit (°F)
  - Kelvin (K)
- The system must handle responses appropriately in the following scenarios:

### Success Response:
- **HTTP Status Code**: `200`
- **Response Body**:
  ```json
  {
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.65
  }
### Failure Responses:
- Invalid Zip Code (correct format but not valid):
  - **HTTP Status Code**: `422`
  - **Message**: "invalid zipcode"
- Zip Code Not Found:
  - **HTTP Status Code**: `404`
  - **Message**: "can not find zipcode"

## Implementation Details

- **viaCEP API** was used to find the city based on the provided zip code:
  - [viaCEP API](https://viacep.com.br/)

- **WeatherAPI** was used to return the temperatures in different units (Celsius, Fahrenheit, Kelvin):
  - [WeatherAPI](https://www.weatherapi.com/)

## Usage

1. Copy `docker-compose.yaml.dist` to `docker-compose.yaml`:
   ```bash
   cp docker-compose.yaml.dist docker-compose.yaml
   ```

2. Copy `env.dist` to `env`:
   ```bash
   cp env.dist env
   ```

3. Add your **WeatherAPI key** to the `env` file.

4. Build and run the application using Docker Compose:
   ```bash
   docker compose up --build
   ```

5. Test requests are available in the `api` folder for convenience.

