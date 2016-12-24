all: sassgen templegen app-assets appjs sv run
	echo all done


build: sassgen templegen app-assets appjs sv 

get: 
	go get -u honnef.co/go/simple/cmd/gosimple
	go get -u github.com/gopherjs/gopherjs
	go get -u github.com/gopherjs/websocket
	go get -u github.com/go-humble/temple
	go get -u github.com/go-humble/form
	go get -u github.com/go-humble/router
	go get -u github.com/go-humble/locstor
	go get -u github.com/steveoc64/formulate
	go get -u github.com/steveoc64/godev/echocors
	go get -u github.com/steveoc64/godev/sms
	go get -u github.com/steveoc64/godev/smt
	go get -u github.com/steveoc64/godev/db
	go get -u github.com/steveoc64/godev/config
	go get -u honnef.co/go/simple/cmd/gosimple
	go get -u github.com/rs/cors
	go get -u gopkg.in/gomail.v2
	go get -u github.com/labstack/echo
	go get -u github.com/labstack/echo/middleware
	go get -u github.com/lib/pq
	go get -u gopkg.in/mgutz/dat.v1/sqlx-runner
	go get -u github.com/nfnt/resize
	go get -u gopkg.in/gomail.v2
	go get -u github.com/logpacker/PayPal-Go-SDK
	mkdir -p scripts
	mkdir -p backup

help: 
	# sassgen    - make SASS files
	# templegen  - make Templates
	# app-assets - make Asset copy to dist	
	# appjs      - make Frontend app
	# sv         - make Server
	# run        - run  Server

clean:	
	# Delete existing build
	@mplayer -quiet ../audio/trash-empty.oga 2> /dev/null > /dev/null &
	rm -rf dist

sassgen: dist/public/css/ministry.css

dist/public/css/ministry.css: scss/*
	@mplayer -quiet ../audio/attention.oga 2> /dev/null > /dev/null
	@mkdir -p dist/public/css
	cd scss && node-sass --output-style compressed app.sass ../dist/public/css/ministry.css
	cd scss && node-sass app.sass ../dist/public/css/ministry.debug.css

templegen: app/template.go 

app/template.go: templates/*.tmpl 	
	@mplayer -quiet ../audio/attention.oga 2> /dev/null > /dev/null
	temple build templates app/template.go --package main

app-assets: dist/assets.log dist/config.json

dist/config.json: server/config.json
	cp server/config.json dist	

cert:
	mkdir -p cert
	openssl genrsa -out cert/ministry.key 2048
	openssl req -new -x509 -key cert/ministry.key -out cert/ministry.pem -days 3650

dist/assets.log: assets/*.html assets/img/*  assets/fonts/* assets/*.webmanifest  assets/css/*
	@mplayer -quiet ../audio/attention.oga 2> /dev/null > /dev/null
	@mkdir -p dist/public/css dist/public/font dist/public/js
	cp assets/*.html dist/public
	cp assets/*.webmanifest dist/public
	cp -R assets/img dist/public
	cp -R assets/fonts dist/public
	cp -R assets/css dist/public
	@date > dist/assets.log

appjs: dist/public/ministry.js

dist/public/ministry.js: app/*.go shared/*.go 
	@mplayer -quiet ../audio/frontend-compile.ogg 2> /dev/null > /dev/null &
	@mkdir -p dist/public/js
	GOOS=linux gopherjs build app/*.go -o dist/public/ministry.js -m
	@ls -l dist/public/ministry.js

remake: 
	@mplayer -quiet ../audio/server-compile.oga 2> /dev/null > /dev/null &
	rm -f dist/ministry-server
	@gosimple server
	go build -o dist/ministry-website server/*.go
	@ls -l dist/ministry-website

sv: dist/ministry-website 

dist/ministry-website: server/*.go shared/*.go
	@mplayer -quiet ../audio/server-compile.oga 2> /dev/null > /dev/null &
	cd server && gosimple
	go build -o dist/ministry-website server/*.go
	@ls -l dist/ministry-website
	cp cert/ministry.key cert/ministry.pem dist

run: 
	./terminate
	@mplayer -quiet ../audio/running.oga 2> /dev/null > /dev/null &
	@cd dist && ./ministry-website

install: sv
	./terminate
	cp -Rv dist/* ~/ministry/current
	cd ~/ministry/current && nohup ./ministry-website &

tail:
	tail -f -n 200 ~/logs/ministry/* ~/ministry/current/nohup.out

data:
	pg_dump ministry > database/ministry.sql
	scp -P 446 database/ministry.sql freebsd@bsd:/home/freebsd/ministry.sql

loadtempdata:
	./terminate
	dropdb ministry
	createdb ministry
	psql ministry < ~/ministry.sql
	cd ~/ministry/current && nohup ./ministry-website &

loaddata:
	./terminate
	dropdb ministry
	createdb ministry
	psql ministry < database/ministry.sql

