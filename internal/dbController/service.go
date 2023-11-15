package dbController

func (d database) WellIDs() ([]int, error) {

	var s []int
	var id int

	rows, err := d.db.Query("select id from well")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		s = append(s, id)
	}
	return s, nil
}
