// Package girl images
package girl

import (
	"regexp"

	"github.com/tidwall/gjson"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
)

var (
	ua       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36"
	coserURL = "https://api.r10086.com/img-api.php?type=极品美女图片"
	datestr  = regexp.MustCompile(`/\d{4}-\d{2}-\d{2}/`)
)

func init() {
	control.Register("girl", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "美女",
		Help:             "- girl",
	}).ApplySingle(ctxext.DefaultSingle).OnFullMatch("girl").SetBlock(true).Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("少女祈祷中......"))
			data, err := web.RequestDataWith(web.NewDefaultClient(), coserURL, "GET", "", ua)
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			text := gjson.Get(helper.BytesToString(data), "data.Title").String()
			m := message.Message{ctxext.FakeSenderForwardNode(ctx, message.Text(text))}
			ds := ""
			gjson.Get(helper.BytesToString(data), "data.data").ForEach(func(_, value gjson.Result) bool {
				if ds == "" {
					ds = datestr.FindString(value.String())
				} else if ds != datestr.FindString(value.String()) {
					return false
				}
				m = append(m, ctxext.FakeSenderForwardNode(ctx, message.Image(value.String())))
				return true
			})
			if id := ctx.Send(m).ID(); id == 0 {
				ctx.SendChain(message.Text("ERROR: 可能被风控或下载图片用时过长，请耐心等待"))
			}
		})
}
