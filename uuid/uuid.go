package uuid

import(
	"strconv"
	idworker "github.com/gitstliu/go-id-worker"
)

func GenerateNextUuid() string{
	currWoker := &idworker.IdWorker{} 
	currWoker.InitIdWorker(1000, 1)
	newId,_ := currWoker.NextId()
	return strconv.FormatInt(newId,10)
}
