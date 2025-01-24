package main
import (
	"github.com/sirupsen/logrus"
)
func main(){
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("This is our second entry in our log file.")
}