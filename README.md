# Sparrow

------

 ***麻雀虽小，五脏俱全*  --  *The sparrow may be small but it has all the vital organs.***

## Project Description

sparrow is a simple microservice demo (with grpc protocol)

------

## Project Structure

````
```
sparrow                		root directory
├── api                 		 grpc protobuf directory
├── build                    build directory
│   ├── ci                      ci/cd config 
│   ├── external_env            external config (prometheus only for now)
├── cmd                      command 
├── config                   project config
│   ├── config.go            		read config
│   └── dev.yaml             		dev enviroment config
│   └── prod.yaml             	prod enviroment config
├── deploy                   deploy cmd
├── init                     init directory   
├── internal								 business logic
│   ├── app                  		business application
│   ├── pkg                  		pkg (app package will import)
				├── bitmap              		bit map 
        ├── biz_redis               business redis client
        ├── entity              		global define
        ├── fuse              			api fuse 
        ├── middleware              middleware 
        ├── monitor              		prometheus monitor & pprof 
        ├── serror              		error package 
        ├── service_mng             service register & find & loadbalance 
        ├── slog              			log package 
        ├── trace              			jaeger trace  
        ├── util              			util tools 
├── log                      log    
├── scripts                  some scripts
├── README.md                         
└── go.mod                     
```
````

------

## Todo

- [ ] Http 
- [ ] Other...
