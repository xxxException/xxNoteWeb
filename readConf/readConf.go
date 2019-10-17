package readConf

import (
	"bufio"
	"os"
	"strings"
)

func ReadConf(fileUrl string) (map[string]string, error) {
	configMap := make(map[string]string)

	file, openFileErr := os.Open(fileUrl)
	if openFileErr != nil {
		println(openFileErr) //log
		return nil, openFileErr
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineStr, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		lineStr = strings.TrimSpace(lineStr)
		index := strings.Index(lineStr, "=")

		key := strings.TrimSpace(lineStr[:index])
		if len(key) == 0 {
			continue
		}

		value := strings.TrimSpace(lineStr[index+1:])
		if len(value) == 0 {
			continue
		}

		configMap[key] = value
	}

	return configMap, nil
}
