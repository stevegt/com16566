package com16566

import (
	"machine"
)

// https://github.com/sparkfun/SparkFun_Qwiic_Relay_Arduino_Library
// https://github.com/sparkfunX/Qwiic_Quad_Relay

// Default address
const I2C_ADDRESS_DEFAULT = 0x6D

// Address if address jumper is closed
const I2C_ADDRESS_JUMPER = 0x6C

const COMMAND_CHANGE_ADDRESS = 0xC7

// Commands that toggle the relays.
const TOGGLE_ONE = 0x01
const TOGGLE_TWO = 0x02
const TOGGLE_THREE = 0x03
const TOGGLE_FOUR = 0x04
const ALL_OFF = 0xA
const ALL_ON = 0xB
const TOGGLE_ALL = 0xC

// Commands to request the on/off state of the relay
const STATUS_ONE = 0x05
const STATUS_TWO = 0x06
const STATUS_THREE = 0x07
const STATUS_FOUR = 0x08

type COM16566 struct {
	Addr uint16
	I2c  *machine.I2C
}

func (r *COM16566) xeq(cmd byte) (rx byte, err error) {
	txbuf := []byte{cmd}
	rxbuf := make([]byte, 1)
	err = r.I2c.Tx(r.Addr, txbuf, rxbuf)
	rx = rxbuf[0]
	return
}

// Status returns the given relay's on/off state.
func (r *COM16566) Status(relay int) (bit bool, err error) {
	var cmd byte
	switch relay {
	case 1:
		cmd = STATUS_ONE
	case 2:
		cmd = STATUS_TWO
	case 3:
		cmd = STATUS_THREE
	case 4:
		cmd = STATUS_FOUR
	}
	rx, err := r.xeq(cmd)
	bit = (rx > 0)
	return
}

// Toggle the given relay. If the relay is on then it will turn
// it off, and if it is off then it will turn it on. Returns new
// relay state.
func (r *COM16566) Toggle(relay int) (bit bool, err error) {
	var cmd byte
	switch relay {
	case 1:
		cmd = TOGGLE_ONE
	case 2:
		cmd = TOGGLE_TWO
	case 3:
		cmd = TOGGLE_THREE
	case 4:
		cmd = TOGGLE_FOUR
	}
	_, err = r.xeq(cmd)
	if err != nil {
		return
	}
	bit, err = r.Status(relay)
	if err != nil {
		return
	}
	return
}

// Write bit to the given relay.
func (r *COM16566) Write(relay int, bit bool) (err error) {
	state, err := r.Status(relay)
	if err != nil {
		return
	}
	if state != bit {
		_, err = r.Toggle(relay)
		if err != nil {
			return
		}
	}
	return
}

// On turns on relay.
func (r *COM16566) On(relay int) (err error) {
	return r.Write(relay, true)
}

// On turns off relay.
func (r *COM16566) Off(relay int) (err error) {
	return r.Write(relay, false)
}
