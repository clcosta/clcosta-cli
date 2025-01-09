package b64

import (
	"fmt"

	"github.com/clcosta/clcosta-cli/pkg/utils"
)

// Configure sets up SSH configuration
func Configure(s string, decode bool) error {
	if decode {
		decoded, err := utils.DecodeB64(s)
		if err != nil {
			return err
		}
		fmt.Println(decoded)
		return nil
	}
	fmt.Println(utils.EncodeB64([]byte(s)))
	return nil
}
