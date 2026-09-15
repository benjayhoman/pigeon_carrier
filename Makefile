OUTPUT = pigeon_carrier

build:
	go build -o $(OUTPUT) ./cmd/main.go

clean:
	rm -f $(OUTPUT)

run: build
	./$(OUTPUT)