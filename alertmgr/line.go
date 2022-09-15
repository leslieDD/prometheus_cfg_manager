package alertmgr

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"time"

	"github.com/vicanso/go-charts/v2"
)

func ChartLine(cr *CronRule) (string, error) {
	var installFont charts.OptionFunc
	if config.Cfg.Font != "" {
		buf, err := ioutil.ReadFile(config.Cfg.Font)
		if err != nil {
			config.Log.Error(err)
			installFont = charts.FontFamilyOptionFunc("roboto")
		} else {
			err = charts.InstallFont("noto", buf)
			if err != nil {
				config.Log.Error(err)
				installFont = charts.FontFamilyOptionFunc("roboto")
			} else {
				installFont = charts.FontFamilyOptionFunc("noto")
			}
		}
	}
	cr.End = time.Now() // 时间要反向处理
	cr.Start = time.Unix(cr.End.Unix()-ConvertTime(cr.NearTime, cr.Unit), 0)
	api := ConvertApi(cr.API)
	reqUrl := CreateReqUrl(api, cr.Rule, cr.Start, cr.End)
	respData := DoReq(reqUrl)
	if respData == nil {
		config.Log.Error("get chart data error")
		return "", errors.New("get chart data error")
	}

	if len(respData.Data.Result) == 0 {
		config.Log.Error("no data found")
		return "", errors.New("no data found")
	}

	x, y, max := ConvertValueV4(respData)
	p, err := charts.LineRender(
		y,
		charts.TitleTextOptionFunc("Line"),
		charts.XAxisDataOptionFunc(x),
		installFont,
		func(opt *charts.ChartOption) {
			opt.Title = charts.TitleOption{Text: cr.Name}
			opt.Height = 600
			opt.Width = 1000
			opt.Legend.Padding = charts.Box{
				Top:    5,
				Bottom: 10,
			}
			opt.SymbolShow = charts.FalseFlag()
			opt.LineStrokeWidth = 1
			opt.YAxisOptions = []charts.YAxisOption{
				{
					Max: charts.NewFloatPoint(max),
				},
			}
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
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf), nil
}

func ChartLineTest() (string, error) {
	var installFont charts.OptionFunc
	if config.Cfg.Font != "" {
		buf, err := ioutil.ReadFile(config.Cfg.Font)
		if err != nil {
			config.Log.Error(err)
			installFont = charts.FontFamilyOptionFunc("roboto")
		} else {
			err = charts.InstallFont("noto", buf)
			if err != nil {
				config.Log.Error(err)
				installFont = charts.FontFamilyOptionFunc("roboto")
			} else {
				installFont = charts.FontFamilyOptionFunc("noto")
			}
		}
	}
	values := [][]float64{
		utils.RandFloat64N(7),
	}
	p, err := charts.LineRender(
		values,
		charts.TitleTextOptionFunc("Line"),
		installFont,
		charts.XAxisDataOptionFunc([]string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		}),
		charts.LegendLabelsOptionFunc([]string{
			"Email",
		}, "50"),
		func(opt *charts.ChartOption) {
			opt.Legend.Padding = charts.Box{
				Top:    5,
				Bottom: 10,
			}
			opt.FillArea = true
			opt.Title = charts.TitleOption{Text: "发送测试邮件（send test email）"}
			opt.Height = 600
			opt.Width = 1000
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
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf), nil
}
