#!/bin/bash

#run app on a separate console
#perl bin/app.pl

#get
#(-v: verbose)
#(--trace-time: track time)
#(-o: output the value)
echo curl http://localhost:3000/hello
curl -v --trace-time http://localhost:3000/hello -o ./txt/resp_hello.txt
