package redisinfo_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/geoffreybauduin/redis-info"
	td "github.com/maxatome/go-testdeep"
	"github.com/stretchr/testify/assert"
)

func Test_Full(t *testing.T) {
	for _, file := range [2]string{"./resources/output.txt", "./resources/output_backslashr.txt"} {
		t.Run(fmt.Sprintf("testing file %s", file), func(tt *testing.T) {
			content, err := ioutil.ReadFile(file)
			assert.NoError(tt, err)
			info, err := redisinfo.Parse(string(content))
			assert.NoError(tt, err)
			td.CmpStruct(tt, info, redisinfo.Info{
				Server: redisinfo.Server{
					RedisVersion:    "3.2.5",
					RedisGitSha1:    "00000000",
					RedisGitDirty:   false,
					RedisBuildID:    "6f9920d2ae584aa0",
					RedisMode:       "standalone",
					OS:              "Linux 4.14.66-ovh-vps-grsec-zfs-classid x86_64",
					ArchBits:        64,
					MultiplexingAPI: "epoll",
					GCCVersion:      "4.9.2",
					ProcessID:       7,
					RunID:           "a4fcd0061e667352b73ad678f944d41b23c9c67a",
					TCPPort:         6379,
					UptimeInDays:    24,
					UptimeInSeconds: 2089968,
					HZ:              10,
					LRUClock:        1870516,
					Executable:      "/data/redis-server",
					ConfigFile:      "/redis.conf",
				},
				Clients: redisinfo.Client{
					ConnectedClients:        914,
					ClientLongestOutputList: 0,
					ClientBiggestInputBuf:   0,
					BlockedClients:          20,
				},
				Memory: redisinfo.Memory{
					UsedMemory:             6146978688,
					UsedMemoryHuman:        "5.72G",
					UsedMemoryRss:          6454341632,
					UsedMemoryRssHuman:     "6.01G",
					UsedMemoryPeak:         6290131040,
					UsedMemoryPeakHuman:    "5.86G",
					TotalSystemMemory:      135210360832,
					TotalSystemMemoryHuman: "125.92G",
					UsedMemoryLua:          37888,
					UsedMemoryLuaHuman:     "37.00K",
					Maxmemory:              0,
					MaxmemoryHuman:         "0B",
					MaxmemoryPolicy:        "noeviction",
					MemFragmentationRatio:  1.05,
					MemAllocator:           "jemalloc-4.0.3",
				},
				Persistence: redisinfo.Persistence{
					Loading:                  false,
					RdbChangesSinceLastSave:  9629911,
					RdbBgsaveInProgress:      false,
					RdbLastSaveTime:          1578891904,
					RdbLastBgsaveStatus:      "ok",
					RdbLastBgsaveTimeSec:     19,
					RdbCurrentBgsaveTimeSec:  -1,
					AofEnabled:               false,
					AofRewriteInProgress:     false,
					AofRewriteScheduled:      false,
					AofCurrentRewriteTimeSec: -1,
					AofLastRewriteTimeSec:    -1,
					AofLastBgrewriteStatus:   "ok",
					AofLastWriteStatus:       "ok",
				},
				Stats: redisinfo.Stats{
					TotalConnectionsReceived: 1209273,
					TotalCommandsProcessed:   1322487410,
					InstantaneousOpsPerSec:   448,
					TotalNetInputBytes:       498224475818,
					TotalNetOutputBytes:      3210532894008,
					InstantaneousInputKbps:   125.06,
					InstantaneousOutputKbps:  601.62,
					RejectedConnections:      0,
					SyncFull:                 53,
					SyncPartialOk:            77,
					SyncPartialErr:           4,
					ExpiredKeys:              24052947,
					EvictedKeys:              0,
					KeyspaceHits:             227762505,
					KeyspaceMisses:           8586015,
					PubsubChannels:           20373,
					PubsubPatterns:           40,
					LatestForkUsec:           189282,
					MigrateCachedSockets:     false,
				},
				Replication: redisinfo.Replication{
					Role:            "master",
					ConnectedSlaves: 3,
					Slaves: []redisinfo.ReplicationSlave{
						{
							ID:     0,
							IP:     "127.0.0.1",
							Port:   31000,
							State:  "online",
							Offset: 222662780961,
							Lag:    0,
						},
						{
							ID:     1,
							IP:     "127.0.0.1",
							Port:   31003,
							State:  "online",
							Offset: 222662776891,
							Lag:    0,
						},
						{
							ID:     2,
							IP:     "127.0.0.1",
							Port:   31001,
							State:  "online",
							Offset: 222662775148,
							Lag:    0,
						},
					},
					MasterReplOffset:           222662781103,
					ReplBacklogActive:          true,
					ReplBacklogSize:            1048576,
					ReplBacklogFirstByteOffset: 222661732528,
					ReplBacklogHistLen:         1048576,
				},
				CPU: redisinfo.CPU{
					UsedCPUSys:          44232.89,
					UsedCPUUser:         16130.08,
					UsedCPUSysChildren:  29.10,
					UsedCPUUserChildren: 259.26,
				},
				Cluster: redisinfo.Cluster{
					ClusterEnabled: false,
				},
				Keyspace: []redisinfo.Keyspace{
					{
						DB:      1,
						Keys:    58,
						Expires: 0,
						AvgTTL:  0,
					},
					{
						DB:      3,
						Keys:    307122,
						Expires: 307122,
						AvgTTL:  5549135,
					},
					{
						DB:      4,
						Keys:    6,
						Expires: 6,
						AvgTTL:  88504,
					},
					{
						DB:      6,
						Keys:    10453,
						Expires: 0,
						AvgTTL:  0,
					},
					{
						DB:      7,
						Keys:    7,
						Expires: 5,
						AvgTTL:  133782010,
					},
					{
						DB:      11,
						Keys:    32,
						Expires: 32,
						AvgTTL:  137917,
					},
					{
						DB:      12,
						Keys:    38,
						Expires: 38,
						AvgTTL:  115430,
					},
					{
						DB:      13,
						Keys:    98863,
						Expires: 98863,
						AvgTTL:  85569597,
					},
				},
			}, td.StructFields{}, "got correct structure")
		})
	}
}

func Test_OnlyServer(t *testing.T) {
	content, err := ioutil.ReadFile("./resources/output-server.txt")
	assert.NoError(t, err)
	info, err := redisinfo.Parse(string(content))
	assert.NoError(t, err)
	td.CmpStruct(t, info, redisinfo.Info{
		Server: redisinfo.Server{
			RedisVersion:    "3.2.5",
			RedisGitSha1:    "00000000",
			RedisGitDirty:   false,
			RedisBuildID:    "6f9920d2ae584aa0",
			RedisMode:       "standalone",
			OS:              "Linux 4.14.66-ovh-vps-grsec-zfs-classid x86_64",
			ArchBits:        64,
			MultiplexingAPI: "epoll",
			GCCVersion:      "4.9.2",
			ProcessID:       7,
			RunID:           "a4fcd0061e667352b73ad678f944d41b23c9c67a",
			TCPPort:         6379,
			UptimeInDays:    24,
			UptimeInSeconds: 2089968,
			HZ:              10,
			LRUClock:        1870516,
			Executable:      "/data/redis-server",
			ConfigFile:      "/redis.conf",
		},
	}, td.StructFields{}, "got correct structure")
}
