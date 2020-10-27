package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
)

func main() {

	f, err := os.OpenFile("info.txt", os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// get groups
	adminGrp, err := user.LookupGroup("admin")
	if err != nil {
		log.Fatal("Error:", err)
	}

	admGid, _ := strconv.ParseInt(adminGrp.Gid, 10, 64)
	fmt.Println("Group Id =>", admGid)

	// get users
	ukayUser, err := user.Lookup("ukay.itachi")
	if err != nil {
		log.Fatal("Error:", err)
	}

	ukayID, _ := strconv.ParseInt(ukayUser.Uid, 10, 64)
	fmt.Println("User Id =>", ukayID)

	// Change mode and owner
	f.Chmod(0755)
	f.Chown(int(ukayID), int(admGid)) // to ukay:admin_group

}
