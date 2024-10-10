# web-service-gin

There is a config.yml file that contains directory where the participants will be stored/read from. The directory contains all participants as .json file e.g. P1.json, P2.json ...

GET 

http://localhost:8080/participants

It reads all .json files present inside the configured directory (see config.yml)

POST

http://localhost:8080/participants

It creates a .json file e.g. P1.json inside the configured directory (see config.yml)

GET

http://localhost:8080/participants/P1

It reads P1.json file from the configured directory and responds with the content 