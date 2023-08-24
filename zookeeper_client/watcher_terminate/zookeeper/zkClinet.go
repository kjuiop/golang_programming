package zookeeper

import (
	"github.com/curator-go/curator"
	"github.com/samuel/go-zookeeper/zk"
	"golang_programming/zookeeper_client/watcher_terminate/config"
	"log"
	"time"
)

type ZKCon struct {
	curator   curator.CuratorFramework
	connected bool
	cfg       *config.Config
}

func NewZkClient(cfg *config.Config) (*ZKCon, error) {

	zkCon := &ZKCon{
		curator:   nil,
		connected: false,
		cfg:       cfg,
	}

	// 접속 timeout 정책 설정
	retryPolicy := curator.NewExponentialBackoffRetry(time.Second, 3, 15*time.Second)

	// 클라이언트를 생성한다.
	zkCon.curator = curator.NewClient(cfg.ZookeeperInfo.Host, retryPolicy)
	if err := zkCon.curator.Start(); err != nil {
		return nil, err
	}

	// 연결이 완료될 때까지 대기한다.
	if err := zkCon.curator.ZookeeperClient().BlockUntilConnectedOrTimedOut(); err != nil {
		return nil, err
	}

	return zkCon, nil
}

func (zkCon *ZKCon) WatcherChildrenNodeToMap(path string, fn func(*zk.Event)) ([]string, error) {
	val, err := zkCon.curator.GetChildren().UsingWatcher(curator.NewWatcher(fn)).ForPath(path)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (zkCon *ZKCon) CheckExists(path string) error {
	stat, err := zkCon.curator.CheckExists().ForPath(path)
	if err != nil {
		return err
	}

	if stat == nil {
		_, err := zkCon.curator.Create().WithMode(curator.PERSISTENT).ForPath(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (zkCon *ZKCon) EndZookeeper() {
	if zkCon.curator != nil {
		if err := zkCon.curator.Close(); err != nil {
			log.Printf("curator close error : %v\n", err)
		}
	}
	zkCon.connected = false
	log.Println("zookeeper close")
}
