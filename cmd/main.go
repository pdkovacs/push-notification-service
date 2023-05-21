package cmd

import (
	"fmt"

	pnservice "pns"
)

func main() {
	pnService := &pnservice.PnService{}
	fmt.Printf("%#v", pnService)
}
