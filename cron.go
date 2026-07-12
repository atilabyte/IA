package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//funcao para bota o atila  no crontab

var mybin string = "atila"

func main() {

	cron_read := exec.Command("crontab", "-l") //isso pode da errado   com o seguinte aviso ->  no crontab for $USER
	//isso sera aprimorado mais tarde

	out, err_combinedoutput := cron_read.CombinedOutput()

	if err_combinedoutput != nil {

		//fmt.Println("erro em combinedoutput")

		fmt.Println(err_combinedoutput)

		return

	}

	out_str := string(out)

	out2 := strings.Contains(out_str, mybin)

	if out2 == false {

		//fmt.Println("atila nao esta no crontab")

		goto label

	}

	if out2 == true {

		//fmt.Println("atila esta no crontab")

		return

	}

label:

	pwd, err := os.Getwd()

	fmt.Println(err)

	cmd := fmt.Sprintf("crontab -l   ;  echo '* * * * *  pgrep   atila   ||  %s/%s' | crontab -", pwd, mybin)

	cron_exec := exec.Command("sh", "-c", cmd)

	cron_exec.Run()

}
