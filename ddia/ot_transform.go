package main

// Operational Transformation
// Two operation types: Insert(pos, chars), Delete(pos, len)
// T(op1, op2) transforms op1 against op2 so both can be applied in either order.

type OpType int

const (
	Insert OpType = iota
	Delete
)

type Op struct {
	typ   OpType
	pos   int
	chars string // used for insert
	len   int    // used for delete
}

// transform returns op1' such that applying op2 then op1' ≡ applying op1 then op2
func transform(op1, op2 Op) Op {
	// if op1.pos < op2.pos: op1 position unaffected
	// if op1 is insert, op2 is insert, op1.pos > op2.pos: op1.pos += len(op2.chars)
	// if op1 is delete, op2 is insert, op1.pos >= op2.pos: op1.pos += len(op2.chars)
	// if op1 is insert, op2 is delete, op1.pos > op2.pos: op1.pos -= op2.len
	// if op1 is delete, op2 is delete, op1.pos > op2.pos: shift by op2.len
	// if op1.pos == op2.pos: tie-breaking rules (e.g., insert has priority, or site ID order)
	return Op{}
}

// apply applies op to doc and returns new doc
func apply(doc string, op Op) string {
	return ""
}
