package loader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Config struct {
	Port     int    `json:"port"`
	LogLevel string `json:"logLevel"`
}

func FileToSlice(filePath string) ([]int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %w", err)
	}

	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse number: %w", err)
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error while reading file: %w", err)
	}

	return numbers, nil
}

func LoadConfig(configPath string) (Config, error) {
	file, err := os.Open(configPath)

	if err != nil {
		return Config{}, nil
	}

	byteValue, err := io.ReadAll(file)

	if err != nil {
		return Config{}, nil
	}

	var config Config

	err = json.Unmarshal(byteValue, &config)

	if err != nil {
		return Config{}, nil
	}

	return config, nil
}
