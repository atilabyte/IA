

dinamic_bin=1
static_bin=0

export  CGO_ENABLED=$static_bin



go build  atila.go    run_vkzmn.go  down_vkzmn.go  tgram.go  magic.go  down_raw.go 


 if [ $? -eq  0 ] ;  then 

 
 cd  *.dir ; ./upx  ../atila  
  
 

 
sucess=1


fi 


if [ $sucess -eq 1 ]  ; then 



cd ../  ;  bash git.sh || sh -c git.sh 
 


fi

