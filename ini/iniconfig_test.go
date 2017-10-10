package ini

import (
	"os"
	"testing"
)

func TestIni(t *testing.T) {
	var (
		inicontext = `
appname = zhunxun
#comment
httpport =     8080
mysqlport =3360

PI = 3.1415926
ok= true
`
		keyValue = map[string]interface{}{
			"appname":   "zhunxun",
			"httpport":  8080,
			"mysqlport": int64(3360),
			"PI":        float64(3.1415926),
			"ok":        true,
		}
	)

	f, err := os.Create("app.conf")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(inicontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove("app.conf")

	ini := NewIniConfig("app.conf")

	for k, v := range keyValue {
		switch v.(type) {
		case int:
			if ini.DefaultInt(k, 123) != v {
				t.Errorf("DefaultInt is Fail")
			}
		case int64:
			if ini.DefaultInt64(k, int64(123)) != v {
				t.Errorf("DefaultInt64 is Fail")
			}

		case string:
			if ini.DefaultString(k, "test") != v {
				t.Errorf("DefaultString is Fail")
			}
		case bool:
			if ini.DefaultBool(k, false) != v {
				t.Errorf("DefaultBool is Fail")
			}
		case float64:
			if ini.DefaultFloat64(k, float64(123.123)) != v {
				t.Errorf("Defaultfloat64 is Fail")
			}
		}
	}
}
