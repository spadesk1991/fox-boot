projectname = demo
port = 40102
serverip = 154.194.254.4
version = v0.0.1
devip = 192.168.1.222
releaseip = 39.108.208.60


default:
	@rm -rf debug
	@go build -o debug .
	@./debug

build:
	@$(info Build Linux)
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o server
	@rm -f $(projectname).img
	docker build -t $(projectname):$(version) .
	rm -f server
	docker save -o $(projectname).img $(projectname)
	docker rmi $(projectname):$(version)
dev:
	@make build
	scp $(projectname).img root@$(devip):/data/images/
	@rm -rf $(projectname).img
	ssh root@$(devip) "cd /data/images; docker rm -f $(projectname);docker rmi $(projectname):$(version);docker load --input $(projectname).img;docker run --restart=always -p $(port):8000 -e port=$(port) -e GIN_MODE=test -e $(port)=$(port) -v /data/$(projectname)_log:/log --privileged=true --name $(projectname) -itd $(projectname):$(version);rm -f $(projectname).img"
	ssh -o ServerAliveInterval=3 root@$(devip) "docker logs -f $(projectname)"
release:
	@make build
	scp $(projectname).img root@$(releaseip):/data/images/
	docker rmi $(projectname):$(version)
	rm -f server
	ssh root@$(releaseip) "cd /data/images;docker stop $(projectname) && docker rm $(projectname);docker rmi $(projectname):$(version);docker load --input $(projectname).img;docker run --restart=always -p $(port):8000 -e port=$(port) -e GIN_MODE=release -v /data/$(projectname)_log:/log --privileged=true --name $(projectname) -itd $(projectname):$(version);rm -f $(projectname).img"
	ssh -o ServerAliveInterval=3 root@$(releaseip) "docker logs -f $(projectname)"
