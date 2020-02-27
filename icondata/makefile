# Go parameters
    GOCMD=go
    GOINSTALL=$(GOCMD) install
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=icondata
    BINARY_MACOS=$(BINARY_NAME)_macos
    BINARY_UNIX=$(BINARY_NAME)_unix
    BINARY_WIN=$(BINARY_NAME)_win.exe
    
    all: test build
    test: 
	$(GOTEST) -v ./...
    install:
	$(GOINSTALL) -v
    clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
    run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
    
    # Cross cleaning
    clean-all: clean-windows clean-macos clean-linux
    clean-windows: 
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOCLEAN) -i
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_WIN)
    clean-macos: 
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOCLEAN) -i
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_MACOS)
    clean-linux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCLEAN) -i
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

    # Cross compilation
    install-all: install-windows install-macos install-linux
    install-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOINSTALL) -v
    install-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOINSTALL) -v
    install-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOINSTALL) -v