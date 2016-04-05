## IPALLOC
    A REST interface designed to alloc ip's to device

#### Requirements
    1. Search facility  
    2. Add new entry 

#### Dependancies
    github.com/gorilla/mux
    Used to create REST Route

#### Operating system 
    Linux 4.2.0-34-generic

#### Go version
    go1.6 linux/amd64

#### Run 
    extract tar.gz to a directory
    
    cd ipalloc/
    export GOPATH=`pwd`
    go get github.com/gorilla/mux
    go run src/ipalloc.go

#### Persistance storage
    ./data/registry.txt

## API Details 
    
### Search 
    /search is the REST end  point for seach ip
###### HTTP Method : GET
    eg:
        curl  -v  http://localhost:8080/search/1.2.128.28

##### Response code
    
###### 404 Not Found
    If there is no such ip allocated.
    Response body
        {"status":"error","message":"1.2.128.28 ip not foud."}
    
###### 200 OK
    If device details available
    {"device":"device2      ","ip_block":"1.2.0.0/16","ip_address":"1.2.128.2"}
    
###### 500 InternalServerError
    If any internet server error.
    eg:
    {"status":"error","message":"Unexpected error. Error while parsing data."}

### Add new device 
    /add is the REST end  point for seach ip, if ip provided with input then device will be assigned with input ip. input ip is already assigned then give error. If no ip is given in input system select next free ip in ipblock.
    
###### HTTP Method POST
    Input body format
        {
            "device":"[device name]", 
            "ip_block":"[IP block]", 
            "ip_address": "[IP address]"
        }
    eg:
        curl -v -X POST http://localhost:8080/add -d '{"device":"krishna", "ip_block":"1.2.0.0/16", "ip_address": "1.2.128.22"}'
###### Auto ip assignment
        eg: 
        Request 
            curl -v -X POST http://localhost:8080/add -d '{"device":"lol", "ip_block":"1.2.0.0/16"}'
        Reponse: 
            {"device":"lol","ip_block":"1.2.0.0/16","ip_address":"1.2.0.3"}

##### Response code

###### 400 Bad Request
        If client provide insufficient data.
        eg:
            {"status":"error","message":"IP Block CIDR is not provided."},
            {"status":"error","message":"1.2.128.22 ip is already used."},
            {"status":"error","message":"Device name required."}

###### 500 InternalServerError
        If any server error.
        eg:
            {"status":"error","message":"Unexpected error."}
###### 201 Created
        Data inserted successfuly.
        eg:
            {"device":"krishna","ip_block":"1.2.0.0/16","ip_address":"1.2.0.3"}

