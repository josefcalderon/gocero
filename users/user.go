package users

import (
	"fmt"
	"josefcalderon/gocero/models"
	"time"
)

func AltaUsuario() {
	u := new(models.User)
	u.AddUser(10, "Edwin", time.Now(), true)
	fmt.Println(u)
}
