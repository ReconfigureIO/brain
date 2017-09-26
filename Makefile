test: test_bnn

test_bnn:
	@( go test ./src/bnn/ )

goget:
	@( \
		go get github.com/ReconfigureIO/bnn-fpga;
	)

gogetu:
	@( \
		go get -u github.com/ReconfigureIO/bnn-fpga;
	)
