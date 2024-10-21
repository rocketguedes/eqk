# This application displays Earthquakes that ocurred in the last 30 days.

Autor: [Marcelo Pinheiro](http://twitter.com/mpinheir)

## Data Source
[USGS](https://earthquake.usgs.gov/)  

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

- Go installed on your machine
- Internet connection to fetch earthquake data from the USGS API

### Build

1. Clone the repository:

   ```bash
   git clone https://github.com/mpinheir/eqk.git
   ```

2. Change into the project directory:

    ```bash
    cd eqk
    ```

3. Build the binary:

    ```bash
    go build
    ```

4. Run the program:

    ```bash
    go run main.go <arg>
    ```

    ex: `go run main.go 6` will display earthquake(s) greater than 6 degrees

5. Run the binary after building:

    ```bash
    ./eqk <arg>
    ```

    ex: `./eqk 5` will display earthquake(s) greater than 5 degrees


    ```
    -------------------------------------------------------------------
    Earthquake(s) above 5.0 degrees, in the last 30 day:
    -------------------------------------------------------------------
    Epicenter = [Location]
    Magnitude: [Magnitude]
    Time: [Timestamp]
    -------------------------------------------------------------------
    ```

## Contributing
Contributions to this project are welcome! Plese feel free to open issues and pull requests to suggest improvements, report bugs, or add new features.

## License
This project is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License).
