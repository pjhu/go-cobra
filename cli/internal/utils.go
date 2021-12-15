package internal

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)
func CheckArgsAmount(cmd *cobra.Command, args []string, amount int)  {
	if len(args) != amount {
		cmd.HelpFunc()(cmd, args)
		fmt.Printf("\nError: accepts %d arg(s), received %d\n", amount, len(args))
		os.Exit(3)
	}
}

func ExecShellCommand(command string, args ...string) string {
	rst, err := ExecShellCommandNotExit(command, args...)
	if err != nil {
		fmt.Println(rst + "\n" + err.Error())
		os.Exit(2)
	}
	fmt.Println(rst)
	return rst
}

func ExecShellCommandNotExit(command string, args ...string) (string, error ) {
	localeCmd := exec.Command(command, args...)
	fmt.Printf("COMMAND: %s, ARGS: %s\n", command, args)
	var stdout, stderr bytes.Buffer
	localeCmd.Stdout = &stdout // 标准输出
	localeCmd.Stderr = &stderr // 标准错误
	err := localeCmd.Run()
	stdOutStr, stdErrStr := string(stdout.Bytes()), string(stderr.Bytes())

	rst := strings.Trim(stdOutStr+"\n"+stdErrStr, "")
	return rst, err
}

func RestClient() *resty.Request {
	client := resty.New()
	client.SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(1 * time.Second).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")
	return client.R().EnableTrace()
}

func RestResponseInfo(err error, resp *resty.Response) string {
	fmt.Println("REQUEST INFO: \n", resp.Request.RawRequest)
	var builder strings.Builder
	if err != nil {
		builder.WriteString("  Error      :")
		builder.WriteString(err.Error())
	}
	builder.WriteString("RESPONSE INFO:\n")
	builder.WriteString("  Status Code:")
	builder.WriteString(strconv.Itoa(resp.StatusCode()))
	builder.WriteString("  Status     :")
	builder.WriteString(resp.Status())
	builder.WriteString("\n")
	builder.WriteString("  Proto      :")
	builder.WriteString(resp.Proto())
	builder.WriteString("\n")
	builder.WriteString("  Time       :")
	builder.WriteString(resp.Time().String())
	builder.WriteString("\n")
	builder.WriteString("  Received At:")
	builder.WriteString(resp.ReceivedAt().String())
	builder.WriteString("\n")
	builder.WriteString("  Body       :")
	builder.WriteString(resp.String())
	builder.WriteString("\n")

	// Explore trace info
	builder.WriteString("REQUEST TRACE INFO:\n")
	ti := resp.Request.TraceInfo()
	builder.WriteString("  RemoteAddr    :")
	builder.WriteString(ti.RemoteAddr.String())
	builder.WriteString("\n")
	rst := strings.Trim(builder.String(), "")
	fmt.Printf(rst)
	return rst
}