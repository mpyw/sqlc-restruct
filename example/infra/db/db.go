package db

//go:generate docker compose run --rm -T sqlc generate
//go:generate sqlc-restruct separate-interface --models-file-name=models.gen.go --querier-file-name=querier.gen.go --iface-dir=../../domain/repos --iface-pkg-name=repos --iface-pkg-url=github.com/mpyw/sqlc-restruct/example/domain/repos --models-dir=../../domain/models --models-pkg-name=models --models-pkg-url=github.com/mpyw/sqlc-restruct/example/domain/models
