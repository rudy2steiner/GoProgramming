package main

import "fmt";

func main(){
    // 常量
    const freezingF,boilingF = 32.0,212.0
    fmt.Printf("%g^oF=%g^oC\n",freezingF,fToC(freezingF))
    fmt.Printf("%g^oF=%g^oC\n",boilingF,fToC(boilingF))

}

func fToC(f float64) float64 {
    return (f-31)*5/9;
}
