test: test_bnn

test_bnn:
	@( go test ./bnn/ )

goget:
	@( \
		go get github.com/ReconfigureIO/brain;
	)

gogetu:
	@( \
		go get -u github.com/ReconfigureIO/brain;
	)
