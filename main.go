package main

import (
	"fmt"
	"golang.org/x/crypto/ed25519"
)

func main() {
	const message = "alice"
	pk, sk, err := ed25519.GenerateKey(nil)//生成私钥和公钥
	if err != nil {
		fmt.Println("errors")
	}
	fmt.Println(pk)
	fmt.Println(sk)
	r,pi := Evaluate(sk,pk,[]byte(message))//生成随机数，证明
	fmt.Println("result:")
	fmt.Println(r)
	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println()
		fmt.Println("Proof:")
		fmt.Println(pi)
	}
	res,err := Verify(pk,pi,[]byte(message),r)//验证结果
	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println()
		fmt.Println("Result:")
		fmt.Println(res)
	}

}
