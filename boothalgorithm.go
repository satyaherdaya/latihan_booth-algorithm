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

func (L *List) deleteLastNode() {
	temp := L.tail
	L.tail.prev.next = nil
	L.tail = temp.prev
	temp = nil
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

func (L *List) size() int {
	count := 0
	list := L.head
	for list != nil {
		count += 1
		list = list.next
	}
	return count
}

func autoBit(M *List, Q *List, A *List, sizeM int, sizeQ int) {
	if sizeM > sizeQ {
		M.newNode(0)
		for i := 0; i < ((sizeM + 1) - sizeQ); i++ {
			Q.newNode(0)
		}
		for i := 0; i < (sizeM + 1); i++ {
			A.newNode(0)
		}
	} else if sizeQ > sizeM {
		Q.newNode(0)
		for i := 0; i < ((sizeQ + 1) - sizeM); i++ {
			M.newNode(0)
		}
		for i := 0; i < (sizeQ + 1); i++ {
			A.newNode(0)
		}
	} else {
		M.newNode(0)
		Q.newNode(0)
		for i := 0; i < (sizeM + 1); i++ {
			A.newNode(0)
		}
	}
}

func (L *List) minBiner() {
	list := L.tail
	for list != nil {
		if list.val == 0 {
			list.val = 1
		} else if list.val == 1 {
			list.val = 0
		}
		list = list.prev
	}
	list = L.tail
	list.val += 1
	for list != nil {
		if list.val == 2 {
			if list.prev != nil {
				list.prev.val += 1
				list.val = 0
			} else {
				list.val = 0
				break
			}
		} else {
			list = list.prev
		}
	}

}

func APM(A *List, M *List, desimalM int) {
	fmt.Printf("A + M : ")
	fmt.Printf("\tA : ")
	A.display()
	fmt.Printf("\tM : ")
	M.display()

	if desimalM > 0 {
		if M.head.val == 1 {
			M.minBiner()
		}
	} else {
		if M.head.val == 0 {
			M.minBiner()
		}
	}

	listM := M.tail
	listA := A.tail

	for listA != nil {
		listA.val = listA.val + listM.val
		if listA.val == 2 {
			listA.val = 0
			if listA.prev != nil {
				listA.prev.val += 1
			} else {
				break
			}
		} else if listA.val == 3 {
			listA.val = 1
			if listA.prev != nil {
				listA.prev.val += 1
			}
		}
		listA = listA.prev
		listM = listM.prev
	}
	fmt.Printf(" \tA + M : ")
	A.display()
	fmt.Println()
}

func AMM(A *List, M *List) {
	M.minBiner()

	fmt.Printf("A - M : ")
	fmt.Printf("\tA : ")
	A.display()
	fmt.Printf("\t-M : ")
	M.display()

	listM := M.tail
	listA := A.tail

	for listA != nil {
		listA.val = listA.val + listM.val
		if listA.val == 2 {
			listA.val = 0
			if listA.prev != nil {
				listA.prev.val += 1
			} else {
				break
			}
		} else if listA.val == 3 {
			listA.val = 1
			if listA.prev != nil {
				listA.prev.val += 1
			}
		}
		listA = listA.prev
		listM = listM.prev
	}
	fmt.Printf("\tA - M : ")
	A.display()
	fmt.Println()
}

func rshift(A *List, Q *List) {
	listA := A.head

	if listA.val == 0 {
		A.newNode(0)
	} else if listA.val == 1 {
		A.newNode(1)
	}
	listA = A.tail
	Q.newNode(listA.val)
	A.deleteLastNode()
	Q.deleteLastNode()
}

func (L *List) display() {
	print := L.head
	for print != nil {
		fmt.Print(print.val)
		print = print.next
	}
}

func tableDataDisplay(A *List, Q *List, M *List, q int, desimalM int) {
	fmt.Printf("\t ")
	A.display()
	fmt.Printf(" \t | \t")
	Q.display()
	fmt.Printf(" \t | \t ")
	fmt.Printf("%d", q)
	fmt.Printf(" \t | \t")
	if desimalM > 0 {
		if M.head.val == 1 {
			M.minBiner()
		}
	} else {
		if M.head.val == 0 {
			M.minBiner()
		}
	}
	M.display()
	fmt.Printf(" \t | ")
}

func main() {
	ctr := 0
	ctrBit := 1
	q := 0
	Q := &List{}
	M := &List{}
	A := new(List)
	var desimalM, desimalQ int
	fmt.Print("masukkan angka desimal M : ")
	fmt.Scan(&desimalM)
	fmt.Print("masukkan angka desimal Q : ")
	fmt.Scan(&desimalQ)

	if desimalM < 0 && desimalQ < 0 {
		M.newNode(0)
		desM := desimalM * (-1)
		M.tambahValue(ctr, desM)

		Q.newNode(0)
		desQ := desimalQ * (-1)
		Q.tambahValue(ctr, desQ)

		autoBit(M, Q, A, M.size(), Q.size())
		M.minBiner()
		Q.minBiner()
	} else if desimalM < 0 && desimalQ > 0 {
		M.newNode(0)
		desM := desimalM * (-1)
		M.tambahValue(ctr, int(desM))

		Q.newNode(0)
		Q.tambahValue(ctr, desimalQ)

		autoBit(M, Q, A, M.size(), Q.size())
		M.minBiner()
	} else if desimalM > 0 && desimalQ < 0 {
		M.newNode(0)
		M.tambahValue(ctr, desimalM)

		Q.newNode(0)
		desQ := desimalQ * (-1)
		Q.tambahValue(ctr, int(desQ))

		autoBit(M, Q, A, M.size(), Q.size())
		Q.minBiner()
	} else if desimalM > 0 && desimalQ > 0 {
		M.newNode(0)
		M.tambahValue(ctr, desimalM)

		Q.newNode(0)
		Q.tambahValue(ctr, desimalQ)

		autoBit(M, Q, A, M.size(), Q.size())
	}

	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Printf("M : ")
	M.display()
	fmt.Println()

	fmt.Printf("Q : ")
	Q.display()
	fmt.Println()

	bit := M.size()
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Printf("\t A \t | \t Q \t | \t q \t | \t M \t | \t proccess")
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------")

	tableDataDisplay(A, Q, M, q, desimalM)
	fmt.Println()

	for ctrBit <= bit {
		tempTailQ := Q.tail
		if tempTailQ.val == 1 && q == 0 {
			tableDataDisplay(A, Q, M, q, desimalM)
			AMM(A, M)
			q = 1
			tableDataDisplay(A, Q, M, q, desimalM)
			fmt.Printf("After A-M")
			fmt.Println()
		} else if tempTailQ.val == 0 && q == 1 {
			tableDataDisplay(A, Q, M, q, desimalM)
			APM(A, M, desimalM)
			q = 0
			tableDataDisplay(A, Q, M, q, desimalM)
			fmt.Printf("After A+M")
			fmt.Println()
		} else if (tempTailQ.val == 0 && q == 0) || (tempTailQ.val == 1 && q == 1) {
			rshift(A, Q)
			ctrBit += 1
			tableDataDisplay(A, Q, M, q, desimalM)
			fmt.Printf("Right Shift")
			fmt.Println()
			tableDataDisplay(A, Q, M, q, desimalM)
			fmt.Printf("After Right Shift")
			fmt.Println()
		}
	}
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Printf("Hasil : ")
	A.display()
	fmt.Printf(" ")
	Q.display()
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------")
}
