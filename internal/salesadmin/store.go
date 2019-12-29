package salesadmin

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	// _ postgres driver for the database/sql package
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Order is a struct that models the saledata csv and
// orders table data format
type Order struct {
	OrderID         int32
	CustomerName    string
	ItemDescription string
	ItemPrice       float64
	Quantity        int64
	MerchantName    string
	MerchantAddress string
}

const insertOrder = `
INSERT INTO orders (customer_name, item_description, item_price, quantity, merchant_name, merchant_address)
VALUES ($1, $2, $3, $4, $5, $6)`

const queryAllOrders = `SELECT order_id, customer_name, item_description, item_price, quantity, merchant_name, merchant_address FROM orders;`

const queryRevenueTotal = `SELECT item_price, quantity FROM orders;`

const queryCustomerNames = `SELECT customer_name FROM orders;`

const queryMerchantNames = `SELECT merchant_name FROM orders;`

// ParseSalesFile takes the sales data csv slice of byte slices
// and unmarshals it into a slice of Order structs
func ParseSalesFile(ctx context.Context, file [][]byte) ([]Order, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("parsing file bytes and unmarshalling into Order structs")
	var orders []Order
	for i, row := range file {
		if i == 0 {
			continue
		}

		fileRowString := string(row)
		csvReader := csv.NewReader(strings.NewReader(fileRowString))
		record, err := csvReader.Read()
		if err != nil {
			csvReaderErr := fmt.Errorf("error %v while reading row %v", err, i)
			return nil, csvReaderErr
		}

		itemPrice, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			parseFloatError := fmt.Errorf("error %v while converting item price to float64 in row %v", err, i)
			return nil, parseFloatError
		}

		itemQuantity, err := strconv.ParseInt(record[3], 2, 16)
		if err != nil {
			parseIntError := fmt.Errorf("error %v while converting item quantity to int in row %v", err, i)
			return nil, parseIntError
		}

		tempOrder := Order{
			CustomerName:    record[0],
			ItemDescription: record[1],
			ItemPrice:       itemPrice,
			Quantity:        itemQuantity,
			MerchantName:    record[4],
			MerchantAddress: record[5],
		}

		orders = append(orders, tempOrder)
	}
	return orders, nil
}

// InsertOrders executes the insertOder sql statement for each order in the order slice
// and will return an error if unsuccessful
func InsertOrders(ctx context.Context, orders []Order) error {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslkey=%s sslcert=%s sslrootcert=%s sslmode=verify-ca",
		viper.GetString("postgresHost"), viper.GetInt("postgresPort"), viper.GetString("postgresUser"), viper.GetString("postgresPassword"),
		viper.GetString("postgresDB"), viper.GetString("postgresKeyFile"), viper.GetString("postgresCertFile"), viper.GetString("caFile"))

	sugar.Info("connecting to postgres salesadmin database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		dbConnectionError := fmt.Errorf("error %v while opening db connection", err)
		return dbConnectionError
	}
	defer db.Close()

	sugar.Info("inserting orders into salesadming database")
	for i, order := range orders {
		_, err := db.ExecContext(ctx, insertOrder,
			order.CustomerName, order.ItemDescription, order.ItemPrice,
			order.Quantity, order.MerchantName, order.MerchantAddress)
		if err != nil {
			insertOrderError := fmt.Errorf("error %v while inserting order %v", err, i)
			return insertOrderError
		}
	}

	return nil
}

