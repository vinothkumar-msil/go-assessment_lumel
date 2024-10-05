package calculations

import (
	"database/sql"
	"go-backend-assessment/db"
)

// CalculateRevenue calculates the total revenue between the given date range
func CalculateRevenue(startDate, endDate string) (float64, error) {
	var revenue float64

	// Updated query to join the sales and orders tables
	query := `SELECT SUM(s.quantity_sold * s.unit_price - s.discount + sh.shipping_cost) AS total_revenue
              FROM sales s
              JOIN orders o ON s.order_id = o.order_id
              JOIN shipping sh ON o.shipping_id = sh.shipping_id
              WHERE o.order_date BETWEEN $1 AND $2`

	err := db.DB.QueryRow(query, startDate, endDate).Scan(&revenue)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // Return 0 revenue if no rows found
		}
		return 0, err
	}
	return revenue, nil
}
