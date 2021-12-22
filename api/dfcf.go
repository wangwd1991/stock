package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"stock"
)

type dfcf struct {
	host string
	uri  string
}

type dfcfData struct {
	Rc   int `json:"rc"`
	Rt   int `json:"rt"`
	Svr  int `json:"svr"`
	Lt   int `json:"lt"`
	Full int `json:"full"`
	Data struct {
		F43  float64     `json:"f43"`
		F44  float64     `json:"f44"`
		F45  float64     `json:"f45"`
		F46  float64     `json:"f46"`
		F47  float64     `json:"f47"`
		F48  float64     `json:"f48"`
		F49  int         `json:"f49"`
		F50  float64     `json:"f50"`
		F51  float64     `json:"f51"`
		F52  float64     `json:"f52"`
		F55  float64     `json:"f55"`
		F57  string      `json:"f57"`
		F58  string      `json:"f58"`
		F60  float64     `json:"f60"`
		F62  int         `json:"f62"`
		F71  float64     `json:"f71"`
		F92  float64     `json:"f92"`
		F116 float64     `json:"f116"`
		F117 float64     `json:"f117"`
		F135 float64     `json:"f135"`
		F136 float64     `json:"f136"`
		F137 float64     `json:"f137"`
		F138 float64     `json:"f138"`
		F139 interface{} `json:"f139"`
		F140 float64     `json:"f140"`
		F141 float64     `json:"f141"`
		F142 float64     `json:"f142"`
		F143 float64     `json:"f143"`
		F144 float64     `json:"f144"`
		F145 float64     `json:"f145"`
		F146 float64     `json:"f146"`
		F147 float64     `json:"f147"`
		F148 float64     `json:"f148"`
		F149 float64     `json:"f149"`
		F161 int         `json:"f161"`
		F162 float64     `json:"f162"`
		F163 float64     `json:"f163"`
		F164 float64     `json:"f164"`
		F167 float64     `json:"f167"`
		F168 float64     `json:"f168"`
		F169 float64     `json:"f169"`
		F170 float64     `json:"f170"`
		F31  interface{} `json:"f31"`
		F32  interface{} `json:"f32"`
		F33  interface{} `json:"f33"`
		F34  interface{} `json:"f34"`
		F35  interface{} `json:"f35"`
		F36  interface{} `json:"f36"`
		F37  interface{} `json:"f37"`
		F38  interface{} `json:"f38"`
		F39  interface{} `json:"f39"`
		F40  interface{} `json:"f40"`
		F19  interface{} `json:"f19"`
		F20  interface{} `json:"f20"`
		F17  interface{} `json:"f17"`
		F18  interface{} `json:"f18"`
		F15  interface{} `json:"f15"`
		F16  interface{} `json:"f16"`
		F13  interface{} `json:"f13"`
		F14  interface{} `json:"f14"`
		F11  interface{} `json:"f11"`
		F12  interface{} `json:"f12"`
	} `json:"data"`
}

// 0.300059

func (d *dfcf) GetByCode(sk *stock.Stock) error {
	params := make(map[string]string)
	params["ut"] = "fa5fd1943c7b386f172d6893dbfba10b"
	params["invt"] = "2"
	params["fltt"] = "2"
	params["fields"] = "f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163,f116,f60,f45,f52,f50,f48,f167,f117,f71,f161,f49,f530,f135,f136,f137,f138,f139,f141,f142,f144,f145,f147,f148,f140,f143,f146,f149,f55,f62,f162,f92,f173,f104,f105,f84,f85,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f107,f111,f86,f177,f78,f110,f260,f261,f262,f263,f264,f267,f268,f250,f251,f252,f253,f254,f255,f256,f257,f258,f266,f269,f270,f271,f273,f274,f275,f127,f199,f128,f193,f196,f194,f195,f197,f80,f280,f281,f282,f284,f285,f286,f287,f292,f293,f181,f294,f295,f279,f288"
	code := sk.Code
	if code[0] == '3' {
		params["secid"] = "0." + code
	} else if code[0] == '6' {
		params["secid"] = "1." + code
	} else if code[0] == '0' {
		params["secid"] = "0." + code
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", d.host+d.uri, nil)
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("call return code: %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := &dfcfData{}

	err = json.Unmarshal(content, data)
	if err != nil {
		return err
	}
	sk.Open = stock.MyFloat64(data.Data.F46)
	sk.Yesterday = stock.MyFloat64(data.Data.F60)
	sk.Current = stock.MyFloat64(data.Data.F43)
	sk.Max = stock.MyFloat64(data.Data.F44)
	sk.Min = stock.MyFloat64(data.Data.F45)
	sk.Gains = stock.MyFloat64(data.Data.F170)

	if sk.Current > sk.Yesterday {
		sk.Status = 1
	} else if sk.Current == sk.Yesterday {
		sk.Status = 0
	} else {
		sk.Status = -1
	}
	if sk.Number > 0 {
		sk.Earn = sk.Current.Earn(sk.Cost, sk.Number)
		sk.CurEarn = sk.Current.Earn(sk.Yesterday, sk.Number)
	}
	return nil
}

func (d *dfcf) GetByName(sk *stock.Stock) error {
	panic("implement me")
}

func NewDFCF() Service {
	return &dfcf{
		host: "http://push2.eastmoney.com",
		uri:  "/api/qt/stock/get",
	}
}
