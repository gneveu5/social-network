package main

import (
	"os"

	f "socialnetwork/pkg/serv"
)

func main() {
	os.Setenv("SecretKey", "azezhfzuegehfdsyh") // -> debug only, environment variable needs to be setup in docker for production
	f.Runserver()
}
