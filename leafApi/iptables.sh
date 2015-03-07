#! /bin/sh

help()
{
	echo " use: $0 drop | accept | reject "
	exit
}

[ ! $# -eq 1 ] && help

case $1 in  
	"accept" )
		/sbin/iptables -D INPUT 1  > /dev/null 2>&1
		;;
    "drop" )
		/sbin/iptables -I INPUT -s 10.44.3.23 -p tcp --dport 8080  -j DROP > /dev/null 2>&1
		;;
	"reject" )
		/sbin/iptables -I INPUT -s 10.44.3.23 -p tcp --dport 8080 -j REJECT > /dev/null 2>&1
		;;
	* )
	help;;
esac

