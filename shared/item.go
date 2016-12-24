package shared

type Item struct {
	ID               int     `db:"id"`
	SKU              string  `db:"sku"`
	Price            float64 `db:"price"`
	Name             string  `db:"name"`
	Descr            string  `db:"descr"`
	Image            string  `db:"image"`
	Qty              int     `db:"qty"`
	VolumeML         int     `db:"volume_ml"`
	WeightG          int     `db:"weight_g"`
	ShippingVolumeML int     `db:"shipping_volume_ml"`
	ShippingWeightG  int     `db:"shipping_weight_g"`
}
