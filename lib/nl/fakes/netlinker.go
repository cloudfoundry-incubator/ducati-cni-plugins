// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-cni-plugins/lib/nl"
	"github.com/vishvananda/netlink"
)

type Netlinker struct {
	LinkAddStub        func(link netlink.Link) error
	linkAddMutex       sync.RWMutex
	linkAddArgsForCall []struct {
		link netlink.Link
	}
	linkAddReturns struct {
		result1 error
	}
	LinkDelStub        func(link netlink.Link) error
	linkDelMutex       sync.RWMutex
	linkDelArgsForCall []struct {
		link netlink.Link
	}
	linkDelReturns struct {
		result1 error
	}
	LinkListStub        func() ([]netlink.Link, error)
	linkListMutex       sync.RWMutex
	linkListArgsForCall []struct{}
	linkListReturns struct {
		result1 []netlink.Link
		result2 error
	}
	LinkSetUpStub        func(link netlink.Link) error
	linkSetUpMutex       sync.RWMutex
	linkSetUpArgsForCall []struct {
		link netlink.Link
	}
	linkSetUpReturns struct {
		result1 error
	}
	LinkByNameStub        func(name string) (netlink.Link, error)
	linkByNameMutex       sync.RWMutex
	linkByNameArgsForCall []struct {
		name string
	}
	linkByNameReturns struct {
		result1 netlink.Link
		result2 error
	}
	LinkSetNsFdStub        func(link netlink.Link, fd int) error
	linkSetNsFdMutex       sync.RWMutex
	linkSetNsFdArgsForCall []struct {
		link netlink.Link
		fd   int
	}
	linkSetNsFdReturns struct {
		result1 error
	}
	AddrAddStub        func(link netlink.Link, addr *netlink.Addr) error
	addrAddMutex       sync.RWMutex
	addrAddArgsForCall []struct {
		link netlink.Link
		addr *netlink.Addr
	}
	addrAddReturns struct {
		result1 error
	}
	LinkSetMasterStub        func(slave netlink.Link, master *netlink.Bridge) error
	linkSetMasterMutex       sync.RWMutex
	linkSetMasterArgsForCall []struct {
		slave  netlink.Link
		master *netlink.Bridge
	}
	linkSetMasterReturns struct {
		result1 error
	}
	RouteAddStub        func(*netlink.Route) error
	routeAddMutex       sync.RWMutex
	routeAddArgsForCall []struct {
		arg1 *netlink.Route
	}
	routeAddReturns struct {
		result1 error
	}
}

func (fake *Netlinker) LinkAdd(link netlink.Link) error {
	fake.linkAddMutex.Lock()
	fake.linkAddArgsForCall = append(fake.linkAddArgsForCall, struct {
		link netlink.Link
	}{link})
	fake.linkAddMutex.Unlock()
	if fake.LinkAddStub != nil {
		return fake.LinkAddStub(link)
	} else {
		return fake.linkAddReturns.result1
	}
}

func (fake *Netlinker) LinkAddCallCount() int {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return len(fake.linkAddArgsForCall)
}

func (fake *Netlinker) LinkAddArgsForCall(i int) netlink.Link {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return fake.linkAddArgsForCall[i].link
}

func (fake *Netlinker) LinkAddReturns(result1 error) {
	fake.LinkAddStub = nil
	fake.linkAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *Netlinker) LinkDel(link netlink.Link) error {
	fake.linkDelMutex.Lock()
	fake.linkDelArgsForCall = append(fake.linkDelArgsForCall, struct {
		link netlink.Link
	}{link})
	fake.linkDelMutex.Unlock()
	if fake.LinkDelStub != nil {
		return fake.LinkDelStub(link)
	} else {
		return fake.linkDelReturns.result1
	}
}

func (fake *Netlinker) LinkDelCallCount() int {
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	return len(fake.linkDelArgsForCall)
}

