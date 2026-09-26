
all: run-backend-dev

run-backend-dev:
	echo "RUNNING BACKEND SERVER: DEVELOPMENT" && \
	cd cmd && \
	go run main.go

start-local-db:
	npx supabase start

stop-local-db:
	npx supabase stop

run-sorm:
	go run SORM/src/main.go

run-migration-tests:
	cd SORM/src/migrations && go test -v && cd entries && go test -v