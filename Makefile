.PHONY: build_back
build_back:
	go build -o ./bin/app.exe ./app

.PHONY: build_front 
build_front:
	cd ./osau_biblio_front && npm run build && cd ..

.PHONY: build
build:
	make -j 2 build_back build_front

.PHONY: run_back
run_back:
	cd ./bin && ./app.exe

.PHONY: run_front
run_front:
	cd ./osau_biblio_front && npm run start

.PHONY: dev
dev: 
	make -j 2 run_back run_front