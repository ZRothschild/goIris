package tool

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Long int64

func (t Long) MarshalJSON() ([]byte, error) {
	b := fmt.Sprintf(`"%d"`, t)
	return []byte(b), nil
}

func (t *Long) UnmarshalJSON(value []byte) error {
	var v = strings.TrimSpace(strings.Trim(string(value), "\""))
	if v == "" {
		*t = Long(0)
		return nil
	}
	num, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}
	*t = Long(num)
	return nil
}

func (t Long) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Long) UnmarshalText(data []byte) error {
	*t = t.FromString(string(data))
	return nil
}

func (t Long) FromString(str string) Long {
	return ParseLong(str)
}

func ParseLong(str string) Long {
	str = strings.TrimSpace(str)
	num, err := strconv.ParseInt(str, 10, 64)
	if nil != err {
		return 0
	}
	return Long(num)
}

func (t Long) String() string {
	return strconv.FormatInt(int64(t), 10)
}

func (t Long) Value() (driver.Value, error) {
	if t.IsZero() {
		return int64(0), nil
	}
	return int64(t), nil
}

func (t *Long) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.(int64); ok {
		*t = Long(v)
	} else if v, ok := value.([]uint8); ok {
		num, err := strconv.ParseInt(string(v), 10, 64)
		if err == nil {
			*t = Long(num)
		}
	}
	return nil
}

func (t Long) IsZero() bool {
	return t == 0
}

// 转换 工具
func LongToInt64(ls []Long) []uint64 {
	nums := make([]uint64, 0, len(ls))
	for _, l := range ls {
		nums = append(nums, uint64(l))
	}
	return nums
}