func (fake *Netlinker) LinkDelArgsForCall(i int) netlink.Link {
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	return fake.linkDelArgsForCall[i].link
}

func (fake *Netlinker) LinkDelReturns(result1 error) {
	fake.LinkDelStub = nil
	fake.linkDelReturns = struct {
		result1 error
	}{result1}
}

func (fake *Netlinker) LinkList() ([]netlink.Link, error) {
	fake.linkListMutex.Lock()
	fake.linkListArgsForCall = append(fake.linkListArgsForCall, struct{}{})
	fake.linkListMutex.Unlock()
	if fake.LinkListStub != nil {
		return fake.LinkListStub()
	} else {
		return fake.linkListReturns.result1, fake.linkListReturns.result2
	}
}

func (fake *Netlinker) LinkListCallCount() int {
	fake.linkListMutex.RLock()
	defer fake.linkListMutex.RUnlock()
	return len(fake.linkListArgsForCall)
}

func (fake *Netlinker) LinkListReturns(result1 []netlink.Link, result2 error) {
	fake.LinkListStub = nil
	fake.linkListReturns = struct {
		result1 []netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *Netlinker) LinkSetUp(link netlink.Link) error {
	fake.linkSetUpMutex.Lock()
	fake.linkSetUpArgsForCall = append(fake.linkSetUpArgsForCall, struct {
		link netlink.Link
	}{link})
	fake.linkSetUpMutex.Unlock()
	if fake.LinkSetUpStub != nil {
		return fake.LinkSetUpStub(link)
	} else {
		return fake.linkSetUpReturns.result1
	}
}

func (fake *Netlinker) LinkSetUpCallCount() int {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return len(fake.linkSetUpArgsForCall)
}

func (fake *Netlinker) LinkSetUpArgsForCall(i int) netlink.Link {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return fake.linkSetUpArgsForCall[i].link
}

func (fake *Netlinker) LinkSetUpReturns(result1 error) {
	fake.LinkSetUpStub = nil
	fake.linkSetUpReturns = struct {
		result1 error
	}{result1}
}

func (fake *Netlinker) LinkByName(name string) (netlink.Link, error) {
	fake.linkByNameMutex.Lock()
	fake.linkByNameArgsForCall = append(fake.linkByNameArgsForCall, struct {
		name string
	}{name})
	fake.linkByNameMutex.Unlock()
	if fake.LinkByNameStub != nil {
		return fake.LinkByNameStub(name)
	} else {
		return fake.linkByNameReturns.result1, fake.linkByNameReturns.result2
	}
}

func (fake *Netlinker) LinkByNameCallCount() int {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return len(fake.linkByNameArgsForCall)
}

func (fake *Netlinker) LinkByNameArgsForCall(i int) string {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return fake.linkByNameArgsForCall[i].name
}

func (fake *Netlinker) LinkByNameReturns(result1 netlink.Link, result2 error) {
	fake.LinkByNameStub = nil
	fake.linkByNameReturns = struct {
		result1 netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *Netlinker) LinkSetNsFd(link netlink.Link, fd int) error {
	fake.linkSetNsFdMutex.Lock()
	fake.linkSetNsFdArgsForCall = append(fake.linkSetNsFdArgsForCall, struct {
		link netlink.Link
		fd   int
	}{link, fd})
	fake.linkSetNsFdMutex.Unlock()
	if fake.LinkSetNsFdStub != nil {
		return fake.LinkSetNsFdStub(link, fd)
	} else {
		return fake.linkSetNsFdReturns.result1
	}
}

func (fake *Netlinker) LinkSetNsFdCallCount() int {
	fake.linkSetNsFdMutex.RLock()
	defer fake.linkSetNsFdMutex.RUnlock()
	return len(fake.linkSetNsFdArgsForCall)
}

func (fake *Netlinker) LinkSetNsFdArgsForCall(i int) (netlink.Link, int) {
	fake.linkSetNsFdMutex.RLock()
	defer fake.linkSetNsFdMutex.RUnlock()
	return fake.linkSetNsFdArgsForCall[i].link, fake.linkSetNsFdArgsForCall[i].fd
}

func (fake *Netlinker) LinkSetNsFdReturns(result1 error) {
	fake.LinkSetNsFdStub = nil
	fake.linkSetNsFdReturns = struct {
		result1 error
	}{result1}
}

func (fake *Netlinker) AddrAdd(link netlink.Link, addr *netlink.Addr) error {
	fake.addrAddMutex.Lock()
	fake.addrAddArgsForCall = append(fake.addrAddArgsForCall, struct {
		link netlink.Link
		addr *netlink.Addr
	}{link, addr})
	fake.addrAddMutex.Unlock()
	if fake.AddrAddStub != nil {
		return fake.AddrAddStub(link, addr)
	} else {
		return fake.addrAddReturns.result1
	}
}

func (fake *Netlinker) AddrAddCallCount() int {
	fake.addrAddMutex.RLock()
	defer fake.addrAddMutex.RUnlock()
	return len(fake.addrAddArgsForCall)
}

func (fake *Netlinker) AddrAddArgsForCall(i int) (netlink.Link, *netlink.Addr) {
	fake.addrAddMutex.RLock()
	defer fake.addrAddMutex.RUnlock()
	return fake.addrAddArgsForCall[i].link, fake.addrAddArgsForCall[i].addr
}

func (fake *Netlinker) AddrAddReturns(result1 error) {
	fake.AddrAddStub = nil
	fake.addrAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *Netlinker) LinkSetMaster(slave netlink.Link, master *netlink.Bridge) error {
	fake.linkSetMasterMutex.Lock()
	fake.linkSetMasterArgsForCall = append(fake.linkSetMasterArgsForCall, struct {
		slave  netlink.Link
		master *netlink.Bridge
	}{slave, master})
	fake.linkSetMasterMutex.Unlock()
	if fake.LinkSetMasterStub != nil {
		return fake.LinkSetMasterStub(slave, master)
	} else {
		return fake.linkSetMasterReturns.result1
	}
}

func (fake *Netlinker) LinkSetMasterCallCount() int {
	fake.linkSetMasterMutex.RLock()
	defer fake.linkSetMasterMutex.RUnlock()
	return len(fake.linkSetMasterArgsForCall)
}

func (fake *Netlinker) LinkSetMasterArgsForCall(i int) (netlink.Link, *netlink.Bridge) {
	fake.linkSetMasterMutex.RLock()
	defer fake.linkSetMasterMutex.RUnlock()
	return fake.linkSetMasterArgsForCall[i].slave, fake.linkSetMasterArgsForCall[i].master
}

func (fake *Netlinker) LinkSetMasterReturns(result1 error) {
	fake.LinkSetMasterStub = nil
	fake.linkSetMasterReturns = struct {
		result1 error
	}{result1}
}

func (fake *Netlinker) RouteAdd(arg1 *netlink.Route) error {
	fake.routeAddMutex.Lock()
	fake.routeAddArgsForCall = append(fake.routeAddArgsForCall, struct {
		arg1 *netlink.Route
	}{arg1})
	fake.routeAddMutex.Unlock()
	if fake.RouteAddStub != nil {
		return fake.RouteAddStub(arg1)
	} else {
		return fake.routeAddReturns.result1
	}
}

func (fake *Netlinker) RouteAddCallCount() int {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return len(fake.routeAddArgsForCall)
}

func (fake *Netlinker) RouteAddArgsForCall(i int) *netlink.Route {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return fake.routeAddArgsForCall[i].arg1
}

func (fake *Netlinker) RouteAddReturns(result1 error) {
	fake.RouteAddStub = nil
	fake.routeAddReturns = struct {
		result1 error
	}{result1}
}

var _ nl.Netlinker = new(Netlinker)
