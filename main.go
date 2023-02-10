package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-resty/resty/v2"
)

func main() {
	url := "http://localhost:3000/api/web-console/v1/ondemand/jobs"
	cookie := "_ga=GA1.1.484072228.1662045000; _biz_uid=f06e970d51f44231a03f8e64a50e7f2b; _biz_flagsA=%7B%22Version%22%3A1%2C%22ViewThrough%22%3A%221%22%2C%22XDomain%22%3A%221%22%7D; hubspotutk=cbc533be50f7620d527227e2034176db; apt.uid=AP-S8VLIJURGKY0-2-4-1662045004869-98433362.0.2.965c0cd2-500b-424e-a457-f8b6c21d6e35; __hstc=181257784.cbc533be50f7620d527227e2034176db.1662045002166.1667244546928.1668181648383.17; _biz_nA=108; _biz_pendingA=%5B%22m%2Fipv%3F_biz_r%3D%26_biz_h%3D802059049%26_biz_u%3Df06e970d51f44231a03f8e64a50e7f2b%26_biz_s%3D5ac105%26_biz_l%3Dhttp%253A%252F%252Flocalhost%253A3000%252Fresults%252Fgithub.com%252Fobarra-dev%252Fminesweeper-API%252F01GK2H4MYYV5J06DQJT09X73X9%252F%253Fph%253DPhaseIntroduced%26_biz_t%3D1669753816188%26_biz_i%3DSonatype%2520Lift%2520--%2520Console%26_biz_n%3D107%26rnd%3D871882%22%5D; muserAuth=2aVWSduSQ1uwGbh8hAUGZUi96HjosSrvhwaxioSRwWSgS1FFte1ZCGUvuMw9BfphrR8hWQAb9c7cwS7BNh8FSbZcLfk9tJZm5HXGt3dzbWpmqrbV3hZF9JzZ5BRT9at3FNytjPcpKA8eZ4jxG5JQE55QA3T4yc63KAWR2vSNiqB1EUvRSP; muserSiteCookie=Ks5v1yoE6Xv1jRZKWk9TyYgT4YBZJAJfWMsbyyKqM98sC55j3TtF6gXuWayYQTnfppXUWgBwiMuMYFASpY4t4bTWJSthwyr5MgRUXTHF"
	repos := getRepos()

	fmt.Println("Count repos: ", len(repos))
	var wg sync.WaitGroup
	client := resty.New()

	for _, repo := range repos {
		wg.Add(1)

		go func(repo string) {
			defer wg.Done()
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("cookie", cookie).
				SetBody(repo).
				Post(url)

			failOnError(err, "Error doing POST")
			fmt.Println("Resp: ", resp, "for: ", repo)
		}(repo)

	}

	wg.Wait()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s, %s", msg, err, err.Error())
	}
}
