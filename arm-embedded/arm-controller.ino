/*
  arm-controller.ino

  controller software for interfacing with the Braccio arm via
   a keyboard CLI connnected over serial and running on Linux

 Created on 18 Nov 2015
 by Joshua VonFricken
 */

#include <Braccio.h>
#include <Servo.h>

Servo base;
Servo shoulder;
Servo elbow;
Servo wrist_rot;
Servo wrist_ver;
Servo gripper;

int targetBase = 90;
int targetShoulder = 90;
int targetElbow = 90;
int targetWristVert = 90;
int targetWristRot = 90;
int targetGripper = 73;

void setup() {
  Serial.begin(9600);
  Braccio.begin();

  //(step delay, M1, M2, M3, M4, M5, M6);
  Braccio.ServoMovement(20, targetBase, targetShoulder, targetElbow,
                        targetWristVert, targetWristRot, targetGripper);
}

void loop() {
  /*
   Step Delay: a milliseconds delay between the movement of each servo.  Allowed
   values from 10 to 30 msec. M1=base degrees. Allowed values from 0 to 180
   degrees M2=shoulder degrees. Allowed values from 15 to 165 degrees M3=elbow
   degrees. Allowed values from 0 to 180 degrees M4=wrist vertical degrees.
   Allowed values from 0 to 180 degrees M5=wrist rotation degrees. Allowed
   values from 0 to 180 degrees M6=gripper degrees. Allowed values from 10 to 73
   degrees. 10: the toungue is open, 73: the gripper is closed.
  */

  while (Serial.available() > 0) {
    int motor = Serial.read();
    int amount = Serial.parseInt();

    if (amount < 0 || amount > 180) {
      continue;
    }

    switch (motor) {
      case 66:  // B
        targetBase = amount;
        break;
      case 83:  // S
        targetShoulder = amount;
        break;
      case 69:  // E
        targetElbow = amount;
        break;
      case 86:  // V
        targetWristVert = amount;
        break;
      case 82:  // R
        targetWristRot = amount;
        break;
      case 71:  // G
        targetGripper = amount;
        break;
      default:
        Serial.println(0);
    }

    Braccio.ServoMovement(20, targetBase, targetShoulder, targetElbow,
                          targetWristVert, targetWristRot, targetGripper);
    Serial.println(1);

    while (Serial.available()) Serial.read();
  }
}