package ini

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	bEmpty   = ""
	bComment = "#"
	bEqual   = "="
)

type IniConfig struct {
	fileName string
	data     map[string]string
	sync.RWMutex
}

func NewIniConfig(fileName string) *IniConfig {
	ini := &IniConfig{fileName, make(map[string]string), sync.RWMutex{}}
	ini.parseFile()
	return ini
}

func (p *IniConfig) parseFile() {
	file, err := os.Open(p.fileName)
	if err != nil {
		fmt.Println("open config file err:", err)
		return
	}
	defer file.Close()
	p.Lock()
	defer p.Unlock()
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		//空行，以#开头注释，或者不是**=**格式的行都不处理
		if bEmpty == line || strings.HasPrefix(line, bComment) || !strings.Contains(line, bEqual) {
			continue
		}
		keyValue := strings.SplitN(line, bEqual, 2)
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		p.data[key] = value
	}

	return
}

func (p *IniConfig) DefaultInt(key string, defaultValue int) (ret int) {
	value, err := p.getData(key)
	if err != nil {
		ret = defaultValue
		return
	}
	ret, err = strconv.Atoi(value)
	if err != nil {
		ret = defaultValue
		return
	}
	return
}

func (p *IniConfig) DefaultInt64(key string, defaultValue int64) (ret int64) {
	value, err := p.getData(key)
	if err != nil {
		ret = defaultValue
		return
	}
	ret, err = strconv.ParseInt(value, 10, 64)
	if err != nil {
		ret = defaultValue
		return
	}
	return
}

func (p *IniConfig) DefaultString(key string, defaultValue string) (ret string) {
	ret, err := p.getData(key)
	if err != nil {
		ret = defaultValue
		return
	}
	return
}

func (p *IniConfig) DefaultBool(key string, defaultValue bool) (ret bool) {
	value, err := p.getData(key)
	if err != nil {
		ret = defaultValue
		return
	}
	ret, err = strconv.ParseBool(value)
	if err != nil {
		ret = defaultValue
		return
	}
	return
}

func (p *IniConfig) DefaultFloat64(key string, defaultValue float64) (ret float64) {
	value, err := p.getData(key)
	if err != nil {
		ret = defaultValue
		return
	}
	ret, err = strconv.ParseFloat(value, 64)
	if err != nil {
		ret = defaultValue
		return
	}
	return
}

func (p *IniConfig) getData(key string) (value string, err error) {
	p.Lock()
	defer p.Unlock()
	value, ok := p.data[key]
	if !ok {
		err = fmt.Errorf("failed to get data where key=%s", key)
	}
	return
}
