# go-road-fare

The [Volvo Cars congestion tax calculator challenge](https://github.com/volvo-cars/congestion-tax-calculator/blob/main/ASSIGNMENT.md), solved in Go.

## Overview

The challenge was to develop an application for calculating congestion tax fees for vehicles. This REST API provides a straightforward endpoint to obtain congestion tax information.

Currently, the API is configured to calculate congestion tax fees based on the rules of the city of Gothenburg. It supports various vehicle types and time ranges, offering precise calculations within the defined tax regulations.

### API Endpoints

- `/api/calculator/{vehicle}/{time}`

`vehicle`: Type of the vehicle (e.g., car, motorbike).

`time`: Time range specifying when the vehicle entered and exited a congestion tax area.

### Example

`GET /api/calculator/car/2023-05-05T12:00:00,2023-05-05T18:15:00`

## License

This project is licensed under the [MIT License](LICENSE).
