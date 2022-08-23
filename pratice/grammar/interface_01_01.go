package grammar

import "fmt"

// 일반적인 구조체
type member struct {
	first string
	last  string
}

/*
값은 하나 이상의 타입일수 있다.
*/

type secretMember struct {
	member
	ltk bool
}

func (s secretMember) speak() {
	fmt.Println("I am secretMember", s.first, s.last)
}

func (m member) speak() {
	fmt.Println("I am person", m.first, m.last)
}

type human interface {
	// 인터페이스는 값이 하나이상의 타입이 될 수 있게해준다.
	// speak() 가 붙은 모든 메소드는 type human이기도 하다
	speak()
}

func bar(h human) {
	fmt.Println("pass into bar", h)
}

func InterfaceExample() {
	sm1 := secretMember{
		member: member{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretMember{
		member: member{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	m1 := member{
		first: "Hi",
		last:  "Bye",
	}

	sm1.speak()
	sa2.speak()
	m1.speak()
	fmt.Println(m1)
	fmt.Println("=======================")

	// bar는 human과 secretMember을 취하며 secretMember와 member 모두 type human이기때문에 bar로 사용이가능하다.
	bar(sm1)
	bar(sa2)
	bar(m1)

}
