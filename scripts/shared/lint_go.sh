@echo -e "Linting project...\n"
@echo "Running go vet..."
@go vet .
@echo -e "Successfully run go vet.\n"
@echo "Running go fmt..."
@go fmt .
@echo -e "Successfully formatted go fmt.\n"