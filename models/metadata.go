package models

// Metadata represents metadata information
type Metadata struct {
	OrderNumber         string `json:"orderNumber" bson:"orderNumber"`
	InvoiceNumber       string `json:"invoiceNumber" bson:"invoiceNumber"`
	PurchaseOrderNumber string `json:"purchaseOrderNumber" bson:"purchaseOrderNumber"`
	CustomerOrderNumber string `json:"customerOrderNumber" bson:"customerOrderNumber"`
	BillOfLading        string `json:"billOfLading" bson:"billOfLading"`
	LineItemNumber      string `json:"lineItemNumber" bson:"lineItemNumber"`
	ReferenceNumber     string `json:"referenceNumber" bson:"referenceNumber"`
	ContainerNumber     string `json:"containerNumber" bson:"containerNumber"`
	ShipmentNumber      string `json:"shipmentNumber" bson:"shipmentNumber"`
	SealNumber          string `json:"sealNumber" bson:"sealNumber"`
	BookingNumber       string `json:"bookingNumber" bson:"bookingNumber"`
	MasterBillOfLading  string `json:"masterBillOfLading" bson:"masterBillOfLading"`
	HouseBillOfLading   string `json:"houseBillOfLading" bson:"houseBillOfLading"`
	ProNumber           string `json:"proNumber" bson:"proNumber"`
	TrackingNumber      string `json:"trackingNumber" bson:"trackingNumber"`
	PackageNumber       string `json:"packageNumber" bson:"packageNumber"`
}
