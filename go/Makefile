
output/image.png: output main
	./main
	cp output/image.png output/image.$(shell date +%Y%m%d_%H%I%S).png

output:
	mkdir output

main: $(wildcard *.go)
	go build -o main . 
