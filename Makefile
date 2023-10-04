export MIGRATION_DIR="./infra/db/migrations"
export COVERAGE_PATH=./coverage.out

test:
	go test ./... -cover -coverprofile=$(COVERAGE_PATH) && go tool cover -html=$(COVERAGE_PATH)