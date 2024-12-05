// Code generated by cdpgen. DO NOT EDIT.

// Package systeminfo implements the SystemInfo domain. The SystemInfo domain
// defines methods and events for querying low-level system information.
package systeminfo

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the SystemInfo domain. The SystemInfo domain
// defines methods and events for querying low-level system information.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the SystemInfo domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// GetInfo invokes the SystemInfo method. Returns information about the
// system.
func (d *domainClient) GetInfo(ctx context.Context) (reply *GetInfoReply, err error) {
	reply = new(GetInfoReply)
	err = rpcc.Invoke(ctx, "SystemInfo.getInfo", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "SystemInfo", Op: "GetInfo", Err: err}
	}
	return
}

// GetFeatureState invokes the SystemInfo method. Returns information about
// the feature state.
func (d *domainClient) GetFeatureState(ctx context.Context, args *GetFeatureStateArgs) (reply *GetFeatureStateReply, err error) {
	reply = new(GetFeatureStateReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "SystemInfo.getFeatureState", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "SystemInfo.getFeatureState", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "SystemInfo", Op: "GetFeatureState", Err: err}
	}
	return
}

// GetProcessInfo invokes the SystemInfo method. Returns information about all
// running processes.
func (d *domainClient) GetProcessInfo(ctx context.Context) (reply *GetProcessInfoReply, err error) {
	reply = new(GetProcessInfoReply)
	err = rpcc.Invoke(ctx, "SystemInfo.getProcessInfo", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "SystemInfo", Op: "GetProcessInfo", Err: err}
	}
	return
}