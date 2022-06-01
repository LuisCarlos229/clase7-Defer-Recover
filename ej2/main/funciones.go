package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

//lee el archivo customers.txt y lo devuelve como un array del tipo Cliente
//para poder manejar de mejor manera sus datos
func readCliente() ([]Cliente, int64, error) {
	// Leer archivo de texto...
	// data := "0001,Lucas Baltazar,44333222,37550001,AV. Picapau 50**0002,Lautaro vega,44333555,37550002,AV. Castell 10"

	data_bytes, err := os.ReadFile("customers.txt")
	if err != nil {
		return nil, 0, err
	}

	indx := int64(len(data_bytes))
	data := fmt.Sprintf("%s", data_bytes)

	clients_str := strings.Split(data, "**")
	var clients []Cliente
	for _, client_str := range clients_str {
		client_arr := strings.Split(client_str, ",")
		client := Cliente{
			legajo:    client_arr[0],
			nombre:    client_arr[1],
			dni:       client_arr[2],
			teléfono:  client_arr[3],
			domicilio: client_arr[4],
		}
		clients = append(clients, client)
	}
	return clients, indx, nil
}

//genera un nuevo id para la carga de un cliente
func ID(str []Cliente) (string, error) {
	id := string(len(str)/5 + 1)
	if BuscarID(id, str) {
		return "", errors.New("el numero de legajo ya existe o es invalido")
	} else {
		return id, nil
	}
}

//verifica si ya no existe tal id que se desea cargar
func BuscarID(buscado string, str []Cliente) bool {
	for i, componente := range str {
		if i%5 == 0 {
			if componente.legajo == buscado {
				return true
			}
		}
	}
	return false

	/*var buscado string

	fmt.Println("\ningrese el legajo buscado: ")
	fmt.Scan(&buscado)
	buscado = string(buscado)

	line, c := 0, 0
	for i := 0; i < len(str); i++ { //recorro el txt
		if (line % 6) == 0 { //examino solo las lineas que son multiplos de 6, ya que son las que lineienen el ID del usuario
			c = i
			for str[c:c+1] != "\n" { // este for recorre la linea desde el comienso a fin
				if buscado == str[c:c+1] {
					return true
				}
				c += 1
			}
		}
		if str[i:i+1] == "\n" { //busco cada salto de linea para linear a estas
			line += 1
		}
	}
	return false
	*/
}

//verifico la existencia de un cliente, a partir de su dni
func VerifyClient(buscado string, str []Cliente) bool {
	for i, componente := range str {
		if (i-2)%5 == 0 {
			if componente.dni == buscado {
				return true
			}
		}
	}
	return false
}

//verifico que cada componente de mi cliente sea distinta de un valor nulo
func VerifyNil(c *Cliente) error {
	if c.legajo == "" {
		return errors.New("error: legajo sin registrar")
	}
	if c.nombre == "" {
		return errors.New("error: nombre sin registrar")
	}
	if c.domicilio == "" {
		return errors.New("error: domicilio sin registrar")
	}
	if c.dni == "" {
		return errors.New("error: DNI sin registrar")
	}
	if c.teléfono == "" {
		return errors.New("error: telefono sin registrar")
	}
	return nil
}

//muestra por pantalla el registro de clientes
func MisClientes(str []Cliente) {
	for i, c := range str {
		if (i+1)%5 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print(c)
		}
	}
}

func LoadClient(f *os.File, c Cliente, indxs int64) {
	Ntxt := ""

	Ntxt += c.legajo + ","
	Ntxt += c.nombre + ","
	Ntxt += c.dni + ","
	Ntxt += c.teléfono + ","
	Ntxt += c.domicilio + "\n"
	data2 := []byte(Ntxt)

	fmt.Println(f, Ntxt, c)

	_, err := f.WriteAt(data2, indxs)

	if err != nil {
		log.Panic(err)
	}
}
