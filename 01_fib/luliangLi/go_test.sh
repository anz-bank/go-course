go test -coverprofile=main_coverage.out

go tool cover -html=main_coverage.out
