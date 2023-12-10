package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddPortRule(t *testing.T) {
	tests := []struct {
		name           string
		formDataValues url.Values
		expectedStatus int
	}{
		{
			name: "AddPortRule_Good",
			formDataValues: url.Values{
				"Port":   {"22"},
				"Target": {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddPortRule_Good",
			formDataValues: url.Values{
				"Port":   {"22"},
				"Target": {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddPortRule_MissingPort_Bad",
			formDataValues: url.Values{
				"Port":   {""},
				"Target": {"ACCEPT"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddPortRule_MissTarget_Bad",
			formDataValues: url.Values{
				"Port":   {"22"},
				"Target": {""},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddPortRule_PortNotNumber_Bad",
			formDataValues: url.Values{
				"Port":   {"Test"},
				"Target": {"ACCEPT"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddPortRule_TargetNotAllCaps_Bad",
			formDataValues: url.Values{
				"Port":   {"22"},
				"Target": {"Accept"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddPortRule_TargetNotAllCaps_Bad",
			formDataValues: url.Values{
				"Port":   {"22"},
				"Target": {"Drop"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			govalidator.SetFieldsRequiredByDefault(true)

			req := httptest.NewRequest(http.MethodPost, "/add-frwl-port", strings.NewReader(tt.formDataValues.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			//log.Println(strings.NewReader(tt.formDataValues.Encode()))

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := AddPortRule(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

		})
	}
}

func TestAddFirewallRule(t *testing.T) {
	tests := []struct {
		name           string
		formDataValues url.Values
		expectedStatus int
	}{
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {"8.8.8.0"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"tcp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"udp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"icmp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {"8.8.8.0"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {"8.8.8.0"},
				"Destination":   {""},
				"Port_blocking": {"22"},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {"22"},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {"22"},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"22"},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"ACCEPT"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {"8.8.8.0"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"tcp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"udp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"icmp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {"8.8.8.0"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {"8.8.8.0"},
				"Destination":   {""},
				"Port_blocking": {"22"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {"22"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {"8.8.8.1"},
				"Destination":   {""},
				"Port_blocking": {"22"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"22"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"all"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {""},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"tcp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"udp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Good",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddFirewallRule_Bad",
			formDataValues: url.Values{
				"IP_Table":      {""},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddFirewallRule_Bad",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {""},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddFirewallRule_Bad",
			formDataValues: url.Values{
				"IP_Table":      {"OUTPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {""},
				"Port_blocking": {"23"},
				"Target":        {"DROP"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddFirewallRule_Bad",
			formDataValues: url.Values{
				"IP_Table":      {"INPUT"},
				"Protocol":      {"icmp"},
				"Source":        {""},
				"Destination":   {"8.8.8.1"},
				"Port_blocking": {"23"},
				"Target":        {""},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			govalidator.SetFieldsRequiredByDefault(true)

			req := httptest.NewRequest(http.MethodPost, "/add-frwl-firewall", strings.NewReader(tt.formDataValues.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := AddFirewallRule(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestAddGeoRule(t *testing.T) {
	tests := []struct {
		name           string
		formDataValues url.Values
		expectedStatus int
	}{
		{
			name: "AddGeoRule_Good",
			formDataValues: url.Values{
				"Location": {"Texas"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddGeoRule_Bad",
			formDataValues: url.Values{
				"Location": {""},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			govalidator.SetFieldsRequiredByDefault(true)

			req := httptest.NewRequest(http.MethodPost, "/add-frwl-geo", strings.NewReader(tt.formDataValues.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := AddGeoRule(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestAddIPrateRule(t *testing.T) {
	tests := []struct {
		name           string
		formDataValues url.Values
		expectedStatus int
	}{
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"all"},
				"TimeFrame":   {"s"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"icmp"},
				"TimeFrame":   {"s"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"tcp"},
				"TimeFrame":   {"s"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"udp"},
				"TimeFrame":   {"s"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"all"},
				"TimeFrame":   {"s"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"all"},
				"TimeFrame":   {"s"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"icmp"},
				"TimeFrame":   {"m"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"tcp"},
				"TimeFrame":   {"h"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"udp"},
				"TimeFrame":   {"h"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Good",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"all"},
				"TimeFrame":   {"d"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddIPrateRule_Bad",
			formDataValues: url.Values{

				"Source":      {""},
				"Protocol":    {"all"},
				"TimeFrame":   {"d"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddIPrateRule_Bad",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {""},
				"TimeFrame":   {"d"},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddIPrateRule_Bad",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"all"},
				"TimeFrame":   {""},
				"Connections": {"10"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddIPrateRule_Bad",
			formDataValues: url.Values{

				"Source":      {"8.8.8.0"},
				"Protocol":    {"all"},
				"TimeFrame":   {"d"},
				"Connections": {""},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			govalidator.SetFieldsRequiredByDefault(true)

			req := httptest.NewRequest(http.MethodPost, "/add-frwl-iprate", strings.NewReader(tt.formDataValues.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := AddIPrateRule(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestAddBandwidthRule(t *testing.T) {
	tests := []struct {
		name           string
		formDataValues url.Values
		expectedStatus int
	}{
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"d"},
				"Type":     {"p"},
				"Bamount":  {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"tcp"},
				"Time":     {"m"},
				"Type":     {"p"},
				"Bamount":  {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"icmp"},
				"Time":     {"h"},
				"Type":     {"p"},
				"Bamount":  {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"udp"},
				"Time":     {"s"},
				"Type":     {"p"},
				"Bamount":  {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {"mb"},
				"Bamount":  {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"h"},
				"Type":     {"b"},
				"Bamount":  {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"s"},
				"Type":     {"p"},
				"Bamount":  {"10"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Good",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {"mb"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {"d"},
				"Bamount":  {"d"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {"1"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"0"},
				"Type":     {"d"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"0"},
				"Time":     {"m"},
				"Type":     {"d"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"0"},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {"d"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {"d"},
				"Bamount":  {""},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {""},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {"all"},
				"Time":     {""},
				"Type":     {"d"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {"8.8.8.0"},
				"Protocol": {""},
				"Time":     {"m"},
				"Type":     {"d"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
		{
			name: "AddBandwidthRule_Bad",
			formDataValues: url.Values{

				"Source":   {""},
				"Protocol": {"all"},
				"Time":     {"m"},
				"Type":     {"d"},
				"Bamount":  {"100"},
			},
			expectedStatus: http.StatusNotAcceptable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			govalidator.SetFieldsRequiredByDefault(true)

			req := httptest.NewRequest(http.MethodPost, "/add-frwl-bandwidth", strings.NewReader(tt.formDataValues.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := AddBandwidthRule(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestGetDashboardData(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "GetDashboardData_Good",
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			govalidator.SetFieldsRequiredByDefault(true)

			req := httptest.NewRequest(http.MethodGet, "/dashboard-data", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := GetDashboardData(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestRotateLogTable(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "GetDashboardData_Good",
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			govalidator.SetFieldsRequiredByDefault(true)

			req := httptest.NewRequest(http.MethodGet, "/archive-logs", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := RotateLogTable(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