// RunQueryAllOrders executes the provided query string
func RunQueryAllOrders(ctx context.Context, query string) ([]Order, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslkey=%s sslcert=%s sslrootcert=%s sslmode=verify-ca",
		viper.GetString("postgresHost"), viper.GetInt("postgresPort"), viper.GetString("postgresUser"), viper.GetString("postgresPassword"),
		viper.GetString("postgresDB"), viper.GetString("postgresKeyFile"), viper.GetString("postgresCertFile"), viper.GetString("caFile"))

	sugar.Info("connecting to database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		dbConnectionError := fmt.Errorf("error %v while opening db connection", err)
		return nil, dbConnectionError
	}
	defer db.Close()

	sugar.Info("querying salesadmin database for all orders")
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		queryAllOrdersError := fmt.Errorf("error %v recevied while querying database", err)
		return nil, queryAllOrdersError
	}

	sugar.Info("scanning result rows and unmarshalling into Order structs")
	var order Order
	var orders []Order
	for rows.Next() {
		err = rows.Scan(&order.OrderID, &order.CustomerName,
			&order.ItemDescription, &order.ItemPrice, &order.Quantity,
			&order.MerchantName, &order.MerchantAddress)
		if err != nil {
			scanRowError := fmt.Errorf("error scan row %v", err)
			return nil, scanRowError
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return orders, nil
}

// RunQueryTotalRevenue executes the provided query string
func RunQueryTotalRevenue(ctx context.Context, query string) ([]Order, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslkey=%s sslcert=%s sslrootcert=%s sslmode=verify-ca",
		viper.GetString("postgresHost"), viper.GetInt("postgresPort"), viper.GetString("postgresUser"), viper.GetString("postgresPassword"),
		viper.GetString("postgresDB"), viper.GetString("postgresKeyFile"), viper.GetString("postgresCertFile"), viper.GetString("caFile"))

	sugar.Info("connecting to postgres salesadmin database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		dbConnectionError := fmt.Errorf("error %v while opening db connection", err)
		return nil, dbConnectionError
	}
	defer db.Close()

	sugar.Info("querying database for all item prices and quantities from each order")
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		queryAllOrdersError := fmt.Errorf("error %v recevied while querying database", err)
		return nil, queryAllOrdersError
	}

	sugar.Info("scanning result rows and unmarshalling into Order structs")
	var order Order
	var orders []Order
	for rows.Next() {
		err = rows.Scan(&order.ItemPrice, &order.Quantity)
		if err != nil {
			scanRowError := fmt.Errorf("error scan row %v", err)
			return nil, scanRowError
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return orders, nil
}

// RunQueryCustomerNames executes the provided query string
func RunQueryCustomerNames(ctx context.Context, query string) ([]Order, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslkey=%s sslcert=%s sslrootcert=%s sslmode=verify-ca",
		viper.GetString("postgresHost"), viper.GetInt("postgresPort"), viper.GetString("postgresUser"), viper.GetString("postgresPassword"),
		viper.GetString("postgresDB"), viper.GetString("postgresKeyFile"), viper.GetString("postgresCertFile"), viper.GetString("caFile"))

	sugar.Info("connecting to postgres salesadmin database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		dbConnectionError := fmt.Errorf("error %v while opening db connection", err)
		return nil, dbConnectionError
	}
	defer db.Close()

	sugar.Info("querying database for all customer names")
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		queryAllOrdersError := fmt.Errorf("error %v recevied while querying database", err)
		return nil, queryAllOrdersError
	}

	sugar.Info("scanning result rows and unmarshalling into Order structs")
	var order Order
	var orders []Order
	for rows.Next() {
		err = rows.Scan(&order.CustomerName)
		if err != nil {
			scanRowError := fmt.Errorf("error scan row %v", err)
			return nil, scanRowError
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return orders, nil
}

// RunQueryMerchantNames executes the provided query string
func RunQueryMerchantNames(ctx context.Context, query string) ([]Order, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslkey=%s sslcert=%s sslrootcert=%s sslmode=verify-ca",
		viper.GetString("postgresHost"), viper.GetInt("postgresPort"), viper.GetString("postgresUser"), viper.GetString("postgresPassword"),
		viper.GetString("postgresDB"), viper.GetString("postgresKeyFile"), viper.GetString("postgresCertFile"), viper.GetString("caFile"))

	sugar.Info("connecting to postgres salesadmin database")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		dbConnectionError := fmt.Errorf("error %v while opening db connection", err)
		return nil, dbConnectionError
	}
	defer db.Close()

	sugar.Info("querying database for all merchant names")
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		queryAllOrdersError := fmt.Errorf("error %v recevied while querying database", err)
		return nil, queryAllOrdersError
	}

	sugar.Info("scanning result rows and unmarshalling into Order structs")
	var order Order
	var orders []Order
	for rows.Next() {
		err = rows.Scan(&order.MerchantName)
		if err != nil {
			scanRowError := fmt.Errorf("error scan row %v", err)
			return nil, scanRowError
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return orders, nil
}
