# 포인터

포인터는 메모리 주소를 값으로 갖는 타입이다.

    var a int

    //아래 코드는 int타입의 메모리 주소를 값으로 가질수 있다.
    var p *int

    p = &a  

    *p = 20

위 코드는 a의 메모리 주소를 포인터 변수 p에 대입합니다. %는 주소, *는 주소가 가르키는 값 으로 이해하면 쉽다.




    type account_no_pointer struct {
        balance int
    }

    func (a account_no_pointer) changeBalance(b int){
        a.balance = b
    }

    func main() {
        account := account_no_pointer{balance: 100}
        account.balance = account.changeBalance(10)

        fmt.Printf("balnce = %d\n", account.balance)

        //balnce = 100
    }

account_no_pointer라는 타입의 account 변수를 선언했고(account라는 메모리 저장) changeBalance()라는 함수를 실행했다. go 에서는 모든 함수는 call은 복사로 일어난다. 다른 공간에 a account_no_pointer 를 저장 할 공간을 따로 만들고 기존의 값(첫줄에 선언한 account)을 복사해온다. changeBalance() 함수를 실행해도 값이 변경되지 않았다. 

    func (*a account_no_pointer) changeBalance(b int){
        a.balance = b // account의 메모리 주소가 가르키고 있는 값(balance)에 해당하는 속성을 변경시켜라
    }

위 코드로 변경하고 changeBalance(10)을 수행하면 balnce가 10으로 나게된다. 동작 방식은 다른 공간에 a account_no_pointer 를 만들고  제일 처음 선언한 account의 메모리 주소가 복사된다. changeBalance() 를 수행하면 account의 메모리 주소가 가르키고 있는 값(balance)에 해당하는 속성을 변경시킨다.

    func (a *account_no_pointer) changeBalance(b int) {
        fmt.Printf("a의 메모리주소 %p\n", a)
        a.balance = b
    }

    func main() {
        ...
        account := account_no_pointer{balance: 100}
        account.changeBalance(10)

        fmt.Printf("account의 메모리주소 %p\n", &account)
        ....
    }

두 메모리 주소를 print해보면 동일한 값이 출력된다.