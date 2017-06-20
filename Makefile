all:
	go build -o "tic-tac-toe" main.go
	@echo "--COMPILED !"

clean:
	rm "tic-tac-toe"
	@echo "--CLEANED !"
