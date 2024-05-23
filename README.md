# Bitcoin Transaction Timestamp Service

This Go application provides a web service to fetch the last active timestamp of a Bitcoin address. It retrieves the latest transaction IDs associated with the given address and then fetches the timestamp of the most recent transaction.

## Features

- Home page with basic instructions.
- Endpoint to get the last active timestamp of a Bitcoin address.

## Endpoints

- `GET /`: Home page with instructions.
- `GET /api/address/{btcaddress}`: Fetches the last active timestamp for the specified Bitcoin address.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Abhinav7903/TestAPI_BTC.git
    cd TestAPI_BTC
    ```

2. Install the required packages:

    ```sh
    go get -u github.com/gorilla/mux
    ```

### Running the Application

1. Build and run the application:

    ```sh
    go run main.go
    ```

2. The server will start listening on port `8080`. You should see the following output:

    ```sh
    Server listening on port 8080
    ```

### Usage

- To access the home page, open your browser and navigate to `http://localhost:8080/`.

    You should see a message saying:

    ```
    Hello! This is the home page.
    To get the last active timestamp for a Bitcoin address, make a GET request to /api/address/{btcaddress}
    Example: /api/address/1a2b3c4d5e6f7g8h9i0j
    ```

- To get the last active timestamp for a specific Bitcoin address, make a GET request to `http://localhost:8080/api/address/{btcaddress}`. Replace `{btcaddress}` with the actual Bitcoin address.

    Example request:

    ```sh
    curl http://localhost:8080/api/address/1a2b3c4d5e6f7g8h9i0j
    ```

    Example response:

    ```json
    {
      "message": "success",
      "last_active_timestamp": "2023-05-23 15:04:05"
    }
    ```

## Code Overview

### Main Components

- **Main Function**: Sets up the router and starts the server.
- **Home Handler**: Provides basic instructions on the home page.
- **Timestamp Handler**: Fetches the last active timestamp for a Bitcoin address.

### Error Handling

The application handles various error scenarios, such as failing to fetch data from the external API or failing to decode the response. Appropriate error messages are returned in the response.

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux): A powerful URL router and dispatcher for Go.

## Acknowledgements

- [Gorilla Mux](https://github.com/gorilla/mux) for the routing library.

