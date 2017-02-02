package main

import (
	"os"
	"text/template"
)

const entitlements = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>com.apple.security.app-sandbox</key>
	<true/>

    <!-- Network -->
    {{if .Network.In}}
    <key>com.apple.security.network.server</key>
	<true/>
    {{end}}
    {{if .Network.Out}}
	<key>com.apple.security.network.client</key>
	<true/>
    {{end}}

    <!-- Hadrware -->
    {{if .Hardware.Camera}}
	<key>com.apple.security.device.camera</key>
	<true/>
    {{end}}
    {{if .Hardware.Microphone}}
	<key>com.apple.security.device.microphone</key>
	<true/>
    {{end}}
    {{if .Hardware.USB}}
	<key>com.apple.security.device.usb</key>
	<true/>
    {{end}}
    {{if .Hardware.Printing}}
	<key>com.apple.security.print</key>
	<true/>
    {{end}}
    {{if .Hardware.Bluetooth}}
	<key>com.apple.security.device.bluetooth</key>
	<true/>
    {{end}}

    <!-- AppData -->
    {{if .AppData.Contacts}}
	<key>com.apple.security.personal-information.addressbook</key>
	<true/>
    {{end}}
    {{if .AppData.Location}}
	<key>com.apple.security.personal-information.location</key>
	<true/>
    {{end}}
    {{if .AppData.Calendar}}
	<key>com.apple.security.personal-information.calendars</key>
	<true/>
    {{end}}

    <!-- FileAccess -->
    {{if len .FileAccess.UserSelected}}
    <key>com.apple.security.files.user-selected.{{.FileAccess.UserSelected}}</key>
	<true/>
    {{end}}
    {{if len .FileAccess.Downloads}}
	<key>com.apple.security.files.downloads.{{.FileAccess.Downloads}}</key>
	<true/>
    {{end}}
    {{if len .FileAccess.Pictures}}
	<key>com.apple.security.assets.pictures.{{.FileAccess.Pictures}}/key>
	<true/>
    {{end}}
    {{if len .FileAccess.Music}}
	<key>com.apple.security.assets.music.{{.FileAccess.Music}}</key>
	<true/>
    {{end}}
    {{if len .FileAccess.Movies}}
	<key>com.apple.security.assets.movies.{{.FileAccess.Movies}}/key>
	<true/>
    {{end}}
</dict>
</plist>
`

func createEntitlements(cap capabilities) error {
	name := "mac.entitlements"
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("entitlements").Parse(entitlements))
	return tmpl.Execute(f, cfg.Capabilities)
}

func deleteEntitlements() {
	os.Remove("mac.entitlements")
}
