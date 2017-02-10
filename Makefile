HOST_ADDR := localhost:8080

all : goagen gormgen swagger/dist

goagen:
	go get github.com/goadesign/goa/goagen

goagen :
	env HOST_ADDR=$(HOST_ADDR) goagen -d github.com/riku179/regisys-server/design bootstrap

gormgen : goagen
	env HOST_ADDR=$(HOST_ADDR) goagen --design=github.com/riku179/regisys-server/design gen --pkg-path=github.com/goadesign/gorma

swagger/dist : swagger-ui
	cp -r swagger-ui/dist swagger/dist

js :
	env HOST_ADDR=$(HOST_ADDR) goagen -d github.com/riku179/regisys/design js

swagger-ui :
	@git clone https://github.com/swagger-api/swagger-ui.git
