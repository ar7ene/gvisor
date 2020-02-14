// automatically generated by stateify.

// +build go1.12
// +build !go1.15

package vfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *epollInterestList) beforeSave() {}
func (x *epollInterestList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *epollInterestList) afterLoad() {}
func (x *epollInterestList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *epollInterestEntry) beforeSave() {}
func (x *epollInterestEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *epollInterestEntry) afterLoad() {}
func (x *epollInterestEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *VirtualDentry) beforeSave() {}
func (x *VirtualDentry) save(m state.Map) {
	x.beforeSave()
	m.Save("mount", &x.mount)
	m.Save("dentry", &x.dentry)
}

func (x *VirtualDentry) afterLoad() {}
func (x *VirtualDentry) load(m state.Map) {
	m.Load("mount", &x.mount)
	m.Load("dentry", &x.dentry)
}

func init() {
	state.Register("pkg/sentry/vfs.epollInterestList", (*epollInterestList)(nil), state.Fns{Save: (*epollInterestList).save, Load: (*epollInterestList).load})
	state.Register("pkg/sentry/vfs.epollInterestEntry", (*epollInterestEntry)(nil), state.Fns{Save: (*epollInterestEntry).save, Load: (*epollInterestEntry).load})
	state.Register("pkg/sentry/vfs.VirtualDentry", (*VirtualDentry)(nil), state.Fns{Save: (*VirtualDentry).save, Load: (*VirtualDentry).load})
}