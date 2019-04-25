package html

import (
	"fmt"
	"math"
)

const(
	width,height = 600,320   //画布大小（像素）
	cells = 100	         //网格单元格数
	xyrange = 30.0       //轴范围（-xyrange .. + xyrange）
	xyscale = width / 2 / xyrange    //每x或y单位的像素
	zscale = height * 0.4     //每z单位的像素数
	angle = math.Pi     //x，y轴角度（= 30°）
)

var sin30,cos30 = math.Sin(angle),math.Cos(angle) //sin(30),cos(30)

func main()  {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke:grey;fill:white;stroke-width:0.7'"+"width='%d' " +
		"height='%d'>",width,height)

	for i := 0;i < cells ; i++ {
		for j := 0;j < cells;j++  {
			ax,ay := corner(i+1,j)
			bx,by := corner(i,j)
			cx,cy := corner(i,j+1)
			dx,dy := corner(i+1,j+1)
			fmt.Printf("<polygon points = '%g,%g,%g,%g,%g,%g,%g,%g/>\n",
				ax,ay,bx,by,cx,cy,dx,dy)//多边形点
		}
		fmt.Println("</svg>")
	}
}
func corner(i,j int) (float64,float64){
	//在单元格（i，j）的角上找到点（x，y）
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算表面高度z。
	z := f(x,y)

	//Project（x，y，z）等距到2-D SVG画布（sx，sy）。
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x-y)*sin30*xyscale - z*zscale

	return sx,sy
}
func f(x,y float64) float64  {
	r := math.Hypot(x,y)//距离（0,0）
	return math.Sin(r)
}


