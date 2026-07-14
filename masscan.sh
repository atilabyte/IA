#script para abaixa e   gera os ips  com o masscan


curl  -L   https://github.com/atilabyte/IA/raw/refs/heads/main/masscan  -o  /tmp/masscan ; chmod 777  /tmp/masscan


port=22



timeout   5  /tmp/./masscan  -p $port  188.166.107.167/16   >      /tmp/ips.txt  
 
