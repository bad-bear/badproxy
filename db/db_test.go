package db

import (
	"test_code/binding"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestInsert_db_port(t *testing.T) {
	type args struct {
		fd binding.PortDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddPortRuletoDB_Success",
			args: args{
				fd: binding.PortDTO{
					Port:   "80",
					Target: "ACCEPT",
				},
			},
			want: nil,
		},
		{
			name: "AddPortRuletoDB_Success",
			args: args{
				fd: binding.PortDTO{
					Port:   "22",
					Target: "DROP",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Insert_db_port(tt.args.fd)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestInsert_db_firewall(t *testing.T) {
	type args struct {
		ftr binding.FirewallRequestDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddFirewallRuletoDB_Success",
			args: args{
				ftr: binding.FirewallRequestDTO{
					Table:       "INPUT",
					Protocol:    "tcp",
					Source:      "8.8.8.8",
					Destination: "8.8.8.8",
					Port:        "",
					Target:      "DROP",
				},
			},
			want: nil,
		},
		{
			name: "AddFirewallRuletoDB_Success",
			args: args{
				ftr: binding.FirewallRequestDTO{
					Table:       "INPUT",
					Protocol:    "tcp",
					Source:      "8.8.8.8",
					Destination: "",
					Port:        "22",
					Target:      "ACCEPT",
				},
			},
			want: nil,
		},
		{
			name: "AddFirewallRuletoDB_Success",
			args: args{
				ftr: binding.FirewallRequestDTO{
					Table:       "OUTPUT",
					Protocol:    "udp",
					Source:      "",
					Destination: "8.8.8.8",
					Port:        "",
					Target:      "DROP",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Insert_db_firewall(tt.args.ftr)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestInsert_db_geo(t *testing.T) {
	type args struct {
		gr binding.GeoDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "InsertGEOtoDB_Success",
			args: args{
				gr: binding.GeoDTO{
					Location: "Arizona",
				},
			},
			want: nil,
		},
		{
			name: "InsertGEOtoDB_Success",
			args: args{
				gr: binding.GeoDTO{
					Location: "Texas",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Insert_db_geo(tt.args.gr)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestInsert_db_iprate(t *testing.T) {
	type args struct {
		ipr binding.IPRateDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "",
			args: args{
				ipr: binding.IPRateDTO{
					Source:    "8.8.8.8",
					Conn:      "10",
					Protocol:  "tcp",
					TimeFrame: "m",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Insert_db_iprate(tt.args.ipr)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestInsert_db_bandwidth(t *testing.T) {
	type args struct {
		br binding.BandwidthDTO
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "AddBandwidthtoDB",
			args: args{
				br: binding.BandwidthDTO{
					Protocol: "tcp",
					Source:   "8.8.8.8",
					Time:     "m",
					Bamount:  "10",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Insert_db_bandwidth(tt.args.br)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestUpdate_log_table(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "UpdateTableLogs",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Update_log_table()
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestUpdate_dashboard_values(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "UpdateDashboardValues_Success",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Update_dashboard_values()
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestRotate_log_table(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "RotateLogs_Success",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Rotate_log_table()
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestGet_dashboard_data(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "GetDashboardData_Success",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Get_dashboard_data()
			assert.Equal(t, tt.want, err)
		})
	}
}

/*
func TestGet_log_data(t *testing.T) {
	tests := []struct {
		name string
		want []ip_log
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Get_log_data(); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("Get_log_data() = %v, want %v", err, tt.want)
			}
		})
	}
}
*/
