package chs

import (
	"context"
	"fmt"
	"gitlab.arksec.cn/wpf1118/toolbox/tools/db"
	"gitlab.arksec.cn/wpf1118/toolbox/tools/flag"
	"gitlab.arksec.cn/wpf1118/toolbox/tools/help"
	"gitlab.arksec.cn/wpf1118/toolbox/tools/logging"
	"testing"
	"time"
)

func TestNewChs(t *testing.T) {
	redisOpts := flag.NewDefaultRedisOpts()
	redisOpts.Endpoint = "www.zzrs.xyz:26379"
	db.RedisInit(redisOpts)

	redis := db.NewRedis()
	ctx := context.Background()

	// 意图：启动10个协程，处理任务
	// 每个任务，输出一段内容
	fn := func(i interface{}) (interface{}, error) {
		// 模拟耗时
		time.Sleep(time.Millisecond * 100)
		if v, ok := i.(int); ok {
			key := fmt.Sprintf("%d", v)
			value := fmt.Sprintf("%d", v)
			_, err := redis.Set(ctx, key, value)
			if err != nil {
				logging.ErrorF("%v", err)
				return nil, err
			}

			logging.InfoF("content: %d %s", v, help.RandStrForNow())
			return nil, nil
		}

		return nil, nil
	}

	chs := NewChs("测试任务", 100, fn)

	var tasks []interface{}
	for i := 0; i < 10001; i++ {
		tasks = append(tasks, i)
	}

	chs.Work(tasks)

	total := chs.GetTotal()

	fmt.Println(total)
}
