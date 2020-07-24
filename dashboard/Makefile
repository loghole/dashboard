APP = dashboard

default-image:
	docker build -t $(APP):latest -f docker/default/Dockerfile .

build-frontend:
	go-bindata -fs -o backend/bindata/bindata.go -prefix "./frontend/dist" -pkg "bindata" ./frontend/dist/...