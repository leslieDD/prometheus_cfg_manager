package alertmgr

import (
	"bytes"
	"encoding/base64"
	"errors"
	"pro_cfg_manager/config"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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
	respValues := ConvertValueV2(respData)
	p := plot.New()

	p.Title.Text = cr.Name
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		cr.Name, generateLineItems2(respValues))
	if err != nil {
		config.Log.Error(err)
		return "", err
	}

	// if err = p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
	// 	config.Log.Error(err)
	// 	return nil, err
	// }

	// titles := make([]string, 0, len(respValues))
	// for _, v := range respValues {
	// 	titles = append(titles, time.Unix(v.Unix, 0).Format("2006-01-02 15:04:05"))
	// }
	wio, err := p.WriterTo(12*vg.Inch, 6*vg.Inch, "png")
	if err != nil {
		config.Log.Error(err)
		return "", err
	}
	buffer := bytes.Buffer{}
	if _, err := wio.WriteTo(&buffer); err != nil {
		config.Log.Error(err)
		return "", err
	}
	// utils.WIoutilByte("aaa.png", buffer.Bytes())
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// generate random data for line chart
func generateLineItems2(vs []*Value) plotter.XYs {
	items := make(plotter.XYs, len(vs))
	for n, v := range vs {
		items[n].X = float64(v.Unix)
		items[n].Y = v.Val
	}
	return items
}
