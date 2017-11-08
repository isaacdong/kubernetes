/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClientConnectionConfiguration contains details for constructing a client.
type ClientConnectionConfiguration struct {
	// kubeConfigFile is the path to a kubeconfig file.
	KubeConfigFile string `json:"kubeconfig"`
	// acceptContentTypes defines the Accept header sent by clients when connecting to a server, overriding the
	// default value of 'application/json'. This field will control all connections to the server used by a particular
	// client.
	AcceptContentTypes string `json:"acceptContentTypes"`
	// contentType is the content type used when sending data to the server from this client.
	ContentType string `json:"contentType"`
	// cps controls the number of queries per second allowed for this connection.
	QPS float32 `json:"qps"`
	// burst allows extra queries to accumulate when a client is exceeding its rate.
	Burst int `json:"burst"`
}

// KubeProxyIPTablesConfiguration contains iptables-related configuration
// details for the Kubernetes proxy server.
type KubeProxyIPTablesConfiguration struct {
	// masqueradeBit is the bit of the iptables fwmark space to use for SNAT if using
	// the pure iptables proxy mode. Values must be within the range [0, 31].
	MasqueradeBit *int32 `json:"masqueradeBit"`
	// masqueradeAll tells kube-proxy to SNAT everything if using the pure iptables proxy mode.
	MasqueradeAll bool `json:"masqueradeAll"`
	// syncPeriod is the period that iptables rules are refreshed (e.g. '5s', '1m',
	// '2h22m').  Must be greater than 0.
	SyncPeriod metav1.Duration `json:"syncPeriod"`
	// minSyncPeriod is the minimum period that iptables rules are refreshed (e.g. '5s', '1m',
	// '2h22m').
	MinSyncPeriod metav1.Duration `json:"minSyncPeriod"`
}

// KubeProxyIPVSConfiguration contains ipvs-related configuration
// details for the Kubernetes proxy server.
type KubeProxyIPVSConfiguration struct {
	// syncPeriod is the period that ipvs rules are refreshed (e.g. '5s', '1m',
	// '2h22m').  Must be greater than 0.
	SyncPeriod metav1.Duration `json:"syncPeriod"`
	// minSyncPeriod is the minimum period that ipvs rules are refreshed (e.g. '5s', '1m',
	// '2h22m').
	MinSyncPeriod metav1.Duration `json:"minSyncPeriod"`
	// ipvs scheduler
	Scheduler string `json:"scheduler"`
}

// KubeProxyConntrackConfiguration contains conntrack settings for
// the Kubernetes proxy server.
type KubeProxyConntrackConfiguration struct {
	// max is the maximum number of NAT connections to track (0 to
	// leave as-is).  This takes precedence over conntrackMaxPerCore and conntrackMin.
	Max int32 `json:"max"`
	// maxPerCore is the maximum number of NAT connections to track
	// per CPU core (0 to leave the limit as-is and ignore conntrackMin).
	MaxPerCore int32 `json:"maxPerCore"`
	// min is the minimum value of connect-tracking records to allocate,
	// regardless of conntrackMaxPerCore (set conntrackMaxPerCore=0 to leave the limit as-is).
	Min int32 `json:"min"`
	// tcpEstablishedTimeout is how long an idle TCP connection will be kept open
	// (e.g. '2s').  Must be greater than 0.
	TCPEstablishedTimeout metav1.Duration `json:"tcpEstablishedTimeout"`
	// tcpCloseWaitTimeout is how long an idle conntrack entry
	// in CLOSE_WAIT state will remain in the conntrack
	// table. (e.g. '60s'). Must be greater than 0 to set.
	TCPCloseWaitTimeout metav1.Duration `json:"tcpCloseWaitTimeout"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeProxyConfiguration contains everything necessary to configure the
// Kubernetes proxy server.
type KubeProxyConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// featureGates is a comma-separated list of key=value pairs that control
	// which alpha/beta features are enabled.
	//
	// TODO this really should be a map but that requires refactoring all
	// components to use config files because local-up-cluster.sh only supports
	// the --feature-gates flag right now, which is comma-separated key=value
	// pairs.
	FeatureGates string `json:"featureGates"`

	// bindAddress is the IP address for the proxy server to serve on (set to 0.0.0.0
	// for all interfaces)
	BindAddress string `json:"bindAddress"`
	// healthzBindAddress is the IP address and port for the health check server to serve on,
	// defaulting to 0.0.0.0:10256
	HealthzBindAddress string `json:"healthzBindAddress"`
	// metricsBindAddress is the IP address and port for the metrics server to serve on,
	// defaulting to 127.0.0.1:10249 (set to 0.0.0.0 for all interfaces)
	MetricsBindAddress string `json:"metricsBindAddress"`
	// enableProfiling enables profiling via web interface on /debug/pprof handler.
	// Profiling handlers will be handled by metrics server.
	EnableProfiling bool `json:"enableProfiling"`
	// clusterCIDR is the CIDR range of the pods in the cluster. It is used to
	// bridge traffic coming from outside of the cluster. If not provided,
	// no off-cluster bridging will be performed.
	ClusterCIDR string `json:"clusterCIDR"`
	// hostnameOverride, if non-empty, will be used as the identity instead of the actual hostname.
	HostnameOverride string `json:"hostnameOverride"`
	// clientConnection specifies the kubeconfig file and client connection settings for the proxy
	// server to use when communicating with the apiserver.
	ClientConnection ClientConnectionConfiguration `json:"clientConnection"`
	// iptables contains iptables-related configuration options.
	IPTables KubeProxyIPTablesConfiguration `json:"iptables"`
	// ipvs contains ipvs-related configuration options.
	IPVS KubeProxyIPVSConfiguration `json:"ipvs"`
	// oomScoreAdj is the oom-score-adj value for kube-proxy process. Values must be within
	// the range [-1000, 1000]
	OOMScoreAdj *int32 `json:"oomScoreAdj"`
	// mode specifies which proxy mode to use.
	Mode ProxyMode `json:"mode"`
	// portRange is the range of host ports (beginPort-endPort, inclusive) that may be consumed
	// in order to proxy service traffic. If unspecified (0-0) then ports will be randomly chosen.
	PortRange string `json:"portRange"`
	// resourceContainer is the bsolute name of the resource-only container to create and run
	// the Kube-proxy in (Default: /kube-proxy).
	ResourceContainer string `json:"resourceContainer"`
	// udpIdleTimeout is how long an idle UDP connection will be kept open (e.g. '250ms', '2s').
	// Must be greater than 0. Only applicable for proxyMode=userspace.
	UDPIdleTimeout metav1.Duration `json:"udpTimeoutMilliseconds"`
	// conntrack contains conntrack-related configuration options.
	Conntrack KubeProxyConntrackConfiguration `json:"conntrack"`
	// configSyncPeriod is how often configuration from the apiserver is refreshed. Must be greater
	// than 0.
	ConfigSyncPeriod metav1.Duration `json:"configSyncPeriod"`
}

// Currently two modes of proxying are available: 'userspace' (older, stable) or 'iptables'
// (newer, faster). If blank, use the best-available proxy (currently iptables, but may
// change in future versions).  If the iptables proxy is selected, regardless of how, but
// the system's kernel or iptables versions are insufficient, this always falls back to the
// userspace proxy.
type ProxyMode string

const (
	ProxyModeUserspace ProxyMode = "userspace"
	ProxyModeIPTables  ProxyMode = "iptables"
)