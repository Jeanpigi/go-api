package handlererror

import (
	"fmt"
	"log"
)

func CheckError (err error, mensaje string) {
	if err != nil {
		log.Fatal("Se ha presentado el siguiente error: ", err)
	}
	fmt.Println(mensaje)
}