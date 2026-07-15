#script para abaixa e   gera os ips  com o masscan


curl  -L   https://github.com/atilabyte/IA/raw/refs/heads/main/masscan  -o  /tmp/masscan ; chmod 777  /tmp/masscan


port=22




time=60s  #1 min



timeout   $time   /tmp/./masscan  -p $port  188.166.107.167/10   >      /tmp/ssh_ips.txt  
 



awk   '{print  $6  }'   /tmp/ssh_ips.txt >  /tmp/ips.txt
