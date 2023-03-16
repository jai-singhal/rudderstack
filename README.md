
## Rudderstack Assignment
This repository contains my submission for the Rudderstack Full-Stack SDE 2 assignment

https://rudderstacks.notion.site/Full-Stack-SDE-2-THA-25f2da47127944fd971296b126fad5de.

## How to Run
This is a Dockerized application, which means you can run it using Docker Compose.

1. Clone this repo.
2. Run the following command to build and start the Docker containers:
    
    ```
    docker-compose up --build
    ```

You can access at http://localhost:3000.

Note: Make sure you have Docker and Docker Compose installed on your machine.
If not, you can use these steps

1. Make sure you have mysql running at port 3306
2. Change the DB_HOST in .env/.local to 'localhost' instead of 'db'
3. Run a command in one terminal
    ```
    go build -o rudderstack && go run rudderstack
    ```
4. Run another command in other terminal
    ```
    npm start
    ```