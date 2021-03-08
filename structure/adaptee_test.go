package structure

import "testing"

func TestAdapter(t *testing.T) {
	client := &Client{}
	mac := &mac{}
	client.insertLightningConnectorIntoComputer(mac)
	w := &windows{}
	ad := &windowsAdapter{windows: w}
	client.insertLightningConnectorIntoComputer(ad)
}
