package main

func signPackage(cfg Config) error {
	return execCmd("codesign",
		"--entitlements",
		"mac.entitlements",
		"-s",
		cfg.SignID,
		cfg.appName(),
	)
}

func signCheck(cfg Config) error {
	return execCmd("codesign",
		"--verify",
		"--deep",
		"--strict",
		"--verbose=2",
		cfg.appName(),
	)
}

func spctlCheck(cfg Config) error {
	return execCmd("spctl",
		"-a",
		"-t",
		"exec",
		"-vv",
		cfg.appName(),
	)
}

func packageForAppStore(cfg Config) error {
	return execCmd("productbuild",
		"--component",
		cfg.appName(),
		"/Applications",
		"--sign",
		cfg.SignID,
		cfg.Name+".pkg",
	)
}
