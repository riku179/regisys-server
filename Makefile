HOST_ADDR := localhost:8080
# 開発環境構築・dockerの依存関係解決
dep :
	go get -u github.com/golang/dep/cmd/dep

develop : dep
	dep ensure

build : generate
	docker-compose run --rm app go-wrapper download

generate : goagen gormgen swagger/dist

goagen :
	env HOST_ADDR=$(HOST_ADDR) goagen -d github.com/riku179/regisys-server/design bootstrap

gormgen : goagen
	env HOST_ADDR=$(HOST_ADDR) goagen --design=github.com/riku179/regisys-server/design gen --pkg-path=github.com/goadesign/gorma

swagger/dist : swagger-ui
	cp -r swagger-ui/dist swagger/dist

swagger-ui :
	@git clone https://github.com/swagger-api/swagger-ui.git
