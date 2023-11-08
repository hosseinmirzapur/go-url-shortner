# URL Shortner with Go

This is a Go web application that utilizes Fiber framework, Docker and Redis to create a service to shorten the input urls.

## Features

- **Containerized**: The docker configuration used in this project will allow you to containerize your application
- **Fast**: Uses Redis as database to hold key-value pairs for urls and their short forms, and well... Redis is FAST :)
- **Fiber as web server**: Fiber framework has one of the fastest http routers for Go according to its website

## Requirements

- **go**: should be version 1.20 or higher
- **docker/docker-desktop + docker-compose**: It's recommended (but not necessary) to have Docker Desktop installed for ease of use
- **git**: If not installed, you can download zip file of this project

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/hosseinmirzapur/go-url-shortner.git


   ```

2. Navigate to the project directory:

   ```bash
   cd go-url-shortner


   ```

3. Install the required dependencies:

   ```bash
   go mod download


   ```

4. Run below command to boot the application on docker:

   ```bash
   docker-compose up -d

   # you can also run the command as sudo-user, also non-detached without the "-d" flag

   ```

## Usage

The web application will contain 2 endpoints:

- `/:url`: Tries to retrieve the provided url from database. (GET REQUEST) Example usage:

  ```bash
  curl http://localhost:3000/short-url-entered-by-user

  ```

- `/api/v1`: Generates a short url for the url provided by user in the request body. `check shorten.go file`


## Configuration

All configurable environment values are available inside the .env.example file, be sure to check it out.

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please create a new issue on the project's GitHub repository. Pull requests are also encouraged.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- This web app was developed using the Golang programming language.
- It utilizes Docker, Fiber Framework, and Redis APIs for its purpose.

## Contact

For any inquiries or questions, please contact me via [email](mailto:hosseinmirzapur@gmail.com).

---

Feel free to customize this README.md file according to your specific project requirements and add any additional sections or information that you think would be relevant.