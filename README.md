fast input data to key value db

* loadtest with 

- wrk -c 100 -d 60s -t 2 -s ./fast-post.lua http://localhost/register
