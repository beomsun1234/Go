# goroutine

![image](https://user-images.githubusercontent.com/68090443/209807098-0cce9786-b69e-4af0-bea7-85ff8beff4bf.png)

![image](https://user-images.githubusercontent.com/68090443/209807128-a67181ed-9894-4013-83bc-bafe8f314559.png)

##  sync.WaitGroup

sync.WaitGroup은 모든 고루틴이 종료될 때까지 대기해야 할 때 사용한다. 


    Add(delta int): WaitGroup에 대기 중인 고루틴 개수 추가
    Done(): 대기 중인 고루틴의 수행이 종료되는 것을 알려줌
    Wait(): 모든 고루틴이 종료될 때까지 대기
    
ex) goRT1, goRT2 가 끝날 때까지 기다리고 'end' 라는 문자열을 출력한다.

    sg := &sync.WaitGroup{}
    sg.Add(2)

    go goRT1("1", ch1, ch3)
    go goRT2(ch2)

    for i := 0; i < 2; i++ {
      select {
      case a := <-ch1:
        fmt.Println("ch1 = ", a, ", ch3 = ", <-ch3)
        sg.Done()
      case b := <-ch2:
        fmt.Println("ch2 =", b)
        sg.Done()
      }
    }
    
    sg.Wait()
    fmt.Println("end")
