package cmd

import (
	"time"

	"git.parallelcoin.io/dev/9/cmd/nine"
	"git.parallelcoin.io/dev/9/cmd/node"
)

var activenetparams = node.ActiveNetParams
var config = getConfig()
var Config = MakeConfig(config)
var stateconfig = node.StateCfg
var tn, sn, rn bool
var DataDir string
var ConfigFile string

func MakeConfig(c *Lines) (out *nine.Config) {
	cfg := *c
	String := func(path string) (out *string) {
		if cfg[path] != nil && cfg[path].Value != nil {
			return cfg[path].Value.(*string)
		}
		return
	}
	Tags := func(path string) (out *[]string) {
		if cfg[path] != nil && cfg[path].Value != nil {
			return cfg[path].Value.(*[]string)
		}
		return
	}
	Map := func(path string) (out *nine.Mapstringstring) {
		if cfg[path] != nil && cfg[path].Value != nil {
			return cfg[path].Value.(*nine.Mapstringstring)
		}
		return
	}
	Int := func(path string) (out *int) {
		if cfg[path] != nil && cfg[path].Value != nil {
			return cfg[path].Value.(*int)
		}
		return
	}
	Bool := func(path string) (out *bool) {
		if cfg[path] != nil && cfg[path].Value != nil {
			return cfg[path].Value.(*bool)
		}
		return
	}
	Float := func(path string) (out *float64) {
		if cfg[path] != nil && cfg[path].Value != nil {
			return cfg[path].Value.(*float64)
		}
		return
	}
	Duration := func(path string) (out *time.Duration) {
		if cfg[path] != nil && cfg[path].Value != nil {
			return cfg[path].Value.(*time.Duration)
		}
		return
	}

	out = &nine.Config{
		ConfigFile:               &ConfigFile,
		AppDataDir:               String("app.appdatadir"),
		DataDir:                  &DataDir,
		LogDir:                   String("app.logdir"),
		LogLevel:                 String("log.level"),
		Subsystems:               Map("log.subsystem"),
		Network:                  String("p2p.network"),
		AddPeers:                 Tags("p2p.addpeer"),
		ConnectPeers:             Tags("p2p.connect"),
		MaxPeers:                 Int("p2p.maxpeers"),
		Listeners:                Tags("p2p.listen"),
		DisableListen:            Bool("p2p.nolisten"),
		DisableBanning:           Bool("p2p.disableban"),
		BanDuration:              Duration("p2p.banduration"),
		BanThreshold:             Int("p2p.banthreshold"),
		Whitelists:               Tags("p2p.whitelist"),
		Username:                 String("rpc.user"),
		Password:                 String("rpc.pass"),
		ServerUser:               String("rpc.user"),
		ServerPass:               String("rpc.pass"),
		LimitUser:                String("limit.user"),
		LimitPass:                String("limit.pass"),
		RPCConnect:               String("rpc.connect"),
		RPCListeners:             Tags("rpc.listen"),
		RPCCert:                  String("tls.cert"),
		RPCKey:                   String("tls.key"),
		RPCMaxClients:            Int("rpc.maxclients"),
		RPCMaxWebsockets:         Int("rpc.maxwebsockets"),
		RPCMaxConcurrentReqs:     Int("rpc.maxconcurrentreqs"),
		RPCQuirks:                Bool("rpc.quirks"),
		DisableRPC:               Bool("rpc.disable"),
		NoTLS:                    Bool("tls.disable"),
		DisableDNSSeed:           Bool("p2p.nodns"),
		ExternalIPs:              Tags("p2p.externalips"),
		Proxy:                    String("proxy.address"),
		ProxyUser:                String("proxy.user"),
		ProxyPass:                String("proxy.pass"),
		OnionProxy:               String("proxy.address"),
		OnionProxyUser:           String("proxy.user"),
		OnionProxyPass:           String("proxy.pass"),
		Onion:                    Bool("proxy.tor"),
		TorIsolation:             Bool("proxy.isolation"),
		TestNet3:                 &tn,
		RegressionTest:           &rn,
		SimNet:                   &sn,
		AddCheckpoints:           Tags("chain.addcheckpoints"),
		DisableCheckpoints:       Bool("chain.disablecheckpoints"),
		DbType:                   String("chain.dbtype"),
		Profile:                  Int("app.profile"),
		CPUProfile:               String("app.cpuprofile"),
		Upnp:                     Bool("app.upnp"),
		MinRelayTxFee:            Float("p2p.minrelaytxfee"),
		FreeTxRelayLimit:         Float("p2p.freetxrelaylimit"),
		NoRelayPriority:          Bool("p2p.norelaypriority"),
		TrickleInterval:          Duration("p2p.trickleinterval"),
		MaxOrphanTxs:             Int("p2p.maxorphantxs"),
		Algo:                     String("mining.algo"),
		Generate:                 Bool("mining.generate"),
		GenThreads:               Int("mining.genthreads"),
		MiningAddrs:              Tags("mining.addresses"),
		MinerListener:            String("mining.listener"),
		MinerPass:                String("mining.pass"),
		BlockMinSize:             Int("block.minsize"),
		BlockMaxSize:             Int("block.maxsize"),
		BlockMinWeight:           Int("block.minweight"),
		BlockMaxWeight:           Int("block.maxweight"),
		BlockPrioritySize:        Int("block.prioritysize"),
		UserAgentComments:        Tags("p2p.useragentcomments"),
		NoPeerBloomFilters:       Bool("p2p.nobloomfilters"),
		NoCFilters:               Bool("p2p.nocfilters"),
		SigCacheMaxSize:          Int("chain.sigcachemaxsize"),
		BlocksOnly:               Bool("p2p.blocksonly"),
		TxIndex:                  Bool("chain.txindex"),
		AddrIndex:                Bool("chain.addrindex"),
		RelayNonStd:              Bool("chain.relaynonstd"),
		RejectNonStd:             Bool("chain.rejectnonstd"),
		TLSSkipVerify:            Bool("tls.skipverify"),
		Wallet:                   Bool("wallet.enable"),
		NoInitialLoad:            Bool("wallet.noinitialload"),
		WalletPass:               String("wallet.pass"),
		WalletServer:             String("rpc.wallet"),
		CAFile:                   String("tls.cafile"),
		OneTimeTLSKey:            Bool("tls.onetime"),
		ServerTLS:                Bool("tls.server"),
		LegacyRPCListeners:       Tags("rpc.listen"),
		LegacyRPCMaxClients:      Int("rpc.maxclients"),
		LegacyRPCMaxWebsockets:   Int("rpc.maxwebsockets"),
		ExperimentalRPCListeners: &[]string{},
	}
	return
}