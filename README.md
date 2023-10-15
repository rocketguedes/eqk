# Earthquake Data Fetcher in Go

This Go program fetches and displas Earthquake data.

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

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mpinheir/Terremoto-em-go.git

2. Change into the project directory:

    ```bash
    cd Terremoto-em-go

3. Run the program:
    ```bash
    go run main.go

## Usage

The program fetches data from the USGS Earthquake API and displays information about them, including location, magnitude, and time.

```bash
-------------------------------------------------------------------
Earthquakes above 6 Richter scale in the last 30 days:
-------------------------------------------------------------------
Epicenter = [Location]
Magnitude: [Magnitude]
Time: [Timestamp]
-------------------------------------------------------------------
```

## Contributing
Contributions to this project are welcome! Feel free to open issues and pull requests to suggest improvements, report bugs, or add new features.

## License
This project is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License).
