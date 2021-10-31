package status

// Does a SELECT 1 to check if database is alive and accepting queries
func (s *StatusTransaction) Ping() (err error) {
	_, err = s.tx.Exec(`SELECT 1`)
	return
}
