package main

func maini2u3ihfiajoij1() {

}

type OrderedStream struct {
	ptr    int
	stream []string
}

func Constructorasdf(n int) OrderedStream {
	return OrderedStream{
		ptr:    1,
		stream: make([]string, n+1),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	if this.stream[this.ptr] == "" {
		return []string{}
	}
	var i = this.ptr
	for ; i < len(this.stream) && this.stream[i] != ""; i++ {
	}
	old := this.ptr
	this.ptr = i
	return this.stream[old:i]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
