build:
	@echo ">  Building driver binary..."
	mkdir -p bin
	go build -a -o ./bin/$* ./cmd/ibm-cosi-driver/$*
test:
unit:
clean:
