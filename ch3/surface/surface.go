package main

import (
	"fmt"
	"math"
)

const (
	width  = 600
	height = 300
	//画布的大小像素
	cells   = 100                 //网格单元个数
	xyrange = 30.0                //坐标轴的范围（-xyrange..+xyrange）
	xyscale = width / 2 / xyrange //x 或 y 轴上每个单位长度的像素
	zscale  = height * 0.4        //z 轴上每个单位长度的像素
	angle   = math.Pi / 6         //x、y轴的角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf(`<svg xmlns='http://www.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>`, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf(`<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n`, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y) // 计算曲面高度
	// 将 xyz 等角投射到二维SVG绘图平面上，坐标是sx，sy
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 到 0，0 的距离
	return math.Sin(r)
}
