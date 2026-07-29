package main

// Jupiter Protocol — state vector (cs, ss) per connected client
// cs: number of ops client has sent to server
// ss: number of ops server has sent to client

type JupiterState struct {
	cs int // client sent count
	ss int // server sent count
}

type JupiterOp struct {
	op    Op
	state JupiterState
}

// clientSend: send op with current (cs, ss)
//   increment cs; store op in pending buffer

// serverReceive: lookup client's last recorded state (lcs, lss)
//   n = cs - lcs (how many pending ops server has already processed from this client)
//   m = ss - lss (how many ops server sent since that state)
//   op' = transform(op, ops[n-1]) then track together, shift through m client-visible ops
//   apply op', record new state, broadcast to other clients

// serverAck: server sends back ack with (cs', ss')
//   client discards pending op matched by cs'
//   transforms remaining pending ops against the ack's implied state delta

func (j *JupiterState) clientSend(pending []JupiterOp, op Op) JupiterOp {
	return JupiterOp{}
}

func (j *JupiterState) serverReceive(clientState JupiterState, op Op) Op {
	return Op{}
}

func (j *JupiterState) serverAck(clientState JupiterState) {
}
