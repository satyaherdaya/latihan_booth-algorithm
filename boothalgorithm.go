package main

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) newNode(nilai int) {
	baru := &Node{
		next: nil,
		prev: nil,
		val:  nilai,
	}
	if L.head == nil && L.tail == nil {
		L.tail = baru
	} else {
		baru.next = L.head
		L.head.prev = baru
	}
	L.head = baru
}

func (L *List) tambahValue(ctr int, desimal int) {
	baru := L.tail
	for ctr < desimal {
		if baru != nil {
			baru.val += 1
		menu:
			if baru.val == 2 {
				baru.val = 0
				if baru.prev == nil {
					baru := &Node{
						next: nil,
						prev: nil,
						val:  1,
					}
					if L.head == nil && L.tail == nil {
						L.head = baru
					} else {
						baru.next = L.head
						L.head.prev = baru
					}
					L.head = baru
				} else {
					baru.prev.val += 1
					baru = baru.prev
					goto menu
				}
			}
			baru = L.tail
		}
		ctr += 1
	}
}

func (L *List) display() {
	print := L.head
	for print != nil {
		fmt.Print(print.val)
		print = print.next
	}
}

func main() {
	ctr := 0
	Q := &List{}
	M := &List{}
	var desimalM, desimalQ int
	fmt.Print("masukkan angka desimal M : ")
	fmt.Scan(&desimalM)
	fmt.Print("masukkan angka desimal Q : ")
	fmt.Scan(&desimalQ)

	fmt.Printf("M : ")
	M.newNode(0)
	M.tambahValue(ctr, desimalM)
	M.display()
	fmt.Println()

	fmt.Printf("Q : ")
	Q.newNode(0)
	Q.tambahValue(ctr, desimalQ)
	Q.display()
	fmt.Println()
}
