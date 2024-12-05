// Code generated by cdpgen. DO NOT EDIT.

// Package fedcm implements the FedCM domain. This domain allows interacting
// with the FedCM dialog.
package fedcm

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the FedCM domain. This domain allows
// interacting with the FedCM dialog.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the FedCM domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the FedCM method.
func (d *domainClient) Enable(ctx context.Context, args *EnableArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "FedCm.enable", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "FedCm.enable", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "FedCM", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the FedCM method.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "FedCm.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "FedCM", Op: "Disable", Err: err}
	}
	return
}

// SelectAccount invokes the FedCM method.
func (d *domainClient) SelectAccount(ctx context.Context, args *SelectAccountArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "FedCm.selectAccount", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "FedCm.selectAccount", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "FedCM", Op: "SelectAccount", Err: err}
	}
	return
}

// ClickDialogButton invokes the FedCM method.
func (d *domainClient) ClickDialogButton(ctx context.Context, args *ClickDialogButtonArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "FedCm.clickDialogButton", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "FedCm.clickDialogButton", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "FedCM", Op: "ClickDialogButton", Err: err}
	}
	return
}

// OpenURL invokes the FedCM method.
func (d *domainClient) OpenURL(ctx context.Context, args *OpenURLArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "FedCm.openUrl", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "FedCm.openUrl", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "FedCM", Op: "OpenURL", Err: err}
	}
	return
}

// DismissDialog invokes the FedCM method.
func (d *domainClient) DismissDialog(ctx context.Context, args *DismissDialogArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "FedCm.dismissDialog", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "FedCm.dismissDialog", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "FedCM", Op: "DismissDialog", Err: err}
	}
	return
}

// ResetCooldown invokes the FedCM method. Resets the cooldown time, if any,
// to allow the next FedCM call to show a dialog even if one was recently
// dismissed by the user.
func (d *domainClient) ResetCooldown(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "FedCm.resetCooldown", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "FedCM", Op: "ResetCooldown", Err: err}
	}
	return
}

func (d *domainClient) DialogShown(ctx context.Context) (DialogShownClient, error) {
	s, err := rpcc.NewStream(ctx, "FedCm.dialogShown", d.conn)
	if err != nil {
		return nil, err
	}
	return &dialogShownClient{Stream: s}, nil
}

type dialogShownClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *dialogShownClient) GetStream() rpcc.Stream { return c.Stream }

func (c *dialogShownClient) Recv() (*DialogShownReply, error) {
	event := new(DialogShownReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "FedCM", Op: "DialogShown Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DialogClosed(ctx context.Context) (DialogClosedClient, error) {
	s, err := rpcc.NewStream(ctx, "FedCm.dialogClosed", d.conn)
	if err != nil {
		return nil, err
	}
	return &dialogClosedClient{Stream: s}, nil
}

type dialogClosedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *dialogClosedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *dialogClosedClient) Recv() (*DialogClosedReply, error) {
	event := new(DialogClosedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "FedCM", Op: "DialogClosed Recv", Err: err}
	}
	return event, nil
}
