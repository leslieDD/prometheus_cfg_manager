package alertmgr

import (
	"encoding/base64"
	"errors"
	"pro_cfg_manager/config"
	"time"

	"github.com/vicanso/go-charts/v2"
	"gonum.org/v1/plot/plotter"
)

func ChartLine(cr *CronRule) (string, error) {
	cr.End = time.Now() // 时间要反向处理
	cr.Start = time.Unix(cr.End.Unix()-ConvertTime(cr.NearTime, cr.Unit), 0)
	api := ConvertApi(cr.API)
	reqUrl := CreateReqUrl(api, cr.Rule, cr.Start, cr.End)
	respData := DoReq(reqUrl)
	if respData == nil {
		return "", errors.New("get chart data error")
	}

	if len(respData.Data.Result) == 0 {
		return "", errors.New("no data found")
	}

	x, y, _ := ConvertValueV4(respData)

	// max := float64(3)
	p, err := charts.LineRender(
		y,
		// charts.YAxisOptionFunc(charts.YAxisOption{Max: &max}),
		charts.TitleTextOptionFunc("Line"),
		charts.XAxisDataOptionFunc(x),
		// charts.LegendLabelsOptionFunc(labels, "50"),
		func(opt *charts.ChartOption) {
			opt.Legend.Padding = charts.Box{
				Top:    5,
				Bottom: 10,
			}
			opt.SymbolShow = charts.FalseFlag()
			opt.LineStrokeWidth = 1
		},
	)
	if err != nil {
		config.Log.Error(err)
		return "", err
	}
	buf, err := p.Bytes()
	if err != nil {
		config.Log.Error(err)
		return "", err
	}
	// err = writeFile(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// p := plot.New()

	// p.Title.Text = cr.Rule
	// p.X.Label.Text = "X"
	// p.Y.Label.Text = "Y"

	// xs := []string{}

	// for _, line := range respData.Data.Result {
	// 	xys := ConvertValueV3(line)
	// 	if err := plotutil.AddLinePoints(p, xys); err != nil {
	// 		config.Log.Error(err)
	// 		return "", err
	// 	}
	// 	// xs = append(xs, line.Metric["instance"])
	// 	break
	// }

	// ys := []string{}
	// line := respData.Data.Result[0]
	// for _, v := range line.Values {
	// 	ys = append(ys, time.Unix(int64(v[0].(float64)), 0).Local().Format("2006-01-02 15:04:05"))
	// }
	// respValues := ConvertValueV2(respData)

	// err := plotutil.AddLinePoints(p,
	// 	generateLineItems2(respValues))
	// if err != nil {
	// 	config.Log.Error(err)
	// 	return "", err
	// }

	// xs := generateLineX(respValues)

	// p.NominalX(ys...)
	// p.NominalY(xs...)

	// plotter

	// if err = p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
	// 	config.Log.Error(err)
	// 	return nil, err
	// }

	// titles := make([]string, 0, len(respValues))
	// for _, v := range respValues {
	// 	titles = append(titles, time.Unix(v.Unix, 0).Format("2006-01-02 15:04:05"))
	// }
	// wio, err := p.WriterTo(12*vg.Inch, 6*vg.Inch, "png")
	// if err != nil {
	// 	config.Log.Error(err)
	// 	return "", err
	// }
	// buffer := bytes.Buffer{}
	// if _, err := wio.WriteTo(&buffer); err != nil {
	// 	config.Log.Error(err)
	// 	return "", err
	// }
	// utils.WIoutilByte("aaa.png", buffer.Bytes())
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf), nil
}

func generateLineX(vs []*Value) []string {
	xs := []string{}
	for _, v := range vs {
		xs = append(xs, time.Unix(v.Unix, 0).Format("2006-01-02 15:04:05"))
	}
	return xs
}

// generate random data for line chart
func generateLineItems2(vs []*Value) plotter.XYs {

	// x := plotter.NewGlyphBoxes()
	// x.Plot()

	items := make(plotter.XYs, len(vs))
	for n, v := range vs {
		items[n].X = float64(v.Unix)
		items[n].Y = v.Val
	}
	return items
}
