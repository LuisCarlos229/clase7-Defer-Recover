package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Cliente struct {
	legajo    string
	nombre    string
	dni       string
	teléfono  string
	domicilio string
}

func main() {
	centinela := false

	var NewClient func() Cliente

	defer func() {
		if centinela == true {
			fmt.Println("ejecución finalizada")
		} else {
			fmt.Println("Se detectaron varios errores en tiempos de ejecución")
		}
	}()

	file, err := os.OpenFile("customers.txt", os.O_RDWR, 0)
	defer file.Close()
	if err != nil {
		log.Panic("el archivo indicado no fue encontrado o está dañado")
	}
	log.Println(35)
	clientes, indexs, err := readCliente()
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	log.Println(40)
	NewClient = func() Cliente {
		N_Client := Cliente{}
		return N_Client
	}

	client := NewClient()

	id, err := GenerateID(clientes)
	if err != nil {
		log.Panic(err)
	}
	log.Println(52)

	client.legajo = id
	client.nombre = "juan"
	client.dni = "555999888"
	client.teléfono = "37550003"
	client.domicilio = "AV. Los Cedros 5"

	registrado := VerifyClient(client.dni, clientes)
	if err := VerifyNil(&client); err != nil {
		log.Println(err)
		fmt.Println("uno o varios datos del cliente estan sin registrar")
	}

	if registrado == false {
		LoadClient(file, client, indexs)
		fmt.Println(*file)
		fmt.Println(readCliente())
	} else {
		fmt.Println("el cliente ya existe")
	}

	dat, _ := ioutil.ReadFile("customers.txt")
	MisClientes(dat)

	log.Println(79)

	centinela = true
}
