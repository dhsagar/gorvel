package controller

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"./../core/view"
)

func Hello(ctx *macaron.Context) {
	fmt.Println(ctx.Req.Host)
	fmt.Println("Printed Hello")
	view.Data(ctx, "DataName", 5)
	view.Show(ctx, "index")
}

