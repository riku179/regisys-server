HOST_ADDR := localhost:8080

all : goagen gormgen swagger-ui

goagen :
	env HOST_ADDR=$(HOST_ADDR) goagen -d github.com/riku179/regisys/design bootstrap

gormgen : goagen
	env HOST_ADDR=$(HOST_ADDR) goagen --design=github.com/riku179/regisys/design gen --pkg-path=github.com/goadesign/gorma

js :
	env HOST_ADDR=$(HOST_ADDR) goagen -d github.com/riku179/regisys/design js

swagger-ui :
	cp -r dist swagger/
