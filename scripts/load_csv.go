package scripts

import (
	"encoding/csv"
	"fmt"
	"go-backend-assessment/db"
	"os"
	"strconv"
	"strings"
	"time"
)

const batchSize = 10 // Number of rows to insert in a single batch

// LoadCSVData loads the CSV file and inserts data into the database using batch insert
func LoadCSVData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Define slices to hold data for batch inserts
	var customerData []string
	var productData []string
	var regionData []string
	var shippingData []string
	var paymentData []string
	var orderData []string
	var salesData []string

	// Iterate over the rows in the CSV file
	for i, row := range records {
		// Customer data
		customerID := row[2]
		customerName := row[12]
		customerEmail := row[13]
		customerAddress := row[14]
		customerData = append(customerData, fmt.Sprintf("('%s', '%s', '%s', '%s')", customerID, customerName, customerEmail, customerAddress))

		// Product data
		productID := row[1]
		productName := row[3]
		category := row[4]
		productData = append(productData, fmt.Sprintf("('%s', '%s', '%s')", productID, productName, category))

		// Region data
		region := row[5]
		regionData = append(regionData, fmt.Sprintf("('%s')", region))

		// Shipping data
		shippingCost, _ := strconv.ParseFloat(row[10], 64)
		shippingData = append(shippingData, fmt.Sprintf("((SELECT region_id FROM regions WHERE region_name = '%s'), %f)", region, shippingCost))

		// Payment data
		paymentMethod := row[11]
		paymentData = append(paymentData, fmt.Sprintf("('%s')", paymentMethod))

		// Order data
		orderID := row[0]
		dateOfSale := row[6]
		orderDate, _ := time.Parse("2006-01-02", dateOfSale)
		orderData = append(orderData, fmt.Sprintf("('%s', '%s', '%s', (SELECT shipping_id FROM shipping WHERE shipping_cost = %f), (SELECT payment_id FROM payments WHERE payment_method = '%s'))", orderID, customerID, orderDate.Format("2006-01-02"), shippingCost, paymentMethod))

		// Sales data
		quantitySold, _ := strconv.Atoi(row[7])
		unitPrice, _ := strconv.ParseFloat(row[8], 64)
		discount, _ := strconv.ParseFloat(row[9], 64)
		salesData = append(salesData, fmt.Sprintf("('%s', '%s', %d, %f, %f)", orderID, productID, quantitySold, unitPrice, discount))

		// Execute batch insert every batchSize rows
		if (i+1)%batchSize == 0 || i == len(records)-1 {
			fmt.Println(customerData, productData, regionData, shippingData, paymentData, orderData, salesData)
			// err := batchInsert(customerData, productData, regionData, shippingData, paymentData, orderData, salesData)
			// if err != nil {
			// 	return err
			// }

			// Clear the data slices after each batch insert
			customerData = customerData[:0]
			productData = productData[:0]
			regionData = regionData[:0]
			shippingData = shippingData[:0]
			paymentData = paymentData[:0]
			orderData = orderData[:0]
			salesData = salesData[:0]
		}
	}

	return nil
}

// Helper function to perform batch insert
func batchInsert(customers, products, regions, shippings, payments, orders, sales []string) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	// Customer batch insert
	if len(customers) > 0 {
		customerQuery := `INSERT INTO customers (customer_id, customer_name, customer_email, customer_address) VALUES ` + strings.Join(customers, ",") +
			` ON CONFLICT (customer_id) DO NOTHING`
		_, err := tx.Exec(customerQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Product batch insert
	if len(products) > 0 {
		productQuery := `INSERT INTO products (product_id, product_name, category) VALUES ` + strings.Join(products, ",") +
			` ON CONFLICT (product_id) DO NOTHING`
		_, err := tx.Exec(productQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Region batch insert
	if len(regions) > 0 {
		regionQuery := `INSERT INTO regions (region_name) VALUES ` + strings.Join(regions, ",") +
			` ON CONFLICT (region_name) DO NOTHING`
		_, err := tx.Exec(regionQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Shipping batch insert
	if len(shippings) > 0 {
		shippingQuery := `INSERT INTO shipping (region_id, shipping_cost) VALUES ` + strings.Join(shippings, ",") +
			` ON CONFLICT DO NOTHING`
		_, err := tx.Exec(shippingQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Payment batch insert
	if len(payments) > 0 {
		paymentQuery := `INSERT INTO payments (payment_method) VALUES ` + strings.Join(payments, ",") +
			` ON CONFLICT (payment_method) DO NOTHING`
		_, err := tx.Exec(paymentQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Order batch insert
	if len(orders) > 0 {
		orderQuery := `INSERT INTO orders (order_id, customer_id, order_date, shipping_id, payment_id) VALUES ` + strings.Join(orders, ",") +
			` ON CONFLICT (order_id) DO NOTHING`
		_, err := tx.Exec(orderQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Sales batch insert
	if len(sales) > 0 {
		salesQuery := `INSERT INTO sales (order_id, product_id, quantity_sold, unit_price, discount) VALUES ` + strings.Join(sales, ",") +
			` ON CONFLICT DO NOTHING`
		_, err := tx.Exec(salesQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
