package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Query().Get("name")

		// 漏洞：直接使用用户输入的文件路径，存在路径遍历风险
		data, err := ioutil.ReadFile("/var/www/" + fileName)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(data)

				// 漏洞：直接使用用户输入的文件路径，存在路径遍历风险
		data, err = ioutil.ReadFile("/var/www/" + fileName)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(data)
		
	})

	abs(2)
	http.ListenAndServe(":8080", nil)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x
	}
}

func sanitizeUrl(redir string) string {
	if len(redir) > 0 && redir[0] == '/' {
		return redir
	}
	return "/"
}
// package main

// import (
// 	"flag"
// 	"io/ioutil"

// 	"gopkg.in/yaml.v2"

// 	"show-live/config"
// 	"show-live/internal/simullink"
// 	"show-live/pkg/db"
// 	"show-live/pkg/email"
// 	"show-live/pkg/log"
// )

// func main() {
// 	var config config.Simullink
// 	configFilePath := flag.String("config", "config-s.yml", "config file")
// 	if configFilePath != nil {
// 		configFile, err := ioutil.ReadFile(*configFilePath)
// 		if err != nil {
// 			log.Logger.Fatal(err)
// 		}
// 		err = yaml.Unmarshal(configFile, &config)
// 		if err != nil {
// 			log.Logger.Fatal(err)
// 		}
// 	}
// 	log.InitLogger(config.Log.LogSuffix, config.Log.LogDir)
// 	d, err := db.InitCache(config.DBDir)
// 	if err != nil {
// 		log.Logger.Errorf("init cache error %v", err)
// 		return
// 	}
// 	defer func() {
// 		if err := d.Exit(); err != nil {
// 			log.Logger.Errorf("db exits error %v", err)
// 		}
// 	}()
// 	e := email.NewEmailSender(config.Email)
// 	c := simullink.NewSimullinkGetter(d, config.TagsSelected, config.URL, config.CityCode)
// 	events, err := c.GetEventsToNotify()
// 	if err != nil {
// 		log.Logger.Errorf("get events to notify error %v", err)
// 		if err := e.Send("同感获取最新演出出错了", err.Error()); err != nil {
// 			log.Logger.Errorf("send email error %v", err)
// 		}
// 		return
// 	}
// 	if len(events) == 0 {
// 		log.Logger.Info("no new event to send")
// 		return
// 	}
// 	var content string
// 	for _, v := range events {
// 		content += v + "\n"
// 	}
// 	if err := e.Send("同感演出上新了", content); err != nil {
// 		log.Logger.Errorf("send email error %v", err)
// 	}
// 	log.Logger.Infof("%d event sent", len(events))
// }
