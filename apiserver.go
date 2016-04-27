package localkube

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	apiserver "k8s.io/kubernetes/cmd/kube-apiserver/app"
	"k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	etcdstorage "k8s.io/kubernetes/pkg/storage/etcd"
)

const (
	APIServerName       = "apiserver"
	APIServerHost       = "127.0.0.1"
	APIServerSecureHost = "0.0.0.0"
	APIServerPort       = 8080
	APIServerSecurePort = 443
)

var (
	APIServerURL   string
	ServiceIPRange = "10.1.30.0/24"
	APIServerStop  chan struct{}
)

func init() {
	APIServerURL = fmt.Sprintf("http://%s:%d", APIServerHost, APIServerPort)
	if ipRange := os.Getenv("SERVICE_IP_RANGE"); len(ipRange) != 0 {
		ServiceIPRange = ipRange
	}
}

func NewAPIServer() Server {
	return &SimpleServer{
		ComponentName: APIServerName,
		StartupFn:     StartAPIServer,
		ShutdownFn: func() {
			close(APIServerStop)
		},
	}
}

func StartAPIServer() {
	APIServerStop = make(chan struct{})
	config := options.NewAPIServer()

	// use host/port from vars
	config.BindAddress = net.ParseIP(APIServerSecureHost)
	config.InsecureBindAddress = net.ParseIP(APIServerHost)
	config.InsecurePort = APIServerPort
	config.SecurePort = APIServerSecurePort
	config.ClientCAFile = "/root/easy-rsa-master/easyrsa3/pki/ca.crt"
	config.TLSCertFile = "/root/easy-rsa-master/easyrsa3/pki/issued/kubernetes-master.crt"
	config.TLSPrivateKeyFile = "/root/easy-rsa-master/easyrsa3/pki/private/kubernetes-master.key"
	config.AdmissionControl = "NamespaceLifecycle,LimitRanger,SecurityContextDeny,ServiceAccount,ResourceQuota"

	// use localkube etcd
	config.EtcdConfig = etcdstorage.EtcdConfig{
		ServerList: KubeEtcdClientURLs,
	}

	// set Service IP range
	_, ipnet, err := net.ParseCIDR(ServiceIPRange)
	if err != nil {
		panic(err)
	}
	config.ServiceClusterIPRange = *ipnet

	// defaults from apiserver command
	config.EnableProfiling = true
	config.EnableWatchCache = true
	config.MinRequestTimeout = 1800

	fn := func() error {
		return apiserver.Run(config)
	}

	// start API server in it's own goroutine
	go until(fn, os.Stdout, APIServerName, 200*time.Millisecond, SchedulerStop)
}

// notFoundErr returns true if the passed error is an API server object not found error
func notFoundErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasSuffix(err.Error(), "not found")
}
