package ride

import (
	"github.com/pkg/errors"
)

func ccc(f Fsm, node Node) (Fsm, error) {
	switch n := node.(type) {
	case *AssignmentNode:
		fsm, err := ccc(f.Assigment(n.Name), n.Expression)
		if err != nil {
			return fsm, err
		}
		return ccc(fsm.Return(), n.Block)
	case *LongNode:
		return f.Long(n.Value), nil
	case *FunctionCallNode:
		var err error
		f = f.Call(n.Name, uint16(len(n.Arguments)))
		for i := range n.Arguments {
			f, err = ccc(f, n.Arguments[i])
			if err != nil {
				return f, err
			}
		}
		return f.Return(), nil
	case *ReferenceNode:
		return f.Reference(n.Name), nil
	case *BooleanNode:
		return f.Boolean(n.Value), nil
	case *StringNode:
		return f.String(n.Value), nil
	case *BytesNode:
		return f.Bytes(n.Value), nil
	case *ConditionalNode:
		f, err := ccc(f.Condition(), n.Condition)
		if err != nil {
			return f, err
		}
		f, err = ccc(f.TrueBranch(), n.TrueExpression)
		if err != nil {
			return f, err
		}
		f, err = ccc(f.FalseBranch(), n.FalseExpression)
		if err != nil {
			return f, err
		}
		return f.Return(), nil
	case *FunctionDeclarationNode:
		fsm, err := ccc(f.FuncDeclaration(n.Name, n.Arguments), n.Body)
		if err != nil {
			return fsm, err
		}
		return ccc(fsm.Return(), n.Block)
	case *PropertyNode:
		//f = f.Call("$property", 2)
		//f, err := ccc(f, n.Object)
		//if err != nil {
		//	return f, err
		//}
		//f.(params).b.writeByte(OpProperty)

		return ccc(f.Property(n.Name), n.Object)
		//n.Object
		//n.Name
	default:
		return f, errors.Errorf("unknown type %T", node)
	}
}

func CompileSimpleScript(t *Tree) (*Executable, error) {
	return compileSimpleScript(t.LibVersion, t.Verifier)
}

func compileSimpleScript(libVersion int, node Node) (*Executable, error) {
	fCheck, err := selectFunctionChecker(libVersion)
	if err != nil {
		return nil, err
	}
	b := newBuilder()
	c := newConstants()
	r := newReferences(nil)
	u := &uniqid{}

	//predefs := newPredef()
	//predefs.set("tx", math.MaxUint16, tx)

	params := params{
		b:      b,
		c:      c,
		r:      r,
		f:      fCheck,
		u:      u,
		predef: predefined,
	}

	f := NewDefinitionsFsm(params)
	f, err = ccc(f, node)
	if err != nil {
		return nil, err
	}
	// Just to write `OpReturn` to bytecode.
	f = f.Return()

	return f.(BuildExecutable).BuildExecutable(libVersion), nil

}
