# CAT

This is the unix command line tool Cat, as designed in Coding Challenges.

This project uses the cobra-cli.

## Project Structure

The project is organized into the following components:

```
cccat/
├── main.go               
├── test.txt
├── test2.txt
├── quotes
├── go.sum
├── go.mod
|cmd/                  
│──── root.go
│──── root_test.go               
```

## Components

1. **Main Component** (`main.go`):
   - calls Execute() as part of the cobra-cli tool

2. **Global Flag Handler** (`root.go`):
   - processes commands, flags and Stdin
  
## Testing Approach

The project uses go test for unit testing:

## Running Tests

To run all tests:

```bash
go test -v
```

## Running the tool

```bash
go run main.go test.txt

head -n1 test.txt | go run main.go -

go run main.go test.txt test2.txt

head -n3 test.txt | go run main.go n

sed G test.txt | go run main.go n | head -n4

sed G test.txt | go run main.go b | head -n5
```

