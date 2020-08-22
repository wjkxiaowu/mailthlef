all: mail generator thlef

mail:
	CGO_ENABLED=0 go build -a -ldflags '-s -w -extldflags "-static"' -o release/mail src/mail.go

generator:
	go build -o release/generator src/generator.go
	
thlef:
	cp src/thlef.tmpl release/thlef.go
	./release/generator >> ./release/thlef.go
	go build -o ./release/thlef ./release/thlef.go 

clean:
	rm -rf release
    
