// Package meinvcalendar 美女日历
package meinvcalendar

import (
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	control.Register("meinvcalendar", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Brief:            "美女日历",
		Help: "- /启用 meinvcalendar\n" +
			"- /禁用 meinvcalendar\n" +
			"- 记录在\"30 8 * * *\"触发的指令\n" +
			"   - 美女日历",
	}).OnFullMatch("美女日历").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData("https://api.vvhan.com/api/mobil.girl")
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
}
