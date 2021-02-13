package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/tarm/serial"
)

type ArmState struct {
	BaseAngle      int
	ShoulderAngle  int
	ElbowAngle     int
	WristVertAngle int
	WristRotAngle  int
	GripperAngle   int
}

var state ArmState = ArmState{
	BaseAngle:      90,
	ShoulderAngle:  90,
	ElbowAngle:     90,
	WristVertAngle: 90,
	WristRotAngle:  90,
	GripperAngle:   73,
}

const (
	BaseOption      = "Base"
	ShoulderOption  = "Shoulder"
	ElbowOption     = "Elbow"
	WristVertOption = "Wrist Verticle"
	WristRotOption  = "Wrist Rotation"
	GripperOption   = "Gripper"

	DefaultBaud = 9600
)

var validateAngle = func(input string) error {

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("Invalid number")
	}

	maxAngle := 180.0
	if input == WristVertOption {
		maxAngle = 72.0
	}

	if value < 0 || value > maxAngle {
		return errors.New("Invalid angle")
	}

	return nil
}

type ArmController struct {
	serialConn *serial.Port
}

func main() {

	channel := os.Getenv("ARM_SERIAL_CHANNEL")

	serialConfig := &serial.Config{
		Name: channel,
		Baud: DefaultBaud,
	}
	serialConn, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := serialConn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	armController := ArmController{
		serialConn: serialConn,
	}

	options := []string{BaseOption, ShoulderOption, ElbowOption, WristVertOption, WristRotOption, GripperOption}
	printState(state)

	isRunning := true
	for isRunning {

		prompt := promptui.Select{
			Label: "Choose Joint",
			Items: options,
			Size:  len(options),
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case BaseOption:
			state.BaseAngle = fetchAngle()
			armController.writeBaseAngleToArm(state.BaseAngle)
		case ShoulderOption:
			state.ShoulderAngle = fetchAngle()
			armController.writeShoulderAngleToArm(state.ShoulderAngle)
		case ElbowOption:
			state.ElbowAngle = fetchAngle()
			armController.writeElbowAngleToArm(state.ElbowAngle)
		case WristVertOption:
			state.WristVertAngle = fetchAngle()
			armController.writeWristVertAngleToArm(state.WristVertAngle)
		case WristRotOption:
			state.WristRotAngle = fetchAngle()
			armController.writeWristRotAngleToArm(state.WristRotAngle)
		case GripperOption:
			state.GripperAngle = fetchAngle()
			armController.wristGripperAngleToArm(state.GripperAngle)
		}

		printState(state)
	}

}

func printState(state ArmState) {
	fmt.Printf(
		`
		Base: %d (0 - 180)
		Shoulder: %d (15 - 165)
		Elbow: %d (0 - 180)
		Wrist Vert: %d (0 - 180)
		Wrist Rot: %d (0 - 180)
		Gripper %d (10 - 73)
		`,
		state.BaseAngle,
		state.ShoulderAngle,
		state.ElbowAngle,
		state.WristVertAngle,
		state.WristRotAngle,
		state.GripperAngle)
}

func fetchAngle() int {
	prompt := promptui.Prompt{
		Label:    "Choose Angle",
		Validate: validateAngle,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	valueInt, _ := strconv.Atoi(result)

	return valueInt
}

// Blocks
func (ac *ArmController) writeBaseAngleToArm(angle int) {
	ac.writeCmd("B", angle)
}

func (ac *ArmController) writeShoulderAngleToArm(angle int) {
	ac.writeCmd("S", angle)
}

func (ac *ArmController) writeElbowAngleToArm(angle int) {
	ac.writeCmd("E", angle)
}

func (ac *ArmController) writeWristVertAngleToArm(angle int) {
	ac.writeCmd("V", angle)
}

func (ac *ArmController) writeWristRotAngleToArm(angle int) {
	ac.writeCmd("R", angle)
}

func (ac *ArmController) wristGripperAngleToArm(angle int) {
	ac.writeCmd("G", angle)
}

func (ac *ArmController) writeCmd(cmd string, angle int) {
	_, err := ac.serialConn.Write([]byte(cmd + strconv.Itoa(angle)))
	if err != nil {
		log.Fatal(err)
	}

	readBuf := make([]byte, 10)
	_, err = ac.serialConn.Read(readBuf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(readBuf))
}
