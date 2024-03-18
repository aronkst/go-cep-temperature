# Go CEP Temperature

## Overview

The project offers an advanced and practical solution for accessing detailed weather information by using Postal Codes (CEPs) as the basis for the query. Simply by providing a CEP through the URL /?cep=CEP, users receive a quick response in JSON format, which includes current temperatures expressed in the three main thermometric scales: Celsius, Fahrenheit, and Kelvin.

## Features

- **Direct Query by CEP**: Allows quick access to specific temperature information of a location, using only the Postal Code (CEP) as a reference.
- **Rigorous CEP Validation**: Implements validation to ensure that the provided CEP is in the correct format, consisting only of numbers and containing exactly 8 characters, ensuring the accuracy of the queries.
- **Free Authentication**: The system is designed to be accessible without the need for API tokens or any form of authentication, simplifying access to weather information.
- **Responses in JSON Format**: All responses are provided in JSON format, facilitating integration with other applications and data handling.
- **Support for Multiple Temperature Units**: Provides temperatures in Celsius, Fahrenheit, and Kelvin, offering flexibility to meet the preferences and needs of users.
- **Real-Time Weather Updates**: Integrates with reliable meteorology services to provide updated and accurate weather information.

## Usage Example

To consult weather information through the command line, you can use `curl`, a powerful tool available on most operating systems for making HTTP requests. Below are practical examples of how to use curl to obtain the temperature based on a specific CEP.

### Making a Query

To make a query, simply replace CEP with the desired postal code in the URL. Here are some examples:

```bash
curl "http://localhost:8080/?cep=01001000"
```

Expected return:

```json
{"temp_C":27.9,"temp_F":82.22,"temp_K":301.05}
```

In this example, the request returns the temperature for CEP 01001000 (a CEP from São Paulo), showing the temperature in Celsius (temp_C), Fahrenheit (temp_F), and Kelvin (temp_K).

```bash
curl "http://localhost:8080/?cep=80210090"
```

Expected return:

```json
{"temp_C":25.3,"temp_F":77.54,"temp_K":298.45}
```

Here, the temperature is returned for CEP 80210090, which corresponds to a location in Curitiba.

### How Data is Returned

Data is returned in JSON format. Each field in the JSON represents a different temperature measure:

- `temp_C`: Temperature in degrees Celsius.
- `temp_F`: Temperature in degrees Fahrenheit.
- `temp_K`: Temperature in Kelvin.

## Development

I developed the project focusing on the coordinated use of several external APIs to deliver accurate weather information based on a provided Postal Code (CEP). The process to obtain this information follows a logical sequence of steps, where each one makes use of a specific API to achieve the final goal. Below, I detail each step and how each API is employed:

### Address Lookup by CEP with viacep.com.br

The starting point involves collecting detailed information about the address associated with the provided CEP. I use the ViaCEP API for this purpose. Upon receiving a valid CEP, I make a request to the ViaCEP API, which returns data such as street, neighborhood, city, and state corresponding to the CEP. These data are essential for determining the exact geographical location to be used in the following weather queries.

### Longitude and Latitude Lookup with nominatim.openstreetmap.org

Having the address data, the next step is to convert them into geographical coordinates (latitude and longitude). For this, I turn to the Nominatim API, part of the OpenStreetMap project. This API allows me to send location details, such as city and state, and receive in return the precise geographical coordinates of that location. This conversion is vital to ensure the accuracy of weather queries based on coordinates.

### Temperature Lookup with api.open-meteo.com or wttr.in

Having the geographical coordinates, I move on to the phase of querying the current weather conditions. At this point, the process divides, depending on the availability of coordinates:

- With longitude and latitude data available, I turn to the Open-Meteo API. This API allows for detailed weather queries based on geographical coordinates, providing exact temperature data for the desired location.
- Without longitude and latitude data, I use the wttr.in API. This API provides weather information based on location names (such as cities), derived from the data obtained via ViaCEP. Although this method may not be as precise as querying by coordinates, it still offers a useful estimate of weather conditions.

## Error Handling

I implemented error handling at each step to ensure that the system can adequately deal with scenarios such as invalid CEPs, failures in obtaining coordinates, or errors in the responses from the weather APIs.

## Unit Tests

An integral part of developing this project involves the implementation of comprehensive unit tests, ensuring the reliability and robustness of each functionality offered by the application. The approach taken for testing follows best software development practices, focusing on validating each component in isolation to ensure its correct operation in various scenarios.

### Test Coverage

The unit tests cover a wide range of use cases and error scenarios, including, but not limited to:

