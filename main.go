package main

import (
	"flag"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func load(name string) (image.Image, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func save(name string, img image.Image) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}

func toNRGBA(img image.Image) (*image.NRGBA, error) {
	if dst, ok := img.(*image.NRGBA); ok {
		return dst, nil
	}
	dst := image.NewNRGBA(img.Bounds())
	draw.Draw(dst, dst.Bounds(), img, image.ZP, draw.Src)
	return dst, nil
}

func punch2(m *image.NRGBA) {
	p := m.Bounds().Max
	x, y := p.X-1, p.Y-1
	c := m.NRGBAAt(x, y)
	c.A = 0
	m.SetNRGBA(x, y, c)
}

func punch(srcname string, dstname string) error {
	src, err := load(srcname)
	if err != nil {
		return err
	}
	dst, err := toNRGBA(src)
	if err != nil {
		return err
	}
	punch2(dst)
	err = save(dstname, dst)
	if err != nil {
		return err
	}
	return nil
}

func modname(s string) string {
	ext := filepath.Ext(s)
	name := s[:len(s)-len(ext)]
	return name + "+pash.png"
}

func main() {
	flag.Parse()
	code := 0
	for _, s := range flag.Args() {
		err := punch(s, modname(s))
		if err != nil {
			log.Printf("failed to punch %q: %s", s, err)
			code = 1
		}
	}
	if code != 0 {
		os.Exit(code)
	}
}
