package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/murlokswarm/log"
)

func generateIcon(conf Config) error {
	iconName := filepath.Join(conf.Name+".app", "Contents", "Resources", conf.Icon)
	iconsetName := epureIconName(conf.Icon) + ".iconset"
	iconsetName = filepath.Join(conf.Name+".app", "Contents", "Resources", iconsetName)

	f, err := os.Open(iconName)
	if err != nil {
		return err
	}

	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return err
	}

	if err := os.Mkdir(iconsetName, os.ModeDir|0755); err != nil {
		return err
	}

	defer os.RemoveAll(iconsetName)

	createIconsetImg(img, iconsetName, 512, 512, 2)
	createIconsetImg(img, iconsetName, 512, 512, 1)
	createIconsetImg(img, iconsetName, 256, 256, 2)
	createIconsetImg(img, iconsetName, 256, 256, 1)
	createIconsetImg(img, iconsetName, 128, 128, 2)
	createIconsetImg(img, iconsetName, 128, 128, 1)
	createIconsetImg(img, iconsetName, 32, 32, 2)
	createIconsetImg(img, iconsetName, 32, 32, 1)
	createIconsetImg(img, iconsetName, 16, 16, 2)
	createIconsetImg(img, iconsetName, 16, 16, 1)

	return execCmd("iconutil", "-c", "icns", iconsetName)
}

func epureIconName(name string) string {
	name = filepath.Base(name)
	return strings.TrimSuffix(name, filepath.Ext(name))
}

func createIconsetImg(img image.Image, iconsetName string, w int, h int, m int) {
	rimg := imaging.Resize(img, w*m, h*m, imaging.Lanczos)
	name := ""

	switch m {
	case 0:
		log.Errorf("multiplier can't be 0: %v", m)
		return

	case 1:
		name = filepath.Join(iconsetName, fmt.Sprintf("icon_%vx%v.png", w, h))

	default:
		name = filepath.Join(iconsetName, fmt.Sprintf("icon_%vx%v@%vx.png", w, h, m))
	}

	f, err := os.Create(name)
	if err != nil {
		log.Error(err)
		return
	}

	defer f.Close()

	if err = png.Encode(f, rimg); err != nil {
		log.Error(err)
	}
}
