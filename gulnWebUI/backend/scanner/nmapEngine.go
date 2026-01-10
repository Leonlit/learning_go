package scanner

import (
	"fmt"
	"gulnManagement/gulnWebUI/utils"
)

func startNewNmapScan(targetIP, templateUUID string) error {
	var isValidTargetAddr = utils.CheckAddressValid(targetIP)
	fmt.Print(isValidTargetAddr)
	return nil
}
