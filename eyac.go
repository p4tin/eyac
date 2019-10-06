package eyac

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func LoadConfig(filename string, settings interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ProcessedYamlConfig string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newValue := getEnvOrDefault(scanner.Text())
		ProcessedYamlConfig = fmt.Sprintf("%s\n%s", ProcessedYamlConfig, newValue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(ProcessedYamlConfig), settings)
	if err != nil {
		panic(err)
	}
}
/*

ENV
ENV_BOOL
ENV_INT
ENV_LIST

 */
func getEnvOrDefault(line string) string {
	i := strings.Index(line, "#")
	if i > 0 {
		line = strings.TrimSpace(line[:i])
	}
	parts := strings.Split(line, ":")
	if len(parts) < 2 || parts[1] == "" {
		return line
	}

	values := strings.Split(strings.TrimSpace(parts[1]), ",")
	value := strings.TrimSpace(values[0])

	index := 0
	if strings.HasPrefix(value, "<%= ENV['") && strings.HasSuffix(value, "'] %>") {
		env := strings.TrimSuffix(strings.TrimPrefix(value, "<%= ENV['"), "'] %>")
		index = strings.Index(line, "<%= ENV['")
		value = os.Getenv(env)
		if value == "" {
			value = strings.TrimSpace(values[1])
		}
	} else if strings.HasPrefix(value, "<%= ENV_INT['") && strings.HasSuffix(value, "'] %>") {
		env := strings.TrimSuffix(strings.TrimPrefix(value, "<%= ENV_INT['"), "'] %>")
		index = strings.Index(line, "<%= ENV_INT['")
		fmt.Println(env)
		value = os.Getenv(env)
		if value == "" {
			value = strings.TrimSpace(values[1])
		}
	} else if strings.HasPrefix(value, "<%= ENV_BOOL['") && strings.HasSuffix(value, "'] %>") {
		env := strings.TrimSuffix(strings.TrimPrefix(value, "<%= ENV_BOOL['"), "'] %>")
		index = strings.Index(line, "<%= ENV_BOOL['")
		fmt.Println(env)
		value = os.Getenv(env)
		if value == "" {
			value = strings.TrimSpace(values[1])
		}
	} else if strings.HasPrefix(value, "<%= ENV_LIST['") && strings.HasSuffix(value, "'] %>") {
		env := strings.TrimSuffix(strings.TrimPrefix(value, "<%= ENV_LIST['"), "'] %>")
		index = strings.Index(line, "<%= ENV_LIST['")
		fmt.Println(env)
		value = os.Getenv(env)
		if value == "" {
			value = strings.TrimSpace(values[1])
		}
	} else {
		value = strings.TrimSpace(values[1])
	}

	newline := strings.Replace(line,
		line[index:], value, 1)
	return newline
}
