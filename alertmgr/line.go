package alertmgr

import (
	"bytes"
	"errors"
	"pro_cfg_manager/config"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func ChartLine(cr *CronRule) ([]byte, error) {
	reqUrl := CreateReqUrl(cr.API, cr.Rule, cr.Start, cr.End)
	respData := DoReq(reqUrl)
	if respData == nil {
		return nil, errors.New("get chart data error")
	}
	respValues := ConvertValue(respData)
	p := plot.New()

	p.Title.Text = cr.Name
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		cr.Name, generateLineItems2(respValues))
	if err != nil {
		config.Log.Error(err)
		return nil, err
	}

	// if err = p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
	// 	config.Log.Error(err)
	// 	return nil, err
	// }

	// titles := make([]string, 0, len(respValues))
	// for _, v := range respValues {
	// 	titles = append(titles, time.Unix(v.Unix, 0).Format("2006-01-02 15:04:05"))
	// }
	wio, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		config.Log.Error(err)
		return nil, err
	}
	buffer := bytes.Buffer{}
	wio.WriteTo(&buffer)
	return buffer.Bytes(), nil
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
