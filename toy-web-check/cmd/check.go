package cmd

import (
	"fmt"
	"time"

	"github.com/Keepinmind100/golang-sandbox/toy-web-check/internal/checker"
	"github.com/spf13/cobra"
)

var timeout time.Duration

var checkCmd = &cobra.Command{
	Use:   "check [urls]",
	Short: "입력한 URL들을 병렬로 상태 점검합니다",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results := checker.CheckURLs(args, timeout)
		for _, res := range results {
			status := "✅"
			if res.Err != nil || res.StatusCode >= 400 {
				status = "❌"
			}
			fmt.Printf("%s [%d] %s\n", status, res.StatusCode, res.URL)
		}
	},
}

func init() {
	checkCmd.Flags().DurationVarP(&timeout, "timeout", "t", 3*time.Second, "HTTP 요청 timeout")
	rootCmd.AddCommand(checkCmd)
}
