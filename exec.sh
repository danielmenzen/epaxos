#/bin/bash

#echo $BASHPID

bin/master  2>>master.log &

masterpid=$!

#disown %%

bin/server -port 7070 2>>server1.log &

server1pid=$!

bin/server -port 7071 2>>server2.log &

server2pid=$!

bin/server -port 7072 2>>server3.log &

server3pid=$!

bin/client

kill $masterpid $server1pid $server2pid $server3pid