- CEP Validation: Tests to ensure that only valid CEPs in the correct format are accepted, and that appropriate error messages are returned for invalid or incorrectly formatted CEPs.
- External API Queries: Tests to verify the correct interaction with the external APIs used to obtain address information, geographical coordinates, and weather data. This includes simulating API responses to test the proper handling of data and errors.
- Temperature Unit Conversion: Tests that validate the accuracy of temperature conversions between Celsius, Fahrenheit, and Kelvin, ensuring the calculations are correct.
- Error Handling: Specific tests to check the system's robustness when facing errors during information querying, including network failures, external API errors, and unexpected data.

## Makefile

This project includes a Makefile designed to offer an efficient and simplified interface for managing development and production environments, as well as executing automated tests. The commands provided allow optimizing and streamlining the development workflow, testing, and maintenance of the project, ensuring more effective and organized management. It is recommended to use these commands to maximize productivity and ensure consistency throughout the software lifecycle.

### Development Commands

### `make dev-start`

Starts the services defined in the `docker-compose.yml` file for the development environment in detached mode (in the background). This allows the services to run in the background without occupying the terminal.

### `make dev-stop`

Stops the services running in the background for the development environment. This does not remove the containers, networks, or volumes created by `docker compose up`.

### `make dev-down`

Shuts down the development environment services and removes the containers, networks, and volumes associated created by `docker compose up`. Use this command to clean up resources after development.

### `make dev-run`

Starts the application's execution within the development environment, using Docker Compose to execute the `go run` command in the `cmd/server/main.go` file. It is ideal for quickly starting the project server in development mode.

### `make dev-run-tests`

Runs all Go tests within the specified container (`dev-go-cep-temperature`), showing verbose details of each test. This command is useful for running the project's test suite and verifying everything is working as expected.

### Production Commands

### `make prod-start`

Starts the services defined in the `docker-compose.prod.yml` file for the production environment in detached mode. This is useful for running the project in an environment that simulates production.

### `make prod-stop`

Stops the production environment services running in the background, without removing the associated containers, networks, or volumes.

### `make prod-down`

Shuts down the production environment services and removes the associated containers, networks, and volumes, cleaning up resources after use in production.

## Prerequisites

Before starting, make sure you have Docker and Docker Compose installed on your machine. If not, you can download and install them from the following links:

- Docker: https://docs.docker.com/get-docker/

### Clone the Repository

First, clone the project repository to your local machine. Open a terminal and execute the command:

```bash
git clone https://github.com/aronkst/go-cep-temperature.git
```

### Navigate to the Project Directory

After cloning the repository, navigate to the project directory using the cd command:

```bash
cd go-cep-temperature
```

## Development Environment

### Build the Project with Docker Compose

In the project directory, execute the following command to build and start the project using Docker Compose:

```bash
docker compose up --build
```

Or using the Makefile:

```bash
make dev-start
```

This command will build the Docker image of the project and start the container.

### Run the Project with Docker Compose

To start the main service of your project in development mode, you can use the direct Docker Compose command:

```bash
docker compose exec go run cmd/server/main.go
```

Or using the Makefile:

```bash
make dev-start
```

### Access the Project

With the container running, you can access the project through the browser or using tools like curl, pointing to http://localhost:8080/?cep=CEP, replacing CEP with the desired postal code.

### Example curl Command

To test if the project is running correctly, you can use the following curl command in a new terminal:

```bash
curl "http://localhost:8080/?cep=01001000"
```

You should receive a JSON response with temperatures in Celsius, Fahrenheit, and Kelvin.

### Ending the Project

To end the project and stop the Docker container, return to the terminal where Docker Compose is running and press Ctrl+C. To remove the containers created by Docker Compose, execute:

```bash
docker compose down
```

Or using the Makefile:

```bash
make dev-down
```

## Production Environment

### Build and Run the Project with Docker Compose

In the project directory, execute the following command to build and start the project in the production environment using Docker Compose:

```bash
docker compose -f docker-compose.prod.yml up --build
```

Or using the Makefile:

```bash
make prod-start
```

This command will build the Docker image of the project for production and start the containers.

### Example curl Command

To check if the project in production is operational, use the following curl command, adjusting the address as per your setup:

```bash
curl "http://localhost:8080/?cep=01001000"
```

You should receive a JSON response with the requested information, such as temperatures in Celsius, Fahrenheit, and Kelvin.

### Ending the Project

To end the project and stop the production containers, use the following command:

```bash
docker compose -f docker-compose.prod.yml down
```

Or using the Makefile:

```bash
make prod-down
```

This command shuts down all production services and removes the associated containers, networks, and volumes, ensuring the production environment is cleaned up after use.
