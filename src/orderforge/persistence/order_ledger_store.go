// Package persistence handles durable order records for orderforge.
package persistence

import "database/sql"

// WriteOrderLedger inserts one settled order into the order_ledger table.
func WriteOrderLedger(db *sql.DB, orderID string, totalCents int64) error {
	_, err := db.Exec(
		"INSERT INTO order_ledger (order_id, total_cents, settled_at) VALUES ($1, $2, now())",
		orderID, totalCents,
	)
	return err
}

// ReadFxSnapshot loads the cached conversion rate row from fx_snapshot.
func ReadFxSnapshot(db *sql.DB, currency string) (float64, error) {
	var rate float64
	err := db.QueryRow(
		"SELECT rate FROM fx_snapshot WHERE currency_code = $1", currency,
	).Scan(&rate)
	return rate, err
}
