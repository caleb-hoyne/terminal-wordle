EXECUTABLE := wordle

build:
	go build -o $(EXECUTABLE)
	chmod +x $(EXECUTABLE)

run:
run: build
	./$(EXECUTABLE)
