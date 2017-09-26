test: test_network test_bnn

test_bnn:
	@( go test ./bnn/ )

test_network:
	@( go test )

goget:
	@( \
		go get github.com/ReconfigureIO/bnn-fpga;
	)

gogetu:
	@( \
		go get -u github.com/ReconfigureIO/bnn-fpga;
	)
