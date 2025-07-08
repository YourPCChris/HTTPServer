
package main

import (
    "strconv"
    "errors"
)

func Calc(num1, num2, op string) (float64, error){
    a, err := strconv.ParseFloat(num1, 64)
    if err != nil { return 0, errors.New("Calc Failed")}
    b, err := strconv.ParseFloat(num2, 64)
    if err != nil { return 0, errors.New("Calc Failed")}

    switch op{
    case "add":
        return a+b, nil
    case "sub":
        return a-b, nil
    case "mul":
        return a*b, nil
    case "div":
        if b != 0{
            return a/b,nil
        }
    }
    return 0, errors.New("Invalid Operation")
}
        
