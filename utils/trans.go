package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strings"
)

func RangeBeginToEnd(begin, end string) ([]string, error) {
	b := net.ParseIP(begin)
	if b == nil {
		return nil, errors.New("need valid ip")
	}
	e := net.ParseIP(end)
	if e == nil {
		return nil, errors.New("need valid ip")
	}
	bType := V4OrV6(begin)
	eType := V4OrV6(end)
	if bType != eType || bType == 0 {
		return nil, errors.New("need valid ip")
	}
	pools := []string{}
	if bType == 4 {
		bInt64 := InetAtoN(begin)
		eInt64 := InetAtoN(end)
		for i := bInt64; i <= eInt64; i++ {
			pools = append(pools, InetNtoA(i))
		}
	} else {
		bBigInt := IP6toInt(b)
		eBigInt := IP6toInt(e)
		for i := bBigInt; i.Cmp(eBigInt) == -1; i = i.Add(i, big.NewInt(1)) {
			ipStr, err := IPv6ByLong(i)
			if err != nil {
				return nil, err
			}
			pools = append(pools, ipStr)
		}
	}
	return pools, nil
}

func BigIntBeginAndEnd(begin, end string) (*big.Int, *big.Int, error) {
	b := net.ParseIP(begin)
	if b == nil {
		return nil, nil, errors.New("need valid ip")
	}
	e := net.ParseIP(end)
	if e == nil {
		return nil, nil, errors.New("need valid ip")
	}
	bType := V4OrV6(begin)
	eType := V4OrV6(end)
	if bType != eType || bType == 0 {
		return nil, nil, errors.New("need valid ip")
	}
	if bType == 4 {
		bInt64 := big.NewInt(0)
		bInt64.SetBytes(b.To4())
		eInt64 := big.NewInt(0)
		eInt64.SetBytes(e.To4())
		return bInt64, eInt64, nil
	} else {
		bBigInt := IP6toInt(b)
		eBigInt := IP6toInt(e)
		return bBigInt, eBigInt, nil
	}
}

// IPV6
func IP6toInt(IPv6Address net.IP) *big.Int {
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(IPv6Address.To16())
	return IPv6Int
}

// IPv6ByLong
// 将 big.Int 长整型转换成IPV6 地址
// converts a big integer represented by a string into an IPv6 address string
// 53174336768213711679990085974688268287=> "2801:0137:0000:0000:0000:ffff:ffff:ffff"
func IPv6ByLong(bi *big.Int) (string, error) {
	b255 := new(big.Int).SetBytes([]byte{255})
	var buf = make([]byte, 2)
	p := make([]string, 8)
	j := 0
	var i uint
	tmpint := new(big.Int)
	for i = 0; i < 16; i += 2 {
		tmpint.Rsh(bi, 120-i*8).And(tmpint, b255)
		bytes := tmpint.Bytes()
		if len(bytes) > 0 {
			buf[0] = bytes[0]
		} else {
			buf[0] = 0
		}
		tmpint.Rsh(bi, 120-(i+1)*8).And(tmpint, b255)
		bytes = tmpint.Bytes()
		if len(bytes) > 0 {
			buf[1] = bytes[0]
		} else {
			buf[1] = 0
		}
		p[j] = hex.EncodeToString(buf)
		j++
	}

	return strings.Join(p, ":"), nil
}

// IPV4
func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// IPV4
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func InetToBigInt(ip net.IP) *big.Int {
	bInt64 := big.NewInt(0)
	return bInt64.SetBytes(ip.To4())
}

func V4OrV6(s string) int {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return 4
		case ':':
			return 6
		}
	}
	return 0
}
