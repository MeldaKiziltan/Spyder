package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"

	"context"

	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/motor"
	"go.viam.com/rdk/robot"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/rdk/utils"
	"go.viam.com/utils/rpc"
)

var actions chan string

var robo robot.Robot

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Watchdog struct {
	interval time.Duration
	timer    *time.Timer
}

func NewWatchDog(interval time.Duration, callback func()) *Watchdog {
	w := Watchdog{
		interval: interval,
		timer:    time.AfterFunc(interval, callback),
	}
	return &w
}

func (w *Watchdog) Stop() {
	w.timer.Stop()
}

func (w *Watchdog) Kick() {
	w.timer.Stop()
	w.timer.Reset(w.interval)
}

var wdog *Watchdog

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {

	// drive
	drive, err := motor.FromRobot(robo, "drive")
	if err != nil {
		return
	}
	// steering
	steering, err := motor.FromRobot(robo, "steering")
	if err != nil {
		return
	}

	wdog = NewWatchDog(1*time.Second, func() {
		fmt.Println("watchdog stop all motors")
		drive.Stop(context.Background(), nil)
		steering.Stop(context.Background(), nil)
		wdog.Kick()
	})
	defer wdog.Stop()
	defer drive.Stop(context.Background(), nil)
	defer steering.Stop(context.Background(), nil)

	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		msg := strings.Trim(string(p), "\n")
		// print out that message for clarity
		log.Println("got websocket msg:", msg)
		wdog.Kick()

		params := strings.Split(msg, ",")
		if len(params) != 2 {
			log.Println("Incorrect params")
			return
		}

		if params[1] == "stop" {
			if params[0] == "drive" {
				drive.Stop(context.Background(), nil)
			} else {
				steering.Stop(context.Background(), nil)
			}
		} else {
			speed, err := strconv.ParseFloat(strings.TrimSpace(params[1]), 32)
			if err != nil {
				log.Println("incorrect params, cant cast")
				return
			}
			if params[0] == "drive" {
				drive.SetPower(context.Background(), speed, nil)
			} else {
				steering.SetPower(context.Background(), speed, nil)
			}
		}

		fmt.Println("websocket")

		if err := conn.WriteMessage(messageType, []byte("ACK")); err != nil {
			log.Println(err)
			return
		}

	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", ping)
	http.HandleFunc("/ws", wsEndpoint)
}

func setupRobot() {
	logger := golog.NewDevelopmentLogger("client")
	var err error
	robo, err = client.New(
		context.Background(),
		"panda-main.xze1jqek1n.viam.cloud",
		logger,
		client.WithDialOptions(rpc.WithCredentials(rpc.Credentials{
			Type:    utils.CredentialsTypeRobotLocationSecret,
			Payload: "s6vmwz49ljl9eweoeb7qnyla63bqkqskt1dsbf8m5ymg6783",
		})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Resources:")
	logger.Info(robo.ResourceNames())

	// Note that the pin supplied is a placeholder. Please change this to a valid pin.
	// pi
	piComponent, err := board.FromRobot(robo, "pi")
	if err != nil {
		logger.Error(err)
		return
	}
	piReturnValue, err := piComponent.GPIOPinByName("16")
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("pi GPIOPinByName return value: %+v", piReturnValue)

	// drive
	drive, err := motor.FromRobot(robo, "drive")
	if err != nil {
		logger.Error(err)
		return
	}
	driveReturnValue, err := drive.IsMoving(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("drive IsMoving return value: %+v", driveReturnValue)

	// steering
	steering, err := motor.FromRobot(robo, "steering")
	if err != nil {
		logger.Error(err)
		return
	}

	steeringReturnValue, err := steering.IsMoving(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("steering IsMoving return value: %+v", steeringReturnValue)

}

func main() {
	fmt.Println("Hello World")
	setupRobot()
	fmt.Println("robot set up")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
	robo.Close(context.Background())
}
