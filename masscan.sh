#script para abaixa e   gera os ips  com o masscan


curl  -L   https://github.com/atilabyte/IA/raw/refs/heads/main/masscan  -o  /tmp/masscan ; chmod 777  /tmp/masscan


port=22




time=300s  #5 min





timeout   $time   /tmp/./masscan  -p $port  187.101.1.35/10      >      /tmp/ssh_ips.txt  
 

timeout   $time   /tmp/./masscan  -p $port   45.236.241.52/10        >>      /tmp/ssh_ips.txt 

                                                                     


awk   '{print  $6  }'   /tmp/ssh_ips.txt  >  /tmp/ips.txt
