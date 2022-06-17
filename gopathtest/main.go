package gopathtest

import (
	_ "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello, go!")
	logrus.Println(uuid.NewString())
}
