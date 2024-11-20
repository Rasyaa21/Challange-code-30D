    var a int
    var b float64
    var c string

    scanner.Scan()
    input1 := scanner.Text()
    a, err := strconv.Atoi(input1)
    if err != nil {
        fmt.Println(err)
    }
    
    scanner.Scan()
    input2 := scanner.Text()
    b, err2 := strconv.ParseFloat(input2, 64)
    if err != nil {
        fmt.Println(err2)
    }
    
    scanner.Scan()
    c = scanner.Text()
    sum1 := a + int(i)
    sum2 := float64(b + d)

    fmt.Println(sum1)
    fmt.Printf("%.1f\n", sum2)
    fmt.Println(s + c)

