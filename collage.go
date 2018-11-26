package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func collage(in []string) {
	images := make([]image.Image, len(in))
	for i, item := range in {
		file, _ := os.Open(item)
		m, _, _ := image.Decode(file)
		images[i] = m
	}

	cols := 16
	cell := 128
	rows := (len(images) + cols - 1) / cols
	dst := image.NewRGBA(image.Rect(0, 0, cell*cols, cell*rows))
	draw.Draw(dst, dst.Bounds(), image.NewUniform(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}), image.Point{}, draw.Src)

	for i, m := range images {
		col := i % cols
		row := i / cols

		sz := m.Bounds().Size()
		dz := sz
		if sz.X > sz.Y {
			dz.X = cell
			dz.Y = cell * sz.Y / sz.X
		} else {
			dz.Y = cell
			dz.X = cell * sz.X / sz.Y
		}

		z := image.Point{cell * col, cell * row}
		r := image.Rectangle{
			Min: z,
			Max: z.Add(dz),
		}
		r = r.Add(image.Point{cell / 2, cell / 2}).
			Sub(image.Point{dz.X / 2, dz.Y / 2})

		draw.CatmullRom.Scale(dst, r, m, m.Bounds(), draw.Over, nil)
	}

	result, err := os.Create("collage.jpg")
	if err != nil {
		log.Println(err)
		return
	}

	if err := jpeg.Encode(result, dst, &jpeg.Options{Quality: 90}); err != nil {
		log.Println(err)
		return
	}
}
