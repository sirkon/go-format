test:
		PATH=${GOPATH}/bin:${PATH}
		go get -u github.com/stretchr/testify
		go test -test.v github.com/sirkon/go-format
