package backupcr

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type backupOpts struct {
	errOut io.Writer

	pvcName   string
	retention string
	jobName   string
}

func NewBackupCRCommand(errOut io.Writer) *cobra.Command {
	opts := backupOpts{
		errOut: errOut,
	}

	cmd := &cobra.Command{
		Use:   "backup",
		Short: "Creates the etcdBackup CR to do a backup",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	opts.AddFlags(cmd.Flags())

	return cmd
}

func (o *backupOpts) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.pvcName, "pvc", "", "")
	fs.StringVar(&o.retention, "retention", "", "")
	fs.StringVar(&o.jobName, "job-name", "", "")
}

func (o *backupOpts) Validate() error {
	if o.pvcName == "" {
		return errors.New("missing required flag: --pvc")
	}
	if o.jobName == "" {
		return errors.New("missing required flag: --job-name")
	}
	return nil
}

func (o *backupOpts) Run() error {
	batchv1client
	return nil
}
