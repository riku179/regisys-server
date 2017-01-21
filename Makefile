HOST_ADDR := localhost:8080

all : goagen gormgen

goagen :
	env HOST_ADDR=$(HOST_ADDR) goagen -d github.com/riku179/regisys/design bootstrap

gormgen : goagen
	env HOST_ADDR=$(HOST_ADDR) goagen --design=github.com/riku179/regisys/design gen --pkg-path=github.com/goadesign/gorma
