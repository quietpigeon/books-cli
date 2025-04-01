# `book-cli`

```txt
.
├── Makefile
├── README.md
├── api
│   ├── books.go
│   ├── errs.go
│   ├── handlers.go
│   └── handlers_test.go
├── bin
│   ├── book
│   └── server
├── book
│   └── main.go
├── cmd
│   ├── add.go
│   ├── add_test.go
│   ├── cmd.go
│   ├── list.go
│   ├── list_test.go
│   ├── remove.go
│   ├── remove_test.go
│   ├── requests.go
│   ├── update.go
│   └── update_test.go
├── db
│   └── books.db
├── docs
│   ├── api.md
│   └── cli.md
├── go.mod
├── go.sum
└── server
    └── main.go
```

This is a simple implementation of a book management software.

* For the expected user experience, please check out the [user guide](docs/cli.md).
* The documentation of the REST API can be found [here](docs/api.md).

## Usage

There are two binaries for this project, `book` and `server`.

To build the binaries:

```bash
make cli
make server
```

To run all unit tests:

```bash
make tests
```
```
```

To run the server:

```bash
make run-server
```

To install the cli to your path:

 ```bash
 make install-cli
 ```

 (If it doesn't work, please run `go env gobin` and export the path in your shell.)

 (If it still doesn't work, replace `book` with `go run book/main.go` for all of the commands.)

