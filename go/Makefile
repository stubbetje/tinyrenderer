
output/image.png: output ../main
	time ../main
	cp output/image.png output/image.$(shell date +%Y%m%d_%H%I%S).png

output:
	mkdir output

../main: $(wildcard *.go) $(wildcard */*.go)
	go build -o ../main . 
