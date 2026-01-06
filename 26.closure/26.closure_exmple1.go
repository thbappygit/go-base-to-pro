package main // main package — program execution এখান থেকেই শুরু

import "fmt" // fmt package — console এ print করার জন্য

// counter একটি function যেটা আরেকটা function return করে (closure)
func counter(num int) func() int { // num = initial value, return type: func() int

	count := num // count local variable, num এর মান কপি হলো (closure variable)

	// show হলো anonymous function (closure)
	// এটি outer function-এর count variable ব্যবহার করছে
	show := func() int {
		count++      // প্রতি call এ count এর মান ১ বাড়ে
		return count // updated count return করে
	}

	return show // show function return করা হচ্ছে
}

func main() { // main function — program execution শুরু

	c1 := counter(10) // counter call → নতুন closure তৈরি
	// count = 10 (Heap-এ যায়)

	fmt.Println(c1()) // c1 call → count = 11 → print 11
	fmt.Println(c1()) // আবার call → count = 12 → print 12
	fmt.Println(c1()) // আবার call → count = 13 → print 13

	c2 := counter(50) // নতুন counter call → নতুন closure
	// count = 50 (c1 থেকে সম্পূর্ণ আলাদা)

	fmt.Println(c2()) // c2 call → count = 51 → print 51
	fmt.Println(c2()) // আবার call → count = 52 → print 52
	fmt.Println(c2()) // আবার call → count = 53 → print 53
}
