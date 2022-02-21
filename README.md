fast input data to key value db

* test api 

curl -X POST -d 'data={"name":"test"}' http://localhost/register

curl http://localhost/list

* loadtest with 

- wrk -c 100 -d 60s -t 2 -s ./fast-post.lua http://localhost/register
