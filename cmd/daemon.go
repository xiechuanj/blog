package cmd

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/macaron.v1"

	"github.com/xiechuanj/blog/models"
	"github.com/xiechuanj/blog/module"
	"github.com/xiechuanj/blog/utils"
	"github.com/xiechuanj/blog/web"
)

var address string
var port int64

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Web subcommand start/stop/monitor Blog's REST API daemon.",
	Long:  ``,
}

var startDaemonCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Blog's REST API daemon.",
	Long:  ``,
	Run:   startDeamon,
}

var stopDaemonCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop Blog's REST API daemon.",
	Long:  ``,
	Run:   stopDaemon,
}

var monitorDeamonCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor Blog's REST API daemon.",
	Long:  ``,
	Run:   monitorDaemon,
}

func init() {
	RootCmd.AddCommand(daemonCmd)

	daemonCmd.AddCommand(startDaemonCmd)
	startDaemonCmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "http or https listen address.")
	startDaemonCmd.Flags().Int64VarP(&port, "port", "p", 80, "the port of http.")

	daemonCmd.AddCommand(stopDaemonCmd)
	daemonCmd.AddCommand(monitorDeamonCmd)
}

func startDeamon(cmd *cobra.Command, args []string) {
	log.SetOutput(os.Stdout)

	models.DB.Open()

	module.InitTimerTask()

	m := macaron.New()

	web.SetBlogdMacaron(m)

	listenMode := configure.GetString("listenmode")
	switch listenMode {
	case "http":
		listenaddr := fmt.Sprintf("%s:%d", address, port)
		fmt.Println("Blog is listen:", listenaddr)
		if err := http.ListenAndServe(listenaddr, m); err != nil {
			fmt.Printf("Start Blog http service error: %v\n", err.Error())
		}
		break
	case "https":
		listenaddr := fmt.Sprintf("%s:443", address)
		server := &http.Server{Addr: listenaddr, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS10}, Handler: m}
		if err := server.ListenAndServeTLS(configure.GetString("httpscertfile"), configure.GetString("httpskeyfile")); err != nil {
			fmt.Printf("Start Blog https service error: %v\n", err.Error())
		}
		break
	case "unix":
		listenaddr := fmt.Sprintf("%s", address)
		if utils.IsFileExist(listenaddr) {
			os.Remove(listenaddr)
		}

		if listener, err := net.Listen("unix", listenaddr); err != nil {
			fmt.Printf("Start Blog unix socket error: %v\n", err.Error())

		} else {
			server := &http.Server{Handler: m}
			if err := server.Serve(listener); err != nil {
				fmt.Printf("Start Blog unix socket error: %v\n", err.Error())
			}
		}
		break
	default:
		break
	}
}

func stopDaemon(cmd *cobra.Command, args []string) {

}

func monitorDaemon(cmd *cobra.Command, args []string) {

}
