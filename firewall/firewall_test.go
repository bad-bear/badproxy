package firewall

import (
	"test_code/binding"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestAddFrwlPort(t *testing.T) {
	type args struct {
		ptr binding.PortDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddFWRPort_Success",
			args: args{
				ptr: binding.PortDTO{
					Port:   "22",
					Target: "DROP",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddFrwlPort(tt.args.ptr)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestAdd_frwl_firewall(t *testing.T) {
	type args struct {
		ftr binding.FirewallRequestDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddFWRPort_Success",
			args: args{
				ftr: binding.FirewallRequestDTO{
					Table:       "INPUT",
					Protocol:    "udp",
					Source:      "8.8.8.8",
					Destination: "",
					Port:        "22",
					Target:      "DROP",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Add_frwl_firewall(tt.args.ftr)
			assert.Equal(t, tt.want, err)

		})
	}
}

func TestAdd_frwl_geo(t *testing.T) {
	type args struct {
		gtr binding.GeoDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddFWRPort_Success",
			args: args{
				gtr: binding.GeoDTO{
					Location: "Texas",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Add_frwl_geo(tt.args.gtr)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestAdd_frwl_iprate(t *testing.T) {
	type args struct {
		iptr binding.IPRateDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddFWRPort_Success",
			args: args{
				iptr: binding.IPRateDTO{
					Source:    "8.8.8.8",
					Conn:      "100",
					Protocol:  "tcp",
					TimeFrame: "m",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Add_frwl_iprate(tt.args.iptr)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestAdd_fwrl_bandwidth(t *testing.T) {
	type args struct {
		bndw binding.BandwidthDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddFWRPort_Success",
			args: args{
				bndw: binding.BandwidthDTO{
					Protocol: "tcp",
					Source:   "8.8.8.8",
					Time:     "m",
					Type:     "mb",
					Bamount:  "1000",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Add_fwrl_bandwidth(tt.args.bndw)
			assert.Equal(t, tt.want, err)
		})
	}
}
