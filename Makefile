build-rpi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o arm-cli main.go 

deploy-rpi: build
	rsync -a arm-cli pi@$(ARM_PI_IP):/home/pi/arm-cli

build:
	go build -o arm-cli main.go

connect: 
	ssh pi@$(ARM_PI_IP)