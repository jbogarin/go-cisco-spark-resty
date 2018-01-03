package ciscosparkresty

// License is the License definition
type License struct {
	ID            string `json:"id,omitempty"`            // License ID.
	Name          string `json:"name,omitempty"`          // License Display Name.
	TotalUnits    string `json:"totalUnits,omitempty"`    // License quantity total.
	ConsumedUnits string `json:"consumedUnits,omitempty"` // License quantity consumed.
}
