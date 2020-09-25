package id

import (
	"errors"
	"github.com/spf13/viper"
	"net"
	"sync"
	"time"
)

// 获取新的程序生成的编号
func NextId(viper2 *viper.Viper) uint64 {
	id, err := SingletonSnowflakeKeyGen(viper2).NextId()
	if err != nil {
		return 0
	}
	return id
}

/*------------------------------------------------------Singleton----------------------------------------------------*/

var (
	once            = new(sync.Once)
	snowflakeKeyGen *Snowflake
)

func SingletonSnowflakeKeyGen(viper2 *viper.Viper) *Snowflake {
	once.Do(func() {
		snowflakeKeyGen = NewSnowflake(SnowflakeSettings{}, viper2)
	})
	return snowflakeKeyGen
}

/*----------------------------------------------------SnowflakeKeyGen------------------------------------------------*/

// Snowflake程序id生成器
// 源于Twitter的Snowflake算法
// 但由于原版算法对应的分布式层级结构太简单，所以目前的算法实际是Sony对Snowflake算法的改进版本的Sonyflake算法
// Sonyflake算法原版可参考github中的开源项目，当前算法有进一步细微调整
const (
	SnowflakeTimeUnit  = 1e7     // 时间单位，一纳秒的多少倍，1e6 = 一毫秒，1e7 = 百分之一秒，1e8 = 十分之一秒
	BitLenSequence     = 8       // 序列号的个数最多256个(0-255)，即每单位时间并发数，如时间单位是1e7，则单实例qps = 25600
	BitLenDataCenterId = 3       // 数据中心个数最多8个(0-7)，即同一个环境（生产、预发布、测试等）的数据中心（假设一个机房相同数据域的应用服务器集群只有一个，则数据中心数等于机房数）最多有8个
	BitLenMachineId    = 16      // 同一个数据中心下最多65536个应用实例（0-65535），默认是根据实例ip后两段算实例id（k8s环境动态创建Pod，也建议用此方式），所以需要预留255 * 255这么多
	BitLenTime         = 1 << 36 // 时间戳之差最大 = 2的36次方 * 时间单位 / 1e9 秒，目前的设计最多可以用21.79年就需要更新开始时间（随之还需要归档旧数据和更新次新数据id）
	// 总共63位，不超过bit64
)

type SnowflakeSettings struct {
	StartTime      time.Time
	DataCenterId   func() (uint16, error)
	MachineId      func() (uint16, error)
	CheckMachineId func(uint16) bool
}

type Snowflake struct {
	mutex        *sync.Mutex
	startTime    int64
	elapsedTime  int64
	sequence     uint16
	dataCenterId uint16
	machineId    uint16
}

func NewSnowflake(st SnowflakeSettings, viper2 *viper.Viper) *Snowflake {
	sf := new(Snowflake)
	sf.mutex = new(sync.Mutex)
	sf.sequence = uint16(1<<BitLenSequence - 1)

	if st.StartTime.After(time.Now()) {
		return nil
	}
	if st.StartTime.IsZero() {
		sf.startTime = toSnowflakeTime(time.Date(2018, 9, 26, 0, 0, 0, 0, time.UTC)) // 没有配置默认使用此时间
	} else {
		sf.startTime = toSnowflakeTime(st.StartTime)
	}

	var err error
	if st.MachineId == nil {
		sf.machineId, err = GetPrivateIPv4Id() // 没有配置会读机器内网ip后两段，然后计算出一个值
	} else {
		sf.machineId, err = st.MachineId()
	}
	if nil != err {
		err = nil
		sf.machineId = uint16(0)
	}
	if st.DataCenterId == nil {
		if id := viper2.GetInt("data_center_id"); id > 0 { // 没有配置会尝试从配置文件读取数据中心id
			sf.dataCenterId = uint16(id)
		} else { // 如果配置文件也没有，默认数据中心id为0
			sf.dataCenterId = uint16(0)
		}
	} else {
		sf.dataCenterId, err = st.DataCenterId()
		if nil != err {
			sf.dataCenterId = uint16(0)
		}
	}
	if st.CheckMachineId != nil && !st.CheckMachineId(sf.machineId) {
		return nil
	}

	return sf
}

func (sf *Snowflake) NextId() (uint64, error) {
	const maskSequence = uint16(1<<BitLenSequence - 1)

	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	current := getCurrentElapsedTime(sf.startTime)
	if sf.elapsedTime < current {
		sf.elapsedTime = current
		sf.sequence = 0
	} else { // sf.elapsedTime >= current
		sf.sequence = (sf.sequence + 1) & maskSequence
		if sf.sequence == 0 {
			sf.elapsedTime++
			overtime := sf.elapsedTime - current
			time.Sleep(getSleepTime(overtime))
		}
	}

	return sf.toId()
}

func toSnowflakeTime(t time.Time) int64 {
	return t.UTC().UnixNano() / SnowflakeTimeUnit
}

func getCurrentElapsedTime(startTime int64) int64 {
	return toSnowflakeTime(time.Now()) - startTime
}

func getSleepTime(overtime int64) time.Duration {
	return time.Duration(overtime)*10*time.Millisecond -
		time.Duration(time.Now().UTC().UnixNano()%SnowflakeTimeUnit)*time.Nanosecond
}

func (sf *Snowflake) toId() (uint64, error) {
	if sf.elapsedTime >= BitLenTime {
		return 0, errors.New("over the time limit")
	}

	return uint64(sf.elapsedTime)<<(BitLenSequence+BitLenDataCenterId+BitLenMachineId) |
		uint64(sf.sequence)<<(BitLenDataCenterId+BitLenMachineId) |
		uint64(sf.dataCenterId)<<BitLenMachineId |
		uint64(sf.machineId), nil
}

func PrivateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil && len(ip) > 3 &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func GetPrivateIPv4Id() (uint16, error) {
	ip, err := PrivateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func Decompose(id uint64) map[string]uint64 {
	const maskDataCenterId = uint64((1<<BitLenDataCenterId - 1) << BitLenMachineId)
	const maskMachineId = uint64(1<<BitLenMachineId - 1)

	dataCenterId := id & maskDataCenterId >> BitLenMachineId
	machineId := id & maskMachineId
	return map[string]uint64{
		"id":           id,
		"dataCenterId": dataCenterId,
		"machineId":    machineId,
	}
}
