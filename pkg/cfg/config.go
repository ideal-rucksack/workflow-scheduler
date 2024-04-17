package cfg

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/constants"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func SetupBootstrap(models ...string) {
	if models != nil {
		for i := range models {
			if models[i] != "" {
				// 设置
				dir, err := os.Getwd()
				if err != nil {
					log.Println("获取当前目录失败", err)
				}
				err = os.Setenv(constants.HOME, dir)
				if err != nil {
					log.Println(err.Error(), true, err)
				}
				err = godotenv.Load(os.Getenv(constants.HOME) + "/conf/" + models[i] + ".env")
				if err != nil {
					log.Printf("Error loading .env file: %v\n", err.Error())
				}
			}
		}
	} else {
		log.Fatalln("models is empty, please check it.")
	}
}
