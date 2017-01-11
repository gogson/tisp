package vm

import (
	"sync"
	"sync/atomic"
)

type thunkState int32

const (
	illegal thunkState = iota
	normal
	locked
	app
)

type Thunk struct {
	result   Object
	function *Thunk
	// `args` is represented as a slice but not a List to let users optimize
	// their functions. If you want to define functions with arguments fully
	// lazy, just create a function which takes only a thunk of a List as a
	// argument.
	args      []*Thunk
	state     thunkState
	blackHole sync.WaitGroup
}

func Normal(o Object) *Thunk {
	if f, ok := o.(func(...*Thunk) Object); ok {
		o = NewLazyFunction(f)
	} else if f, ok := o.(func(...Object) Object); ok {
		o = NewStrictFunction(f)
	}

	return &Thunk{result: o, state: normal}
}

func App(f *Thunk, args ...*Thunk) *Thunk {
	t := &Thunk{function: f, args: args, state: app}
	t.blackHole.Add(1)
	return t
}

func (t *Thunk) Eval() Object { // return WHNF
	if t.lock() {
		children := make([]*Thunk, 0, 1) // TODO: best capacity?

		for {
			o := t.function.Eval()
			f, ok := o.(Callable)

			if !ok {
				t.result = NotCallableError(o)
				break
			}

			t.result = f.Call(t.args...)
			child, ok := t.result.(*Thunk)

			if !ok {
				break
			}

			t.function, t.args, ok = child.delegateEval()

			if !ok {
				t.result = child.Eval()
				break
			}

			children = append(children, child)
		}

		for _, child := range children {
			child.result = t.result
			child.finalize()
		}

		t.finalize()
	} else {
		t.blackHole.Wait()
	}

	return t.result
}

func (t *Thunk) lock() bool {
	return t.compareAndSwapState(app, locked)
}

func (t *Thunk) delegateEval() (*Thunk, []*Thunk, bool) {
	if t.lock() {
		return t.function, t.args, true
	}

	return nil, nil, false
}

func (t *Thunk) finalize() {
	t.function = nil
	t.args = nil
	t.storeState(normal)
	t.blackHole.Done()
}

func (t *Thunk) compareAndSwapState(old, new thunkState) bool {
	return atomic.CompareAndSwapInt32((*int32)(&t.state), int32(old), int32(new))
}

func (t *Thunk) storeState(new thunkState) {
	atomic.StoreInt32((*int32)(&t.state), int32(new))
}
