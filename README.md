# sqlc-restruct

**EXPERIMENTAL!**

Post-processor for [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)

<img width="1048" alt="overview" src="https://github.com/mpyw/sqlc-restruct/assets/1351893/51b422a3-fb7d-4808-a100-ff7c039546b8">

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
go install github.com/mpyw/sqlc-restruct/cmd/sqlc-restruct@latest
```

## SubCommands

### [separate-interface](./cmd/sqlc-restruct/separate_interface.go)

```ShellSession
user@host:~$ sqlc-restruct separate-interface --help
NAME:
   sqlc-restruct separate-interface - Separates models and the `Querier` interface from the `Queries` struct. This is typically done to adhere to the Dependency Inversion Principle (DIP), allowing for more flexible and testable code.

USAGE:
   sqlc-restruct separate-interface [command options] [arguments...]

OPTIONS:
   --iface-pkg-name value     The package name where the separated Querier will be located.
   --iface-pkg-url value      The package URL where the separated Querier will be located. (e.g. "github.com/<user>/<repo>/path/to/pkg")
   --iface-dir value          The directory path where the separated Querier will be located.
   --models-pkg-name value    The package name where the separated models will be located. (default: --models-pkg-name value)
   --models-pkg-url value     The package URL where the separated models will be located. (default: --models-pkg-url value)
   --models-dir value         The directory path where the separated models will be located. (default: --iface-dir value)
   --impl-dir value           The original directory where the sqlc-generated code is located. (default: ".")
   --impl-sql-suffix value    The suffix for sqlc-generated files from SQL files. (default: ".sql.go")
   --models-file-name value   The file name for the sqlc-generated models file. (default: "models.go")
   --querier-file-name value  The file name for the sqlc-generated Querier file. (default: "querier.go")
   --help, -h                 show help
```

We recommend chaining the `//go:generate` directive for `sqlc-restruct` right after the one for `sqlc` in your code.

```go
//go:generate sqlc generate
//go:generate sqlc-restruct separate-interface --iface-dir=domain/repos --iface-pkg-name=repos --iface-pkg-url=github.com/example/domain/repos
```
