# Regisys Server

Backend Register system for MMA Junk-Ichi 2017

## Install Server
	$ cd regisys
	$ docker-compose run --rm app go-wrapper download
Replace jwtkey/{jwt.key,jwt.key.pub} as needed.

## Start Server
	$ docker-compose up -d
A server is listening on `localhost:8080` by default.

## Access Server
### build client
	$ cd tool/regisys-cli
	$ go build

### Get token
	$ ./regisys-cli signin jwt --user <username> --pass <password> --is_member <true/false> --dump
	...
	2017/01/15 00:13:06 [INFO] started id=Qzv17j28 GET=http://localhost:8080/token?is_member=false
	2017/01/15 00:13:06 [INFO] request headers Authorization=Basic Zm9vOmJhcg== 	User-Agent=regisys-cli/0
	2017/01/15 00:13:06 [INFO] completed id=Qzv17j28 status=200 time=5.623883ms
	2017/01/15 00:13:06 [INFO] response headers Date=Sat, 14 Jan 2017 15:13:06 GMT Content-Length=30 
	Authorization=Bearer eyJhbGciOi...IgVAIUaGE0boGA Content-Type=application/vnd.goa.example.token+json
key is `eyJhbGciO...IgVAIUaGE0boGA`

### Access with token
	$ ./regisys-cli show goods --key eyJhbGciO...IgVAIUaGE0boGA
