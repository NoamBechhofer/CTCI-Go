package threeinone

import (
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type Operation int

const (
	IsEmpty = iota
	Peek
	Pop
	Push
	Size
)

func fuzzTarget(t *testing.T, data []byte) {
	control := []*lib.Stack[int32]{nil, {}, {}, {}}
	test := NewThreeInOne[int32]()

	for i := 0; i+6 <= len(data); i += 6 {
		step := i / 6
		stack := StackId(data[i]%3) + 1
		op := Operation(data[i+1] % 5)
		param := int32(binary.LittleEndian.Uint32(data[i+2 : i+6]))

		switch op {
		case IsEmpty:
			want := control[stack].IsEmpty()
			got := test.IsEmpty(stack)
			if want != got {
				t.Fatalf("step=%d stack=%d op=%d IsEmpty: want %t, got %t",
					step, stack, op, want, got)
			}

		case Peek:
			wantV, wantOK := control[stack].Peek()
			gotV, gotOK := test.Peek(stack)
			if wantV != gotV || wantOK != gotOK {
				t.Fatalf("step=%d stack=%d op=%d Peek: want (%d, %t), got (%d, %t)",
					step, stack, op, wantV, wantOK, gotV, gotOK)
			}

		case Pop:
			wantV, wantOK := control[stack].Pop()
			gotV, gotOK := test.Pop(stack)
			if wantV != gotV || wantOK != gotOK {
				t.Fatalf("step=%d stack=%d op=%d Pop: want (%d, %t), got (%d, %t)",
					step, stack, op, wantV, wantOK, gotV, gotOK)
			}

		case Push:
			control[stack].Push(param)
			test.Push(stack, param)

		case Size:
			want := control[stack].Size()
			got := test.Size(stack)
			if want != got {
				t.Fatalf("step=%d stack=%d op=%d Size: want %d, got %d",
					step, stack, op, want, got)
			}

		default:
			panic("unknown op")
		}
	}
}

type seedStep struct {
	stack byte // logical stack id: 1, 2, 3
	op    Operation
	param int32
}

func step(stack byte, op Operation, param int32) seedStep {
	if stack < 1 || stack > 3 {
		panic("invalid stack")
	}
	return seedStep{stack: stack, op: op, param: param}
}

func seed(steps ...seedStep) []byte {
	out := make([]byte, 0, 6*len(steps))
	for _, st := range steps {
		out = append(out, st.stack-1, byte(st.op), 0, 0, 0, 0) // -1 because fuzzTarget does %3 + 1
		binary.LittleEndian.PutUint32(out[len(out)-4:], uint32(st.param))
	}
	return out
}
func makeRandomProgram(r *rand.Rand, steps int, bursty bool) []byte {
	data := make([]byte, 0, 6*steps)

	currentStack := byte(r.Intn(3)) // fuzzTarget maps 0,1,2 -> stacks 1,2,3
	burstLeft := 0

	for i := 0; i < steps; i++ {
		stack := byte(r.Intn(3))
		if bursty {
			if burstLeft == 0 {
				currentStack = byte(r.Intn(3))
				burstLeft = 1 + r.Intn(20)
			}
			burstLeft--
			stack = currentStack
		}

		var op Operation
		if bursty {
			// More state-changing ops to force collisions, shuffles, and growth.
			switch n := r.Intn(10); {
			case n < 6:
				op = Push
			case n < 8:
				op = Pop
			case n < 9:
				op = Peek
			default:
				op = Size
			}
		} else {
			switch n := r.Intn(10); {
			case n < 5:
				op = Push
			case n < 7:
				op = Pop
			case n < 8:
				op = Peek
			case n < 9:
				op = Size
			default:
				op = IsEmpty
			}
		}

		param := uint32(r.Uint32())

		data = append(data, stack, byte(op), 0, 0, 0, 0)
		binary.LittleEndian.PutUint32(data[len(data)-4:], param)
	}

	return data
}

func FuzzThreeInOne(f *testing.F) {
	// 1. Basic smoke: all ops, all stacks.
	f.Add(seed(
		step(1, IsEmpty, 0),
		step(1, Push, 10),
		step(1, Peek, 0),
		step(1, Size, 0),
		step(1, Pop, 0),

		step(2, IsEmpty, 0),
		step(2, Push, 20),
		step(2, Peek, 0),
		step(2, Size, 0),
		step(2, Pop, 0),

		step(3, IsEmpty, 0),
		step(3, Push, 30),
		step(3, Peek, 0),
		step(3, Size, 0),
		step(3, Pop, 0),
	))

	// 2. Simple LIFO behavior on each stack.
	f.Add(seed(
		step(1, Push, 1), step(1, Push, 2), step(1, Peek, 0), step(1, Pop, 0), step(1, Pop, 0),
		step(2, Push, 3), step(2, Push, 4), step(2, Peek, 0), step(2, Pop, 0), step(2, Pop, 0),
		step(3, Push, 5), step(3, Push, 6), step(3, Peek, 0), step(3, Pop, 0), step(3, Pop, 0),
	))

	// 3. shuffle2ToMiddle(right):
	// size1=3, size2=1, size3=2 => collision on push1 with >1 free slot.
	f.Add(seed(
		step(1, Push, 11), step(1, Push, 12), step(1, Push, 13),
		step(2, Push, 21),
		step(3, Push, 31), step(3, Push, 32),
		step(1, Push, 14),
		step(1, Peek, 0), step(2, Peek, 0), step(3, Peek, 0),
		step(1, Size, 0), step(2, Size, 0), step(3, Size, 0),
	))

	// 4. shuffle2ToMiddle(left):
	// size1=1, size2=2, size3=4 => collision on push2 with >1 free slot.
	f.Add(seed(
		step(1, Push, 11),
		step(2, Push, 21), step(2, Push, 22),
		step(3, Push, 31), step(3, Push, 32), step(3, Push, 33), step(3, Push, 34),
		step(2, Push, 23),
		step(1, Peek, 0), step(2, Peek, 0), step(3, Peek, 0),
		step(1, Size, 0), step(2, Size, 0), step(3, Size, 0),
	))

	// 5. shuffle2Rightward():
	// availableSlots==1 and push1 collides.
	f.Add(seed(
		step(1, Push, 11), step(1, Push, 12), step(1, Push, 13),
		step(2, Push, 21),
		step(3, Push, 31), step(3, Push, 32), step(3, Push, 33), step(3, Push, 34),
		step(1, Push, 14),
		step(1, Peek, 0), step(2, Peek, 0), step(3, Peek, 0),
		step(1, Size, 0), step(2, Size, 0), step(3, Size, 0),
	))

	// 6. shuffle2Leftward():
	// availableSlots==1 and push2 collides.
	f.Add(seed(
		step(1, Push, 11), step(1, Push, 12),
		step(2, Push, 21), step(2, Push, 22), step(2, Push, 23),
		step(3, Push, 31), step(3, Push, 32), step(3, Push, 33),
		step(2, Push, 24),
		step(1, Peek, 0), step(2, Peek, 0), step(3, Peek, 0),
		step(1, Size, 0), step(2, Size, 0), step(3, Size, 0),
	))

	// 7. grow():
	// full array, then push1 collides and must grow.
	f.Add(seed(
		step(1, Push, 11), step(1, Push, 12), step(1, Push, 13),
		step(2, Push, 21),
		step(3, Push, 31), step(3, Push, 32), step(3, Push, 33), step(3, Push, 34), step(3, Push, 35),
		step(1, Push, 14),
		step(1, Peek, 0), step(2, Peek, 0), step(3, Peek, 0),
		step(1, Size, 0), step(2, Size, 0), step(3, Size, 0),
	))

	// 8. Empty-middle edge case: push1 forces stack2 relocation while stack2 is empty.
	f.Add(seed(
		step(1, Push, 1), step(1, Push, 2), step(1, Push, 3),
		step(1, Push, 4),
		step(1, Peek, 0), step(1, Size, 0),
		step(3, Push, 9), step(3, Peek, 0), step(3, Size, 0),
	))

	// 9-12. Deterministic long random programs.
	for _, tc := range []struct {
		seed   int64
		steps  int
		bursty bool
	}{
		{1, 1_000, false},
		{2, 10_000, false},
		{11, 10_000, true},
		{12, 50_000, true},
	} {
		r := rand.New(rand.NewSource(tc.seed))
		f.Add(makeRandomProgram(r, tc.steps, tc.bursty))
	}

	f.Fuzz(fuzzTarget)
}
