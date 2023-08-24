package zookeeper

import (
	"github.com/curator-go/curator"
	"golang_programming/zookeeper_client/watcher_terminate/config"
	"time"
)

type ZKCon struct {
	zkClient  curator.ZookeeperConnection
	curator   curator.CuratorFramework
	connected bool
	cfg       *config.Config
}

func NewZkClient(cfg *config.Config) (*ZKCon, error) {

	zkClient := &ZKCon{
		curator:   nil,
		connected: false,
		cfg:       cfg,
	}

	// 접속 timeout 정책 설정
	retryPolicy := curator.NewExponentialBackoffRetry(time.Second, 3, 15*time.Second)

	// 클라이언트를 생성한다.
	zkClient.curator = curator.NewClient(cfg.ZookeeperInfo.Host, retryPolicy)
	if err := zkClient.curator.Start(); err != nil {
		return nil, err
	}

	// 연결이 완료될 때까지 대기한다.
	if err := zkClient.curator.ZookeeperClient().BlockUntilConnectedOrTimedOut(); err != nil {
		return nil, err
	}

	return zkClient, nil
}
