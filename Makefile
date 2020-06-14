default:
	go build -o $$GOBIN
models:
	easyjson -lower_camel_case -all gofly/models.go
