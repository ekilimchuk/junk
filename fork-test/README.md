git clone
cd ./junk/fork-test
sudo ./start.sh
sudo tail -f start.log

ps aux | grep -E "(go ru[n]|/tmp/go-buil[d])" | awk '{print $2}' | xargs kill

