// automatically generated by stateify.

package refs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *WeakRef) StateTypeName() string {
	return "pkg/refs.WeakRef"
}

func (x *WeakRef) StateFields() []string {
	return []string{
		"obj",
		"user",
	}
}

func (x *WeakRef) beforeSave() {}

func (x *WeakRef) StateSave(m state.Sink) {
	x.beforeSave()
	var obj savedReference = x.saveObj()
	m.SaveValue(0, obj)
	m.Save(1, &x.user)
}

func (x *WeakRef) afterLoad() {}

func (x *WeakRef) StateLoad(m state.Source) {
	m.Load(1, &x.user)
	m.LoadValue(0, new(savedReference), func(y interface{}) { x.loadObj(y.(savedReference)) })
}

func (x *AtomicRefCount) StateTypeName() string {
	return "pkg/refs.AtomicRefCount"
}

func (x *AtomicRefCount) StateFields() []string {
	return []string{
		"refCount",
		"name",
		"stack",
	}
}

func (x *AtomicRefCount) beforeSave() {}

func (x *AtomicRefCount) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.refCount)
	m.Save(1, &x.name)
	m.Save(2, &x.stack)
}

func (x *AtomicRefCount) afterLoad() {}

func (x *AtomicRefCount) StateLoad(m state.Source) {
	m.Load(0, &x.refCount)
	m.Load(1, &x.name)
	m.Load(2, &x.stack)
}

func (x *savedReference) StateTypeName() string {
	return "pkg/refs.savedReference"
}

func (x *savedReference) StateFields() []string {
	return []string{
		"obj",
	}
}

func (x *savedReference) beforeSave() {}

func (x *savedReference) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.obj)
}

func (x *savedReference) afterLoad() {}

func (x *savedReference) StateLoad(m state.Source) {
	m.Load(0, &x.obj)
}

func (x *weakRefList) StateTypeName() string {
	return "pkg/refs.weakRefList"
}

func (x *weakRefList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *weakRefList) beforeSave() {}

func (x *weakRefList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *weakRefList) afterLoad() {}

func (x *weakRefList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *weakRefEntry) StateTypeName() string {
	return "pkg/refs.weakRefEntry"
}

func (x *weakRefEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *weakRefEntry) beforeSave() {}

func (x *weakRefEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *weakRefEntry) afterLoad() {}

func (x *weakRefEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func init() {
	state.Register((*WeakRef)(nil))
	state.Register((*AtomicRefCount)(nil))
	state.Register((*savedReference)(nil))
	state.Register((*weakRefList)(nil))
	state.Register((*weakRefEntry)(nil))
}