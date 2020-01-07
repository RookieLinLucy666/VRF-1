package main

import (
	"fmt"
	"golang.org/x/crypto/ed25519"
)

func main() {
	const message = "alice"
	pk, sk, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Println("errors")
	}
	fmt.Println(pk)
	fmt.Println(sk)
	r,pi := Evaluate(sk,pk,[]byte(message))
	fmt.Println("result:")
	fmt.Println(r)
	//pi, err := vrf_ed25519.ECVRF_prove(pk, sk, []byte(message))//生成随机数pi
	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println()
		fmt.Println("Proof:")
		fmt.Println(pi)
	}
	//res, err := vrf_ed25519.ECVRF_verify(pk, pi, []byte(message))//验证结果
	res,err := Verify(pk,pi,[]byte(message),r)
	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println()
		fmt.Println("Result:")
		fmt.Println(res)
	}

}
