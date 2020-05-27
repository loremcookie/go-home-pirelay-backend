@echo "Compiling for raspberry pi..."
@env GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/server
@echo "Succesfully compiled"