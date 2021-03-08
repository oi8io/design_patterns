package structure

import "fmt"

type Computer interface {
	insertIntoLightningPort()
}

type Client struct {

}

func (c *Client)insertLightningConnectorIntoComputer(com Computer)  {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

type mac struct {
	
}

func (c *mac)insertIntoLightningPort()  {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type windows struct {

}
func (c *windows)insertIntoUSBPort()  {
	fmt.Println("USB connector is plugged into windows machine.")
}

type windowsAdapter struct {
	windows *windows
}

func (w windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windows.insertIntoUSBPort()
}

