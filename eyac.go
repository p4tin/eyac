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

func getEnvOrDefault(line string) string {
	parts := strings.Split(line, ":")
	if len(parts) < 2 || parts[1] == "" {
		return line
	}
	values := strings.Split(strings.TrimSpace(parts[1]), ",")
	value := strings.TrimSpace(values[0])
	if strings.HasPrefix(value, "<%= ENV['") && strings.HasSuffix(value, "'] %>") {
		env := strings.TrimSuffix(strings.TrimPrefix(value, "<%= ENV['"), "'] %>")
		value = os.Getenv(env)
		if value == "" {
			value = strings.TrimSpace(values[1])
		}
	} else {
		value = strings.TrimSpace(values[1])
	}
	index := strings.Index(line, "<%= ENV['")

	newline := strings.Replace(line,
		line[index:], value, 1)
	return newline
}
