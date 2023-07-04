# sqlc-restruct

Post-processor for [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)

<img width="1048" alt="overview" src="https://github.com/mpyw/sqlc-restruct/assets/1351893/4aa8cc38-538f-442e-ba70-990226e9623f">

```ShellSession
user@host:~$ sqlc-restruct --help
NAME:
   sqlc-restruct - Post-processor for kyleconroy/sqlc

USAGE:
   sqlc-restruct [global options] command [command options] [arguments...]

COMMANDS:
   separate-interface  Separates models and the `Querier` interface from the `Queries` struct. This is typically done to adhere to the Dependency Inversion Principle (DIP), allowing for more flexible and testable code.
   help, h             Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Installation

```
go install github.com/mpyw/sqlc-restruct@v0.0.0-alpha.0
```

## SubCommands

### [separate-interface](./cmd/separate_interface.go)

```ShellSession
user@host:~$ sqlc-restruct --help
NAME:
   sqlc-restruct separate-interface - Separates models and the `Querier` interface from the `Queries` struct. This is typically done to adhere to the Dependency Inversion Principle (DIP), allowing for more flexible and testable code.

USAGE:
   sqlc-restruct separate-interface [command options] [arguments...]

OPTIONS:
   --iface-pkg-name Querier     The package name where the separated models and Querier will be located.
   --iface-pkg-url Querier      The package URL where the separated models and Querier will be located (e.g. "github.com/<user>/<repo>/path/to/pkg").
   --iface-dir Querier          The directory path where the separated models and Querier will be located.
   --impl-dir value             The original directory where the sqlc-generated code is located. (default: ".")
   --impl-sql-suffix value      The suffix for sqlc-generated files from SQL files. (default: ".sql.go")
   --models-file-name value     The file name for the sqlc-generated models file. (default: "models.go")
   --querier-file-name Querier  The file name for the sqlc-generated Querier file. (default: "querier.go")
   --help, -h                   show help
```

We recommend chaining the `//go:generate` directive for `sqlc-restruct` right after the one for `sqlc` in your code.

```go
//go:generate sqlc-generate
//go:generate sqlc-restruct separate-interface --models-file-name=models.gen.go --querier-file-name=querier.gen.go --iface-dir=domain/repos --iface-pkg-name=repos --iface-pkg-url=github.com/example/domain/repos
```
