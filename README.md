# checkout-service

Checkout backend service.

## Install

Make sure you have these installed:

* **Docker:** Get it from [https://www.docker.com/get-started](https://www.docker.com/get-started).
* **Go:** Get it from [https://go.dev/dl/](https://go.dev/dl/).

## How to Run

You can run this service in two ways: using Docker or directly with Go.

### Run with Docker

1.  **Copy `.env.example` to `.env`:**

    Open your terminal in the project folder and run:

    ```bash
    cp .env.example .env
    ```

    (You might need to change values in `.env` later.)

2.  **Build the Docker image:**

    In the project folder, run:

    ```bash
    docker build -t checkout-service .
    ```

3.  **Start the service with Docker Compose:**

    Make sure you have `docker-compose.yaml`. Then run:

    ```bash
    docker-compose up -d
    ```

### Run with Command Line (Go)

1.  **Go to the project folder** in your terminal.

2.  **Run the app:**

    ```bash
    go run app/main.go rest
    ```