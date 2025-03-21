delay=0
while true;
do
	delay=$((delay + 1))
	exp_delay=$(($delay * 2))
	echo "Start delay"
	echo "delay = $exp_delay * 60"
	for i in `seq $exp_delay`
	do
		echo "sleep 60"
		sleep 60
	done
	pid=$(pgrep telegraf)
#	echo $pid
	if [ "$pid " != " " ]
	then
		echo "kill $pid"
		sudo kill $pid
	fi
done
