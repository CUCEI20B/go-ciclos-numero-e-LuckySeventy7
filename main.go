func main() {
	var n float64
	var j float64
	var i float64
	var fact float64//factorial
	var e float64

	fmt.Scan(&n)
	for j = 0; j < n; j++ {
		fact = 1 
		for i = 1; i <= n;	i++ {
			fact = fact * i;	
		} 
		e =1 +1+ 1.0/fact;
	}
	fmt.Println(e)
}