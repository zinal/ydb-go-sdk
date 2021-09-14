// Code generated by gtrace. DO NOT EDIT.

package ydb

import (
	"context"
)

// Compose returns a new DriverTrace which has functional fields composed
// both from t and x.
func (t DriverTrace) Compose(x DriverTrace) (ret DriverTrace) {
	switch {
	case t.OnDial == nil:
		ret.OnDial = x.OnDial
	case x.OnDial == nil:
		ret.OnDial = t.OnDial
	default:
		h1 := t.OnDial
		h2 := x.OnDial
		ret.OnDial = func(d DialStartInfo) func(DialDoneInfo) {
			r1 := h1(d)
			r2 := h2(d)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(d DialDoneInfo) {
					r1(d)
					r2(d)
				}
			}
		}
	}
	switch {
	case t.OnGetConn == nil:
		ret.OnGetConn = x.OnGetConn
	case x.OnGetConn == nil:
		ret.OnGetConn = t.OnGetConn
	default:
		h1 := t.OnGetConn
		h2 := x.OnGetConn
		ret.OnGetConn = func(g GetConnStartInfo) func(GetConnDoneInfo) {
			r1 := h1(g)
			r2 := h2(g)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(g GetConnDoneInfo) {
					r1(g)
					r2(g)
				}
			}
		}
	}
	switch {
	case t.OnPessimization == nil:
		ret.OnPessimization = x.OnPessimization
	case x.OnPessimization == nil:
		ret.OnPessimization = t.OnPessimization
	default:
		h1 := t.OnPessimization
		h2 := x.OnPessimization
		ret.OnPessimization = func(p PessimizationStartInfo) func(PessimizationDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PessimizationDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.TrackConnStart == nil:
		ret.TrackConnStart = x.TrackConnStart
	case x.TrackConnStart == nil:
		ret.TrackConnStart = t.TrackConnStart
	default:
		h1 := t.TrackConnStart
		h2 := x.TrackConnStart
		ret.TrackConnStart = func(t TrackConnStartInfo) {
			h1(t)
			h2(t)
		}
	}
	switch {
	case t.TrackConnDone == nil:
		ret.TrackConnDone = x.TrackConnDone
	case x.TrackConnDone == nil:
		ret.TrackConnDone = t.TrackConnDone
	default:
		h1 := t.TrackConnDone
		h2 := x.TrackConnDone
		ret.TrackConnDone = func(t TrackConnDoneInfo) {
			h1(t)
			h2(t)
		}
	}
	switch {
	case t.OnGetCredentials == nil:
		ret.OnGetCredentials = x.OnGetCredentials
	case x.OnGetCredentials == nil:
		ret.OnGetCredentials = t.OnGetCredentials
	default:
		h1 := t.OnGetCredentials
		h2 := x.OnGetCredentials
		ret.OnGetCredentials = func(g GetCredentialsStartInfo) func(GetCredentialsDoneInfo) {
			r1 := h1(g)
			r2 := h2(g)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(g GetCredentialsDoneInfo) {
					r1(g)
					r2(g)
				}
			}
		}
	}
	switch {
	case t.OnDiscovery == nil:
		ret.OnDiscovery = x.OnDiscovery
	case x.OnDiscovery == nil:
		ret.OnDiscovery = t.OnDiscovery
	default:
		h1 := t.OnDiscovery
		h2 := x.OnDiscovery
		ret.OnDiscovery = func(d DiscoveryStartInfo) func(DiscoveryDoneInfo) {
			r1 := h1(d)
			r2 := h2(d)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(d DiscoveryDoneInfo) {
					r1(d)
					r2(d)
				}
			}
		}
	}
	switch {
	case t.OnOperation == nil:
		ret.OnOperation = x.OnOperation
	case x.OnOperation == nil:
		ret.OnOperation = t.OnOperation
	default:
		h1 := t.OnOperation
		h2 := x.OnOperation
		ret.OnOperation = func(o OperationStartInfo) func(OperationDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OperationDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	switch {
	case t.OnStream == nil:
		ret.OnStream = x.OnStream
	case x.OnStream == nil:
		ret.OnStream = t.OnStream
	default:
		h1 := t.OnStream
		h2 := x.OnStream
		ret.OnStream = func(s StreamStartInfo) func(StreamDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	return ret
}
func (t DriverTrace) onDial(d DialStartInfo) func(DialDoneInfo) {
	fn := t.OnDial
	if fn == nil {
		return func(DialDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DialDoneInfo) {
			return
		}
	}
	return res
}
func (t DriverTrace) onGetConn(g GetConnStartInfo) func(GetConnDoneInfo) {
	fn := t.OnGetConn
	if fn == nil {
		return func(GetConnDoneInfo) {
			return
		}
	}
	res := fn(g)
	if res == nil {
		return func(GetConnDoneInfo) {
			return
		}
	}
	return res
}
func (t DriverTrace) onPessimization(p PessimizationStartInfo) func(PessimizationDoneInfo) {
	fn := t.OnPessimization
	if fn == nil {
		return func(PessimizationDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PessimizationDoneInfo) {
			return
		}
	}
	return res
}
func (t DriverTrace) trackConnStart(t1 TrackConnStartInfo) {
	fn := t.TrackConnStart
	if fn == nil {
		return
	}
	fn(t1)
}
func (t DriverTrace) trackConnDone(t1 TrackConnDoneInfo) {
	fn := t.TrackConnDone
	if fn == nil {
		return
	}
	fn(t1)
}
func (t DriverTrace) onGetCredentials(g GetCredentialsStartInfo) func(GetCredentialsDoneInfo) {
	fn := t.OnGetCredentials
	if fn == nil {
		return func(GetCredentialsDoneInfo) {
			return
		}
	}
	res := fn(g)
	if res == nil {
		return func(GetCredentialsDoneInfo) {
			return
		}
	}
	return res
}
func (t DriverTrace) onDiscovery(d DiscoveryStartInfo) func(DiscoveryDoneInfo) {
	fn := t.OnDiscovery
	if fn == nil {
		return func(DiscoveryDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DiscoveryDoneInfo) {
			return
		}
	}
	return res
}
func (t DriverTrace) onOperation(o OperationStartInfo) func(OperationDoneInfo) {
	fn := t.OnOperation
	if fn == nil {
		return func(OperationDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OperationDoneInfo) {
			return
		}
	}
	return res
}
func (t DriverTrace) onStream(s StreamStartInfo) func(StreamDoneInfo) {
	fn := t.OnStream
	if fn == nil {
		return func(StreamDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(StreamDoneInfo) {
			return
		}
	}
	return res
}
func driverTraceOnDial(t DriverTrace, c context.Context, address string) func(_ context.Context, address string, _ error) {
	var p DialStartInfo
	p.Context = c
	p.Address = address
	res := t.onDial(p)
	return func(c context.Context, address string, e error) {
		var p DialDoneInfo
		p.Context = c
		p.Address = address
		p.Error = e
		res(p)
	}
}
func driverTraceOnGetConn(t DriverTrace, c context.Context) func(_ context.Context, address string, _ error) {
	var p GetConnStartInfo
	p.Context = c
	res := t.onGetConn(p)
	return func(c context.Context, address string, e error) {
		var p GetConnDoneInfo
		p.Context = c
		p.Address = address
		p.Error = e
		res(p)
	}
}
func driverTraceOnPessimization(t DriverTrace, c context.Context, address string, cause error) func(_ context.Context, address string, _ error) {
	var p PessimizationStartInfo
	p.Context = c
	p.Address = address
	p.Cause = cause
	res := t.onPessimization(p)
	return func(c context.Context, address string, e error) {
		var p PessimizationDoneInfo
		p.Context = c
		p.Address = address
		p.Error = e
		res(p)
	}
}
func driverTraceTrackConnStart(t DriverTrace, address string) {
	var p TrackConnStartInfo
	p.Address = address
	t.trackConnStart(p)
}
func driverTraceTrackConnDone(t DriverTrace, address string) {
	var p TrackConnDoneInfo
	p.Address = address
	t.trackConnDone(p)
}
func driverTraceOnGetCredentials(t DriverTrace, c context.Context) func(_ context.Context, token bool, _ error) {
	var p GetCredentialsStartInfo
	p.Context = c
	res := t.onGetCredentials(p)
	return func(c context.Context, token bool, e error) {
		var p GetCredentialsDoneInfo
		p.Context = c
		p.Token = token
		p.Error = e
		res(p)
	}
}
func driverTraceOnDiscovery(t DriverTrace, c context.Context) func(_ context.Context, endpoints []Endpoint, _ error) {
	var p DiscoveryStartInfo
	p.Context = c
	res := t.onDiscovery(p)
	return func(c context.Context, endpoints []Endpoint, e error) {
		var p DiscoveryDoneInfo
		p.Context = c
		p.Endpoints = endpoints
		p.Error = e
		res(p)
	}
}
func driverTraceOnOperation(t DriverTrace, c context.Context, address string, m Method, params OperationParams) func(_ context.Context, address string, _ Method, params OperationParams, opID string, issues IssueIterator, _ error) {
	var p OperationStartInfo
	p.Context = c
	p.Address = address
	p.Method = m
	p.Params = params
	res := t.onOperation(p)
	return func(c context.Context, address string, m Method, params OperationParams, opID string, issues IssueIterator, e error) {
		var p OperationDoneInfo
		p.Context = c
		p.Address = address
		p.Method = m
		p.Params = params
		p.OpID = opID
		p.Issues = issues
		p.Error = e
		res(p)
	}
}
func driverTraceOnStream(t DriverTrace, c context.Context, address string, m Method) func(_ context.Context, address string, _ Method, _ error) {
	var p StreamStartInfo
	p.Context = c
	p.Address = address
	p.Method = m
	res := t.onStream(p)
	return func(c context.Context, address string, m Method, e error) {
		var p StreamDoneInfo
		p.Context = c
		p.Address = address
		p.Method = m
		p.Error = e
		res(p)
	}
}
