package config

import (
	"path"

	cmn "github.com/tendermint/tmlibs/common"
)

/****** these are for production settings ***********/
func EnsureRoot(rootDir string, network string) {
	cmn.EnsureDir(rootDir, 0700)
	cmn.EnsureDir(rootDir+"/data", 0700)

	configFilePath := path.Join(rootDir, "config.toml")

	// Write default config file if missing.
	if !cmn.FileExists(configFilePath) {
		cmn.MustWriteFile(configFilePath, []byte(selectNetwork(network)), 0644)
	}
}

var defaultConfigTmpl = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml
fast_sync = true
db_backend = "leveldb"
api_addr = "0.0.0.0:9888"
`

var mainNetConfigTmpl = `chain_id = "mainnet"
[p2p]
laddr = "tcp://0.0.0.0:46657"
seeds = "45.79.213.28:46657,198.74.61.131:46657,212.111.41.245:46657,47.100.214.154:46657,47.100.109.199:46657,47.100.105.165:46657"
`

var testNetConfigTmpl = `chain_id = "testnet"
[p2p]
laddr = "tcp://0.0.0.0:46656"
seeds = "47.96.42.1:46656,172.104.224.219:46656,45.118.132.164:46656"
`

var soloNetConfigTmpl = `chain_id = "solonet"
[p2p]
laddr = "tcp://0.0.0.0:46658"
seeds = ""
`

// Select network seeds to merge a new string.
func selectNetwork(network string) string {
	if network == "testnet" {
		return defaultConfigTmpl + testNetConfigTmpl
	} else if network == "mainnet" {
		return defaultConfigTmpl + mainNetConfigTmpl
	} else {
		return defaultConfigTmpl + soloNetConfigTmpl
	}
}
