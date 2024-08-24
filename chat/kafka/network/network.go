package network

import (
	"chat-kafka/config"
	"chat-kafka/repository"
	"chat-kafka/service"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg        *config.Config
	engine     *gin.Engine
	service    *service.Service
	repository *repository.Repository
	ip         string
	port       string
}

type ServerInfoEvent struct {
	IP     string
	Status bool
}

func NewServer(cfg *config.Config, service *service.Service, repository *repository.Repository) *Server {

	s := &Server{
		cfg:        cfg,
		engine:     gin.New(),
		service:    service,
		repository: repository,
		port:       cfg.Server.Port,
	}

	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())
	s.engine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	registerServer(s)

	return s
}

func (n *Server) setServerInfo() {
	// IP를 가져오고,
	// IP를 기반으로 MySQL serverInfo 테이블 데이터 추가

	if addrs, err := net.InterfaceAddrs(); err != nil {
		log.Fatalln(err.Error())
	} else {
		var ip net.IP

		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					ip = ipNet.IP
					break
				}
			}
		}

		if ip == nil {
			log.Fatalln("no ip address found")
		}

		if err = n.service.ServerSet(fmt.Sprintf("%s:%s", ip.String(), n.port), true); err != nil {
			log.Fatalln(err.Error())
		} else {
			n.ip = ip.String()
		}
	}
}

func (n *Server) StartServer() error {
	n.setServerInfo()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		<-sigCh

		addr := fmt.Sprintf("%s:%s", n.ip, n.port)
		if err := n.service.ServerSet(addr, false); err != nil {
			// 실패케이스 추가 처리 필요
			log.Println("Failed to set server info when close", "err", err.Error())
		}

		// kafka 이벤트 전송
		e := &ServerInfoEvent{IP: addr, Status: false}
		ch := make(chan kafka.Event)

		if v, err := json.Marshal(e); err != nil {
			log.Println("Failed to Marshal")
		} else if result, err := n.service.PublishEvent("chat", v, ch); err != nil {
			log.Println("Failed To Send Event To Kafka", "err", err)
		} else {
			// Send Event To Kafka
			log.Println("Success To Send Event", result)
		}
		os.Exit(1)
	}()

	return n.engine.Run(fmt.Sprintf(":%s", n.cfg.Server.Port))
}
