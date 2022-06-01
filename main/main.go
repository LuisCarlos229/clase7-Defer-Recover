/*package main

import (
	//s "Strings"
	"fmt"
	"log"
	"os"
)

func BuscarID(str string) bool {
	var buscado string
	fmt.Println(str)

	fmt.Println("\n\n\ningrese el legajo buscado: ")
	fmt.Scan(&buscado)
	buscado = string(buscado)
	/*
		fmt.Println(s.Index(str, buscado) + 10)
		indexs1 := s.Index(str, buscado) + 10
		indexs2 := indexs1
		for i := indexs1; i < (indexs1 + 20); i++ {
			if str[indexs2:indexs2+2] == "\n" {
				break
			}
			indexs2 += 1
		}
		fmt.Println("nombre: ", str[indexs1:indexs2])

	line, c := 0, 0
	for i := 0; i < len(str); i++ { //recorro el txt
		if (line % 6) == 0 { //examino solo las lineas que son multiplos de 6, ya que son las que lineienen el ID del usuario
			fmt.Printf("line: %v\n", line)
			c = i
			for str[c:c+1] != "\n" { // este for recorre la linea desde el comienso a fin
				if buscado == str[c:c+1] {
					fmt.Printf("str: %v  buscado: %v\n", str[c:c+1], buscado)
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
}

func main() {
	data, err := os.ReadFile("customers.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)

	b := BuscarID(dataString)
	fmt.Println("el resultado de la busqueda es: ", b)
}
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Client struct {
	id     string
	name   string
	dni    string
	phone  string
	addres string
}

func readDb() ([]Client, error) {
	// Leer archivo de texto...
	// data := "0001,Lucas Baltazar,44333222,37550001,AV. Picapau 50**0002,Lautaro vega,44333555,37550002,AV. Castell 10"

	data_bytes, err := os.ReadFile("test.txt")
	if err != nil {
		return nil, err
	}
	data := fmt.Sprintf("%s", data_bytes)

	clients_str := strings.Split(data, "**") //["0001,Lucas Baltazar,44333222,37550001,AV. Picapau 50","Lautaro vega,44333555,37550002,AV. Castell 10"]
	var clients []Client
	for _, client_str := range clients_str {
		client_arr := strings.Split(client_str, ",") // ["0001","Lucas Baltazar","44333222","37550001","AV. Picapau 50"]
		client := Client{
			id:     client_arr[0],
			name:   client_arr[1],
			dni:    client_arr[2],
			phone:  client_arr[3],
			addres: client_arr[4],
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func main() {
	fmt.Println("Hello World")
	clients, err := readDb()
	if err != nil {
		log.Panic("error: el archivo no fue encontrado o está dañado")
	}
	fmt.Println(clients)
}
