package monitor

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{}
var homeHTML string
var gaugeJS string
var loadCPU []float64

func sumValues(values []string) int {
	sum := 0
	for _, v := range values {
		// TODO check error cases
		i, _ := strconv.Atoi(v)
		sum += i
	}
	return sum
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", homeHTML)
}

func gaugejs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", gaugeJS)
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		for i, v := range loadCPU {
			message := fmt.Sprintf("cpu%v|%v", i, int(v*100))
			err = c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func httpServe() {
	homeHTMLb, err := ioutil.ReadFile("home.html")
	if err != nil {
		log.Fatal(err)
	}

	homeHTML = string(homeHTMLb)

	gaugeJSb, err := ioutil.ReadFile("js/svg-gauge-master/dist/gauge.js")
	if err != nil {
		log.Fatal(err)
	}

	gaugeJS = string(gaugeJSb)

	http.HandleFunc("/", home)
	http.HandleFunc("/ws", ws)
	http.HandleFunc("/gauge.js", gaugejs)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func MonitorCpu() {
	go httpServe()

	ptotal, pidle := monitorCpu()
	loadCPU = make([]float64, len(ptotal))

	var ctotal, cidle []int
	for {
		time.Sleep(1 * time.Second)
		// fmt.Println("\033[H\033[2J")
		ctotal, cidle = monitorCpu()
		for i, _ := range ctotal {
			id := 1 - float64(cidle[i]-pidle[i])/float64(ctotal[i]-ptotal[i])
			// fmt.Println(id * 100)
			loadCPU[i] = id
		}
		ptotal, pidle = ctotal, cidle
	}
}

func monitorCpu() (total, idle []int) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")

		if fields[0] == "cpu" {
			continue
		}

		if strings.HasPrefix(fields[0], "cpu") {
			total = append(total, sumValues(fields[1:]))
			i, _ := strconv.Atoi(fields[4])
			idle = append(idle, i)

		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
