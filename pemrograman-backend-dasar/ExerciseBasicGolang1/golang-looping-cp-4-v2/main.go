package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	//flagging = false //true
	//admin@yahoo.co.id
	//0.................last

	//slicing
	//[admin] [yahoo.co.id]
	emailSplitted := strings.Split(email, "@")

	emailInfo := strings.Split(emailSplitted[1], ".")

	var TLD string

	if len(emailInfo) > 2 {
		TLD = strings.Join(emailInfo[1:], ".")
	} else {
		TLD = emailInfo[1]
	}

	// fmt.Println(TLD)
	// TODO: answer here
	return "Domain: " + emailInfo[0] + " dan TLD: " + TLD // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
