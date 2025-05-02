webcheck/
├── cmd/
│   ├── root.go          # 기본 root 명령
│   ├── check.go         # 'check' 명령: 여러 URL 병렬 상태 점검
│   ├── watch.go         # 'watch' 명령: 일정 간격으로 URL 상태 모니터링
│
├── internal/
│   ├── checker/
│   │   ├── checker.go   # 핵심 로직: goroutine + http check + 채널
│   ├── output/
│   │   ├── printer.go   # 결과 출력 관련 (컬러, json, 스피너 등)
│
├── urls/
│   ├── reader.go        # URL 파일 읽기 처리 (옵션에서 활용)
│
├── go.mod
├── main.go              # cobra root 실행

----
cmd/	            CLI 명령들 정의 (check, watch 등)
internal/checker/	각 URL 상태 체크를 고루틴으로 처리
internal/output/	CLI 출력: 컬러 텍스트, json 포맷 등
urls/reader.go	    URL을 파일에서 읽어오는 유틸
main.go	            cobra 초기화 및 실행

---

go get github.com/spf13/cobra@latest

go get github.com/spf13/viper@latest
-> 설정관리 ( 환경변수  , Config파일 , CLI 플래그 등)을 쉽게 해주는 라이브러리 