package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func InitViper() error {
	viper.SetEnvPrefix("FATECLOUD")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	var altPath = os.Getenv("cUserSpecifiedPATH")
	if altPath != "" {
		if !DirExists(altPath) {
			return fmt.Errorf("%s %s does not exist", cUserSpecifiedPATH, altPath)
		}
	} else {
		// Append the project dir to the config seraching path
		path, _ := filepath.Abs(".")
		viper.AddConfigPath(path)
	}
	return nil
}

func main() {
	fmt.Println("test")
	strs := []string{"1", "2"}
	sli := append([]string(nil), strs...)
	fmt.Println(sli)
	fmt.Printf("%T\n", sli)

	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}
