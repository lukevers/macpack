package main

import "testing"

func TestDefaultConfig(t *testing.T) {
	t.Logf("%+v", defaultConfig())
}

func TestConfigCheckName(t *testing.T) {
	conf := defaultConfig()

	conf.Name = "foo-bar_boo"
	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.Name = ""
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.Name = "foo.bar"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestConfigCheckVersion(t *testing.T) {
	conf := defaultConfig()
	conf.Version = "1.0.0.42"

	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.Version = ""
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.Version = "1.42"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.Version = "1.42.f.1"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestConfigCheckID(t *testing.T) {
	conf := defaultConfig()

	conf.ID = "com.murlok.Test-ID"
	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.ID = ""
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.ID = "foo_bar"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestConfigCheckRole(t *testing.T) {
	conf := defaultConfig()

	conf.Role = "Editor"
	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.Role = "Viewer"
	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.Role = "Shell"
	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.Role = "None"
	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.Role = "Logger"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestConfigCheckDeploymentTarget(t *testing.T) {
	conf := defaultConfig()
	conf.DeploymentTarget = "10.15"

	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.DeploymentTarget = ""
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.DeploymentTarget = "1.a"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.DeploymentTarget = "9.12"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.DeploymentTarget = "10.10"
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestConfigCheckSupportedFiles(t *testing.T) {
	conf := defaultConfig()
	conf.SupportedFiles = []string{
		"public.jpeg",
	}

	if err := conf.check(); err != nil {
		t.Error(err)
	}

	conf.SupportedFiles = append(conf.SupportedFiles, "public_png")
	if err := conf.check(); err == nil {
		t.Error("checkConfig should return an error")
	}
}
