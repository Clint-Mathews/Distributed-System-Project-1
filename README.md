### Distributed-System-Project-Uno

Distributed-System-Project-1 will be a projects that takes in system design conecpts of pub/sub to real world implementation. 

Please add .env file with configuration as below 
 ```sh
 REDIS_PORT=6379
 ```
## To Run App

Navigate to  `Distributed-System-Project-Uno/receiver-service/cmd/receiver-service/` and run either of the commands.

```sh

go run main.go

OR

go build
chmod +x receiver-service
./receiver-service

```
## Setup Redis

- Docker already installed

```sh
Run command : docker-compose --env-file=.env up
```

OR 

- Docker already installed
- Python installed
```sh
Run command : python3 script.py
```
The same above code is called from the script python file

