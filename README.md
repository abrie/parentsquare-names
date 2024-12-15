# parentsquare-names

## Node Application

This repository contains a Node application that reads a JSON file from the CLI, makes a POST request to `https://www.parentsquare.com/sessions` with the specified cookie and body, and logs the response status code and body.

### How to Run the Application

1. Ensure you have Node.js installed on your machine.
2. Clone this repository.
3. Create a JSON file with the following format and save it as `data.json`:
   ```json
   {
     "cookie": "string_here",
     "body": "string_here"
   }
   ```
4. Run the application using the following command:
   ```sh
   node app.js data.json
   ```

### Expected Input Format

The JSON file should have the following format:
```json
{
  "cookie": "string_here",
  "body": "string_here"
}
```

## Go Application

This repository also contains a Go application that reads a JSON file from the CLI, makes a POST request to `https://www.parentsquare.com/sessions` with the specified cookie and body, and logs the response status code and body.

### How to Run the Application

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Create a JSON file with the following format and save it as `data.json`:
   ```json
   {
     "cookie": "string_here",
     "body": "string_here"
   }
   ```
4. Run the application using the following command:
   ```sh
   go run app.go data.json
   ```

### Expected Input Format

The JSON file should have the following format:
```json
{
  "cookie": "string_here",
  "body": "string_here"
}
```
