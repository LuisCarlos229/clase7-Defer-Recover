package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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

//  GenerateID genera un nuevo id para la carga de un cliente
func GenerateID(clientList []Cliente) (string, error) {
	id := strconv.Itoa(len(clientList) + 1)
	if BuscarID(id, clientList) {
		return "", errors.New("el numero de legajo ya existe o es invalido")
	} else {
		return id, nil
	}
}

//verifica si ya no existe tal id que se desea cargar
func BuscarID(buscado string, clientList []Cliente) bool {
	for i, cliente_ := range clientList {
		if i%5 == 0 {
			if cliente_.legajo == buscado {
				return true
			}
		}
	}
	return false
}

//verifico la existencia de un cliente, a partir de su dni
func VerifyClient(buscado string, clientList []Cliente) bool {
	for i, cliente_ := range clientList {
		if (i-2)%5 == 0 {
			if cliente_.dni == buscado {
				return true
			}
		}
	}
	return false
}

//verifico que cada cliente_ de mi cliente sea distinta de un valor nulo
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
func MisClientes(clientList []Cliente) {
	for _, c := range clientList {
		fmt.Println(c.legajo)
		fmt.Println(c.nombre)
		fmt.Println(c.dni)
		fmt.Println(c.teléfono)
		fmt.Println(c.domicilio)
	}
}

func LoadClient(f *os.File, c Cliente, indxs int64) {
	Ntxt := ""

	Ntxt += "\n" + c.legajo + ","
	Ntxt += c.nombre + ","
	Ntxt += c.dni + ","
	Ntxt += c.teléfono + ","
	Ntxt += c.domicilio
	data2 := []byte(Ntxt)

	fmt.Println(f, Ntxt, c)

	_, err := f.WriteAt(data2, indxs)

	if err != nil {
		log.Panic(err)
	}
}
