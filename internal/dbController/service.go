package dbController

import (
	_ "embed"
	"fmt"
	"github.com/Q1KLLI/oilProductionAI/internal/logic"
)

//go:embed query/wells.sql
var queryWells string

func (d *Database) WellIDs() ([]int, error) {

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

func (d *Database) Wells() ([]*logic.Well, error) {

	rows, err := d.db.Query(queryWells)
	if err != nil {
		return nil, fmt.Errorf("unable query wells: %w", err)
	}

	var wells []*logic.Well

	for rows.Next() {

		well := new(logic.Well)

		err = rows.Scan(
			&well.ID,
			&well.Name,
			&well.State,
		)

		wells = append(wells, well)
		if err != nil {
			return nil, fmt.Errorf("unable scan wells: %w", err)
		}
	}
	return wells, nil
}
