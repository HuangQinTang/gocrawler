package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contentBts, err := ioutil.ReadFile("city.txt")
	if err != nil {
		panic(err)
	}
	correct := map[string]string{
		"小扬姐":"http://album.zhenai.com/u/1958903678",
		"遇见你":"http://album.zhenai.com/u/1396738538",
		"期待有缘人":"http://album.zhenai.com/u/1774343032",
		"最美不过能遇见你":"http://album.zhenai.com/u/1529742129",
		"无人知晓":"http://album.zhenai.com/u/1700301913",
		"远离回忆":"http://album.zhenai.com/u/1398432707",
		"董懂":"http://album.zhenai.com/u/1821376705",
		"松哥":"http://album.zhenai.com/u/1491731990",
		"非诚勿扰":"http://album.zhenai.com/u/1693313153",
		"普普通通": "http://album.zhenai.com/u/1761252184",
		"不想": "http://album.zhenai.com/u/1426975040",
		"愿得一人心的美好": "http://album.zhenai.com/u/1677398857",
		"一路是蓝": "http://album.zhenai.com/u/1292525744",
		"山有木兮沫": "http://album.zhenai.com/u/1834384219",
		"随缘": "http://album.zhenai.com/u/1918103401",
		"找个真心人": "http://album.zhenai.com/u/1122288128",
		"Lily": "http://album.zhenai.com/u/1998372165",
		"高级杂工": "http://album.zhenai.com/u/1839731206",
		"山水有相逢": "http://album.zhenai.com/u/1968078839",
		"回家": "http://album.zhenai.com/u/1977189312",
	}
	data := ParseCity(contentBts)
	for k, v := range data.Items {
		if correct[v.(string)] != data.Requests[k].Url {
			t.Errorf("error %s url is %s, not is %s", v, correct[v.(string)], data.Requests[k].Url)
		}
	}
}
