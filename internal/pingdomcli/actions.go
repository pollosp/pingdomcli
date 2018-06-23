package actions

import (
	"fmt"
	"strconv"

	"github.com/russellcardullo/go-pingdom/pingdom"
)

func checkIfMonitorExist(c pingdom.Client, checkHostName string) bool {
	checks, _ := c.Checks.List()

	for index := range checks {
		if checkHostName == checks[index].Hostname {
			return true
		}
	}
	return false
}

//CreateHTTPCheck creates pingdom monitor for the given parameters
func CreateHTTPCheck(c pingdom.Client, checkName string, checkHostName string, userID string) bool {

	if !checkIfMonitorExist(c, checkHostName) {
		id, _ := strconv.Atoi(userID)
		UserIds := []int{id}
		newCheck := pingdom.HttpCheck{Name: checkName, Hostname: checkHostName, Resolution: 5, SendNotificationWhenDown: 2, UserIds: UserIds}
		check, _ := c.Checks.Create(&newCheck)
		fmt.Println("Created check:", check)
		return true
	}
	fmt.Println("Already exist check:", checkHostName)
	return false
}
