# Go Unit Test

Go has a built-in [testing] command `go test` and a package testing

`Go compiler and linker will not ship your test files in any binaries it produces`

The go test tool-chain also includes benchmarking and statement based code coverage similar to [istanbul](Node.js).

## Why Unit test need

Unit tests are crucial to long-term project. We are expected to learn by perceiving, but often we end up dooming ourselves from the start, due to misconceptions or gaps in knowledge. I hope to fill in some of those gaps and provide a broader way of ideas to tackle go unit tests.

Key benefits of unit tests:

- Provide a safety net when refactoring
- Can help identify dead code
- Provide a measure of confidence for management
- Can sometimes find missed use cases
- Define a contract
- Helps produce higher quality code

There are costs associated with writing unit tests as well:

- Time and effort to write and maintain
- False sense of security (poor coverage, duplicate tests, testing the wrong thing, poorly written tests)

## Writing go unit test

Unit testing in Go is just as opinionated as any other aspect of the language like formatting or naming. Go unit test deliberately avoids the use of assertions and leaves the responsibility for checking values.

The requirements for a valid go test file

- File name ends with \_test.go (Ex. `add.go` then the test file should be `add_test.go`)
- Include a package declaration in the file.
- function should begins with the word Test followed by a word or phrase starting with a capital letter and should have only one parameter `t *testing.T`.  
  Ex. `go func TestClientResponse(t *testing.T) {}`

- `t.Error` or `t.Fail` to indicate a failure
- `t.Log` can be used to provide non-failing debug information

### Example

#### Square function we need to test

```go
// main.go
package main

func Square(n int) int {
    return n * n
}

func main() {
    Square(5, 5)
}
```

```go
// main_test.go
package main

import (
    "testing"
)

func TestSquare(t *testing.T) {
    t.Log(Square(2))
}
```

### Executing tests

| Command | Description |
| :-----: | :---------: |
| `go test` | picks up any files matching packagename_test.go                                                   |
|                      `go test github.com/andanhm/gounittest`                      | fully-qualified package name                                                                      |
|                                  `go test ./...`                                  | picks up any files matching \*\_test.go all the packages from directory                           |
|                                   `go test -v`                                    | verbose output with PASS/FAIL result of each test including any extra logging produced by _t.Log_ |
|                                 `go test -cover`                                  | verbose output with code-coverage                                                                 |
| `go test -cover -coverprofile=c.out` `go tool cover -html=c.out -o coverage.html` | generating an HTML coverage report                                                                |

## Subtests

Subtests are built-in to Go. You can target subtests and can nest subtests further if necessary.

```go
func TestSquare(t *testing.T) {

  t.Run(“pass”, func(t *testing.T) {
     if Square(2) != 4 {
         t.Fatal("fail!")
    }
  })
  t.Run(“fail”, func(t *testing.T) {
    if Square(2) != 3 {
        t.Fatal("fail!")
    }
  })
}
```

```sh
go test -run=TestSquare/pass
```

**Go doesn't provide assertions [why?]**

There are third-party libraries [testify], [assert] that replicate the feel of [mocha] a NodeJS unit test library.

## Table based tests

[TBT] "Test based tables" are a way to build a table/array/slice of test input and expected output.
**Uses table-driven tests everywhere**

[![Go unit testing](https://img.youtube.com/vi/yszygk1cpEc/0.jpg)](https://www.youtube.com/watch?v=yszygk1cpEc)

### Example for the writing a unit test with scenario based TBT

```go
func TestSquare(t *testing.T) {
    tests := []struct{
        Input int
        Expected int
    }{
        { 2, 4 },
        { 3, 9 },
    }
    for _, tt := range tests {
        t.Run(fmt.Sprintf("Square(%d)", tt.Input), func(t *testing.T) {
            actual := Square(tt.Input)
            if actual != tt.Expected {
                t.Errorf("expected %d but got %d", tt.Expected, actual)
            }
        })
    }
}
```

- Low overhead to add new test cases
- Makes testing exhaustive scenarios simple. It's easy to see visually if you've covered all cases.
- Makes reproducing reported issues simple

Consider naming the cases in a table-driven test and consider checking negative/error

```go
// tbt/curl_test.go

func TestCurl(t *testing.T) {

    tests := []struct {
        name    string
        url     string
        want    *Response
        wantErr bool
    }{
        {
            "url exist",
            "https://www.andanhm.me/gounittest.json",
            &Response{
                Name:    "gounittest",
                Version: "v1.0.0",
                Status:  true,
            },
            false,
        },
        {
            "url not exist",
            "https://www.andanhm.me/not_exist.json",
            nil,
            true,
        },
        {
            "url provided invalid",
            "andanhm.me/not_exist.json",
            nil,
            true,
        },
        {
            "expected json parser error",
            "https://www.andanhm.me/invalid.json",
            nil,
            true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            response, err := Curl(tt.url)
            if (err != nil) != tt.wantErr {
                t.Errorf("Curl() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(response, tt.want) {
                t.Errorf("Curl() error = %v, want %v", err, tt.want)
                return
            }
        })
    }
}
```

## HTTP Unit testing using http/httptest package

Package httptest provides utilities for HTTP testing.

```go
// client/handlers.go
package handlers

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    // A very simple health check.
    w.WriteHeader(http.StatusOK)

    io.WriteString(w, `Ok`)
}
```

```go
// client/handlers_test.go
package handlers

func TestHealthCheck(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll pass 'nil' as the third parameter.
    request, err := http.NewRequest(http.MethodGet, "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    response := httptest.NewRecorder()
    handler := http.HandlerFunc(HealthCheck)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(response, request)

    // Check the status code is what we expect.
    if status := response.Code; status != http.StatusOK {
        t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `Ok`
    if response.Body.String() != expected {
        t.Errorf("unexpected body: got %v want %v", response.Body.String(), expected)
    }
}
```

## TODO

[ ] Testing Routines and Channels

[ ] Benchmarking

[testing]: https://golang.org/pkg/testing/
[httptest]: https://golang.org/pkg/net/http/httptest
[istanbul]: https://istanbul.js.org/
[mocha]: https://mochajs.org/
[testify]: https://github.com/stretchr/testify
[why?]: https://golang.org/doc/faq#assertions
[assert]: https://github.com/gsamokovarov/assert
