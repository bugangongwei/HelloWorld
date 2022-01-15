package util

/*
关于怎么生成 hex
*/

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	wg  = &sync.WaitGroup{}
	src = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// 随机 uint64 类型转换成 16 进制
func genHexFmt() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	i := rnd.Uint64()
	fmt.Println(i)
	return fmt.Sprintf("%x", i)
}

// 随机 uint64 类型转换成 16 进制
func genHexStrconv() string {
	wg.Done()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	i := rnd.Uint64()
	fmt.Println(i)
	return strconv.FormatUint(i, 16)
}

/*
使用 hex 包
并发安全
source fixed, 会生成重复的值
*/
func genHexRandom(n int) (string, error) {
	b := make([]byte, n)
	n, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

/*
使用 hex 包
非并发安全
source 按照时间 seed, 生成数据不重复
*/
func genHexRandomSource(n int) (string, error) {
	b := make([]byte, n)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func CurrHex() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			str := genHexFmt()
			fmt.Println(str)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("the end")
}
