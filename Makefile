all: update

update:
	curl https://www.alexa.com/topsites > Global.html
	curl https://www.alexa.com/topsites/countries/BR > BR.html
	go run main.go
	rm -Rf Global.html BR.html