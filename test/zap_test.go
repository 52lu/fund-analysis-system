/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/7/5 6:29 下午
 **/
package test

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestZap(t *testing.T) {
	fmt.Printf("DebugLevel: %v\n",zapcore.DebugLevel)
	fmt.Printf("InfoLevel: %v\n",zapcore.InfoLevel)
	fmt.Printf("WarnLevel: %v\n",zapcore.WarnLevel)
	fmt.Printf("-1: %v\n", zapcore.Level(-1))
	fmt.Printf("0: %v\n", zapcore.Level(0))
	fmt.Printf("1: %v\n", zapcore.Level(1))
	fmt.Printf("2: %v\n", zapcore.Level(2))
}
