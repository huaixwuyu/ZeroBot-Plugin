// Package dailycalendar 日历
package dailycalendar

import (
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	control.Register("dailycalendar", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Brief:            "日历",
		Help: "- /启用 dailycalendar\n" +
			"- /禁用 dailycalendar\n" +
			"- 记录在\"00 7 * * *\"触发的指令\n" +
			"   - 日历",
	}).OnFullMatch("日历").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData("https://api.vvhan.com/api/60s")
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
}
