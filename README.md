# Go ChatGPT Prompt Splitter

This project is a web-based application that allows users to input a long text 
string and split it into smaller parts based on a user-defined length. 
The project uses Go for server-side processing and a simple HTML & JavaScript
frontend for user interaction. The application uses Redis for data management
and Docker for containerization and deployment.

This project is a Go port from Python [ChatGPT Prompt Splitter](https://github.com/jupediaz/chatgpt-prompt-splitter) from jupediaz.

## Features
- Input text prompt and desired length to split the prompt into smaller parts.
- Display the number of site visits using a counter.
- Copy each part of the split text to the clipboard with a click of a button.

- ## Requirements
- Go 1.18 or later
- Docker and Docker Compose
- Redis in your local environment or a Redis server URL if not running docker

## Installation
 
1. Clone this repository:

```bash
git clone https://github.com/yourusername/go-chatgpt-prompt-splitter.git
```

2. Change into the project directory:
```bash
cd go-chatgpt-prompt-splitter
```

3. Install the Go dependencies:
```bash
go mod tidy
```

## Usage
### Running Locally
1. Run the server:
```bash
make run
```

The server will start and listen on port 8080 by default. If you want to use a
different port, you can specify the PORT environment variable in the .env file.

2. Open a web browser and navigate to http://localhost:8080 (or the port you 
specified).

![Home Page of Local Go ChatGPT Splitter Server](static/images/go-chatgpt-prompt-splitter-start.png)

3. Enter your long text in the "Prompt" textarea and specify the length of each
part in the "Split Length" field.
   

![Copy the long prompt to the text field](static/images/go-chatgpt-prompt-splitter-input-text.png)

4. Click "Split the text" to split the prompt. The split parts will be displayed
below the form.

![Prompt got splitted based on given chunk length](static/images/go-chatgpt-prompt-splitter-result.png)

5. Click the "Copy to clipboard" button next to each part to copy the part to
the clipboard.

### Running with Docker
1. Build and run the Docker containers:
```bash
make run-docker-build
```
This will start the Go application and Redis server in separate Docker containers.

2. Follow the same steps as above to use the application.

## Configuration
This project uses the UPSTASH_REDIS_URL and PORT environment variables. 
You can specify these variables in the .env file:

```makefile
# URL for the Redis server
UPSTASH_REDIS_URL="redis://redis:6379"

# Port for the application to listen on
PORT=8080
```

Replace redis://redis:6379 with your Redis URL if needed, and 8080 with the port you want the application to listen on.

## Roadmap

Our project is continuously improving and expanding! Here are some of the exciting updates and enhancements we're planning:

- [ ] Add tests coverage for handlers
- [ ] Replace Redis with SQLite: Transition from Redis to SQLite for simplified and efficient data management.
- [ ] Support for PDF files: Introduce ability to upload PDFs, convert to text, and split the text into manageable parts.
- [ ] Improve the UI with Flutter Web: Upgrade front-end to use Flutter Web for a more dynamic, responsive user interface.
- [ ] Connect to OpenAPI API, user can select and use any API from OpenAPI with either GPT-3.5 or GPT-4 and get response directly. 
- [ ] Continuous Improvements: Always open to enhancements and new feature suggestions to better serve our users.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

### License

[MIT](https://choosealicense.com/licenses/mit/) License