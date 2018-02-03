package main

import (
	"github.com/hashicorp/packer/packer"
	"github.com/mitchellh/multistep"
	"github.com/vmware/govmomi/object"
)

type StepCreateSnapshot struct {
	createSnapshot bool
}

func (s *StepCreateSnapshot) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	d := state.Get("driver").(*Driver)
	vm := state.Get("vm").(*object.VirtualMachine)

	if s.createSnapshot {
		ui.Say("Creating snapshot...")

		err := d.CreateSnapshot(vm)
		if err != nil {
			state.Put("error", err)
			return multistep.ActionHalt
		}
	}

	return multistep.ActionContinue
}

func (s *StepCreateSnapshot) Cleanup(state multistep.StateBag) {}
