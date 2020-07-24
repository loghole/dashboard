APP = dashboard

image:
	docker build -t $(APP):latest -f build/Dockerfile .

build-frontend:
	go-bindata -fs -o backend/bindata/bindata.go -prefix "./frontend/dist" -pkg "bindata" ./frontend/dist/...