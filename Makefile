test:
	go test -race -cover ./...

coverprofile=cover.out
webtest:
	go test -race -coverprofile ${coverprofile} ./...
	go tool cover -html ${coverprofile}
