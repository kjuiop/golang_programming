package main

import "fmt"

var myConfig Config

type Config struct {
	BackupDir string `envconfig:"TR_BACKUP_DIR"`
}

func main() {

	content := map[string]interface{}{
		"origin": map[string]interface{}{},
	}

	s := content["origin"].(map[string]interface{})["input_filename"].(string)
	fmt.Println(s)

}
