package main

import "testing"

func TestEntitlements(t *testing.T) {
	cap := capabilities{
		Network: networkCap{
			In:  true,
			Out: true,
		},
		Hardware: harwareCap{
			Camera:     true,
			Microphone: true,
			USB:        true,
			Printing:   true,
			Bluetooth:  true,
		},
		AppData: appDataCap{
			Contacts: true,
			Location: true,
			Calendar: true,
		},
		FileAccess: fileAccessCap{
			Downloads: fileReadWriteAccess,
			Pictures:  fileReadAccess,
			Music:     fileReadAccess,
			Movies:    fileReadAccess,
		},
	}

	defer deleteEntitlements()
	if err := createEntitlements(cap); err != nil {
		t.Error(err)
	}
}
