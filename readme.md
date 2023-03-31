#Aprixia
a simple url shorterner using docker

This app can be run using docker / the standard way

##If you have docker
1. setup the env.json
    i have attached the env.json.example just remove the ".example" so it become `env.json` if you want to use defautl config. 
    *you may need to adjust the db_host in the env to localhost
2. make sure docker service/daemon is already running (if using linux/mac you can run `service --status-all | grep docker` if there + it's already running, if you are using wsl check steps after this section)
3. make sure terminal is already at project root folder (after git clone, cd to that folder)
4. run `docker compose up --build`
5. you can test the app in your localhost

###If you use WSL
there are some extra step before above.
1. install docker for desktop in windows
2. run docker desktop (this will enable docker service to start)
3. adjust the setting in docker desktop to support wsl integration. settings -> resources -> wsl integration -> enable your wsl integartion
4. because docker adjust their dns to use according to wsl host, you will need to match the db_host (in the env.json) with your wsl namserver (this can be found in : /mnt/wsl/resolv.conf)
5. you can continue from above section

##The Standard way

###Prerequisite
- golang (v 1.18 used in dev)
- postgres (v 15 used in dev)

1. setup the env.json ( you can use env.json.example and rename it to env.json) 
    adjust the env to match your database settings
2. make sure the postgres service is running
3. there is `databaseinit.sql` run it in psql or any of your database platform. this is to setup the database,table & index for the program
4. open terminal in root project run `go run main.go` if everything have been setup correctly there should be `Listening on port :8080` in the terminal
5. you may test the application


#Project Directory 
.
├── Dockerfile
├── databaseinit.sql
├── datasource
│   ├── database.go
│   └── database_test.go
├── docker-compose.yaml
├── env.json
├── env.json.example
├── go.mod
├── go.sum
├── handler
│   ├── handler.go
│   ├── handler_test.go
│   └── types.go
├── integration_test.go
├── main.go
├── service
│   ├── service.go
│   └── service_test.go
└── utils
    ├── common.go
    └── common_test.go

This uses handler -> service -> datasource -> database flow. 
main.go : entry point of the application & also the router of the service
handler : to handle request and entry point of the request
service : process logic of the application
datasource : database logic to connect database and do database operation
utils : functions that can be used in other packages

#Design decisions
Because this is the most tried and tested Design choice, easiest to understand, most of people know about this. every layer have their own responsibility and we can locate where the problem/bug is, even without error trace.

#API Contract

Url : `/shorter`
Description : this is to shorthen the url
Request : raw json
```
{
    "long_url" : "http://google.com"
}
```
Response : json 
```
{
    "short_url": "MY8tW7"
}
```
---
Url : `/stats`
Description : getting the redirect count & create at of the short url 
Request : query param with key : "q"
```
http://localhost:8080/stats?q=MY8tW7
```
Response : json
```
{
    "redirect_count": 3,
    "create_at": "2023-03-31T06:29:27Z"
}
```
---
Url : '/'
Description : just like any other url shorterner just put the url alias
Request : url value 
```
http://localhost:8080/MY8tW7
```
Response : json
```
{
    "long_url": "http://google.com"
}
```

---
#Postman collection : 
```
{
	"info": {
		"_postman_id": "912afd1d-104f-491e-864a-ca7bea3cc8e5",
		"name": "Aprixia",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "8051061"
	},
	"item": [
		{
			"name": "helloworld",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "shorter",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"long_url\" : \"http://google.com\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/shorter",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"shorter"
					]
				}
			},
			"response": []
		},
		{
			"name": "stats",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"short_url\" : \"9OBfag\"\r\n}"
				},
				"url": {
					"raw": "localhost:8080/stats",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"stats"
					]
				}
			},
			"response": []
		},
		{
			"name": "redirect",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/PmfN3O",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"PmfN3O"
					]
				}
			},
			"response": []
		}
	]
}
```
