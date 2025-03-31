
https://bugs.freebsd.org/bugzilla/show_bug.cgi?id=283747

- ```git clone https://github.com/ekilimchuk/junk.git```
- ```cd ./junk/fork-test```
- ```sudo ./start.sh```
- ```sudo tail -f start.log```
- Run after a few days: ```ps aux | grep -E "(go ru[n]|/tmp/go-buil[d])" | awk '{print $2}' | xargs kill```

