# pprof 사용방법

---

1. /pprof 에서 pprof 서버를 띄운다.
2. /measurement_goroutine 에서 `go build` 를 진행 후 trace.out 파일을 생성한다.
3. /measurement_goroutine 에서 `curl http://0.0.0.0:6060/debug/pprof/trace\?seconds\=30 --output trace.out` 을 통해 분석 report 를 전송한다.
4. `go tool trace -http 0.0.0.0:9090 trace.out` 을 통해 trace.out 을 분석한 대시보드를 띄운다.
