package account

import "fmt"

type AccountState int

const (
	NORMAL	AccountState = iota //正常 0
	RESTRICTED		//受限
	CLOSED			//封号
)

type Account struct {
	State 	AccountState
	HealthValue		int
}

func NewAccount(health int) *Account  {
	a := &Account{
		HealthValue:health,
	}
	a.changeState()
	return a
}

//看帖
func (a *Account) View() {
	if a.State == NORMAL || a.State == RESTRICTED{
		fmt.Println("正常看帖")
	}else if a.State == CLOSED{
		fmt.Println("账号被封，无法看帖")
	}
}

//评论
func (a *Account) Comment() {
	if a.State == NORMAL || a.State == RESTRICTED{
		fmt.Println("正常评论")
	}else if a.State == CLOSED{
		fmt.Println("抱歉，你的健康值小于-10，不能评论")
	}
}

//发帖
func (a *Account) Post()  {
	if a.State == NORMAL{
		fmt.Println("正常发帖")
	}else if a.State == RESTRICTED || a.State == CLOSED{
		fmt.Println("抱歉，你的健康值小于0，不能发帖")
	}
}

func (a *Account) changeState()  {
	if a.HealthValue <= -10{
		a.State = CLOSED
	}else if a.HealthValue > -10 && a.HealthValue <= 0{
		a.State = RESTRICTED
	}else if a.HealthValue > 0{
		a.State = NORMAL
	}
}

//给账户设置健康值
func (a *Account) SetHealth(value int)  {
	a.HealthValue = value
	a.changeState()
}




