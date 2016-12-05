package main

import "text/template"
import "path/filepath"
import "os"

const plistTmpl = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>   
    <key>CFBundleVersion</key>
	<string>{{.Version}}</string>

    <key>CFBundleName</key>
	<string>{{.Name}}</string>
    
    <key>CFBundleDisplayName</key>
	<string>{{.Name}}</string>
    
    <key>CFBundleExecutable</key>
	<string>{{.Name}}</string>
	
    <key>CFBundleIdentifier</key>
	<string>{{.ID}}</string>
    
    <key>CFBundlePackageType</key>
	<string>APPL</string>
    
    <key>CFBundleSignature</key>
	<string>????</string>
    
    <key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>

    <key>CFBundleSupportedPlatforms</key>
    <array>
	    <string>MacOSX</string>
	</array>

    <key>DTCompiler</key>
	<string>com.apple.compilers.llvm.clang.1_0</string>

    <key>LSMinimumSystemVersion</key>
	<string>{{.OSMinVersion}}</string>

    <key>CFBundleIconFile</key>
	<string>{{.Icon}}</string>

    <key>LSUIElement</key>
	<string>true</string>

    <key>NSPrincipalClass</key>
	<string>NSApplication</string>

	<key>NSAppTransportSecurity</key>
	<dict>
        <key>NSAllowsArbitraryLoads</key>
        <true/>
    </dict>

	{{if .Sandbox}}<key>com.apple.security.app-sandbox</key>
    <true/>{{end}}

	<key>CFBundleDocumentTypes</key>
	<array>
		<dict>
			<key>CFBundleTypeName</key>
        	<string>Supported files</string>
			<key>CFBundleTypeRole</key>
			<string>{{.Role}}</string>
			<key>LSItemContentTypes</key>
			<array>
				{{range .SupportedFiles}}<string>{{.}}</string>
				{{end}}
			</array>
		</dict>
	</array>
</dict>
</plist>
    `

func createPlist(conf Config) error {
	plistName := filepath.Join(conf.Name+".app", "Contents", "Info.plist")

	f, err := os.Create(plistName)
	if err != nil {
		return err
	}

	defer f.Close()

	if len(conf.Icon) != 0 {
		conf.Icon = epureIconName(conf.Icon)
	}

	tmpl := template.Must(template.New("plist").Parse(plistTmpl))
	return tmpl.Execute(f, conf)
}
