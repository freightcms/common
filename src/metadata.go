package models

// Metadata represents metadata information
type Metadata struct {
	OrderNumber         *string
	InvoiceNumber       *string
	PurchaseOrderNumber *string
	CustomerOrderNumber *string
	BillOfLading        *string
	LineItemNumber      *string
	ReferenceNumber     *string
	ContainerNumber     *string
	ShipmentNumber      *string
	SealNumber          *string
	BookingNumber       *string
	MasterBillOfLading  *string
	HouseBillOfLading   *string
	ProNumber           *string
	TrackingNumber      *string
	PackageNumber       *string
}
