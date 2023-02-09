// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package status

import (
	"context"
	"fmt"

	rolloutsapi "github.com/GoogleContainerTools/kpt/rollouts/api/v1alpha1"
	"github.com/GoogleContainerTools/kpt/rollouts/rolloutsclient"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func newRunner(ctx context.Context) *runner {
	r := &runner{
		ctx: ctx,
	}
	c := &cobra.Command{
		Use:     "status",
		Short:   "displays status of a rollout",
		Long:    "displays status of a rollout",
		Example: "displays status of a rollout",
		RunE:    r.runE,
	}
	r.Command = c
	return r
}

func NewCommand(ctx context.Context) *cobra.Command {
	return newRunner(ctx).Command
}

type runner struct {
	ctx     context.Context
	Command *cobra.Command
}

func (r *runner) runE(cmd *cobra.Command, args []string) error {
	rlc, err := rolloutsclient.New()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	if len(args) == 0 {
		fmt.Printf("must provide rollout name")
		return nil
	}

	rollout, err := rlc.Get(r.ctx, args[0])
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	if len(rollout.Status.WaveStatuses) > 0 {
		renderWaveStatusAsTable(cmd, rollout)
		return nil
	}
	renderStatusAsTable(cmd, rollout)
	return nil
}

func renderStatusAsTable(cmd *cobra.Command, rollout *rolloutsapi.Rollout) {
	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())
	t.AppendHeader(table.Row{"CLUSTER", "PACKAGE ID", "PACKAGE STATUS", "SYNC STATUS"})
	for _, cluster := range rollout.Status.ClusterStatuses {
		pkgStatus := cluster.PackageStatus
		t.AppendRow([]interface{}{cluster.Name, pkgStatus.PackageID, pkgStatus.Status, pkgStatus.SyncStatus})
	}
	t.AppendSeparator()
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.Render()
}

func renderWaveStatusAsTable(cmd *cobra.Command, rollout *rolloutsapi.Rollout) {
	t := table.NewWriter()
	t.SetOutputMirror(cmd.OutOrStdout())
	t.AppendHeader(table.Row{"WAVE", "CLUSTER", "PACKAGE ID", "PACKAGE STATUS", "SYNC STATUS"})
	for _, wave := range rollout.Status.WaveStatuses {
		for i, cluster := range wave.ClusterStatuses {
			pkgStatus := cluster.PackageStatus
			waveName := ""
			if i == 0 {
				waveName = wave.Name
			}
			t.AppendRow([]interface{}{waveName, cluster.Name, pkgStatus.PackageID, pkgStatus.Status, pkgStatus.SyncStatus})
		}
		t.AppendSeparator()
	}
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.Render()
}
