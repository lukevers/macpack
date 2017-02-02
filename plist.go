package main

import (
	"path/filepath"
	"text/template"
)

import "os"

const plistTmpl = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>{{.DevRegion}}</string>

	<key>CFBundleExecutable</key>
	<string>{{.Name}}</string>

	<key>CFBundleIconFile</key>
	<string>{{.Icon}}</string>

	<key>CFBundleIdentifier</key>
	<string>{{.ID}}</string>

	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>

	<key>CFBundleName</key>
	<string>{{.Name}}</string>

	<key>CFBundlePackageType</key>
	<string>APPL</string>

	<key>CFBundleSupportedPlatforms</key>
	<array>
		<string>MacOSX</string>
	</array>

	<key>CFBundleShortVersionString</key>
	<string>{{.Version}}</string>

	<key>CFBundleVersion</key>
	<string>{{.BuildNumber}}</string>

	<key>LSMinimumSystemVersion</key>
	<string>{{.DeploymentTarget}}</string>

	<key>LSApplicationCategoryType</key>
	<string>{{.Category}}</string>

	<key>NSHumanReadableCopyright</key>
	<string>{{html .Copyright}}</string>

	<key>NSPrincipalClass</key>
	<string>NSApplication</string>

	<key>NSAppTransportSecurity</key>
	<dict>
		<key>NSAllowsArbitraryLoadsInWebContent</key>
		<true/>
	</dict>

	<key>CFBundleDocumentTypes</key>
	<array>
		<dict>
			<key>CFBundleTypeName</key>
			<string>Supported files</string>
			<key>CFBundleTypeRole</key>
			<string>{{.Role}}</string>
			<key>LSItemContentTypes</key>
			<array>{{range .SupportedFiles}}
				<string>{{.}}</string>{{end}}
			</array>
		</dict>
	</array>
</dict>
</plist>
    `

func createPlist(cfg Config) error {
	plistName := filepath.Join(cfg.appName(), "Contents", "Info.plist")
	f, err := os.Create(plistName)
	if err != nil {
		return err
	}
	defer f.Close()

	if len(cfg.Icon) != 0 {
		cfg.Icon = epureIconName(cfg.Icon)
	}

	tmpl := template.Must(template.New("plist").Parse(plistTmpl))
	return tmpl.Execute(f, cfg)
}
