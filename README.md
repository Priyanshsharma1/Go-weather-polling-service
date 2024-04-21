# Weather Data Poller

This Go program fetches weather data for multiple cities from the OpenWeatherMap API and displays the temperature, feels-like temperature, and humidity at regular intervals.

## Setup

1. Clone the repository to your local machine.
2. Create a `.env` file in the root directory and add your OpenWeatherMap API key as follows:

    ```
    API_KEY=<your-api-key>
    ```

3. Build and run the program.

## Configuration

- `API_KEY`: Your OpenWeatherMap API key. You can obtain one by signing up on [OpenWeatherMap](https://openweathermap.org/).
- `API_ENDPOINT`: The API endpoint for fetching weather data. Default is `http://api.openweathermap.org/data/2.5/weather`.
- `INTERVAL`: Interval at which the program fetches weather data for each city. Default is 5 seconds.
- `CITIES`: List of cities for which weather data will be fetched.

## Usage

Run the program and observe the weather data printed to the console.

## Dependencies

- [github.com/joho/godotenv](https://github.com/joho/godotenv): GoDotEnv is used for loading environment variables from a `.env` file.
- Standard library packages: `encoding/json`, `fmt`, `net/http`, `os`, and `time`.

