package db

//go:generate docker compose run --rm -T sqlc generate
//go:generate go run ../../../cmd/sqlc-restruct separate-interface --impl-dir=articles --models-file-name=entities.articles.gen.go --querier-file-name=querier.articles.gen.go --iface-dir=../../domain/repos/articles --iface-pkg-name=articles --iface-pkg-url=github.com/mpyw/sqlc-restruct/example/domain/repos/articles --models-dir=../../domain/entities --models-pkg-name=entities --models-pkg-url=github.com/mpyw/sqlc-restruct/example/domain/entities
//go:generate docker compose run --rm -T sqlc generate
//go:generate go run ../../../cmd/sqlc-restruct separate-interface --impl-dir=users --models-file-name=entities.users.gen.go --querier-file-name=querier.users.gen.go --iface-dir=../../domain/repos/users --iface-pkg-name=users --iface-pkg-url=github.com/mpyw/sqlc-restruct/example/domain/repos/users --models-dir=../../domain/entities --models-pkg-name=entities --models-pkg-url=github.com/mpyw/sqlc-restruct/example/domain/entities
