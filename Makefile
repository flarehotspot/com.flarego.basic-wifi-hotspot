
prod:
	go build -buildmode=plugin -o plugin.so -tags=prod ./main.go
