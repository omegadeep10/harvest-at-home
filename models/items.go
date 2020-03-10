package models

type Item struct {
	id             int
	title          string
	description    string
	price_per_unit float32
	unit_size      string
	image          string
}

func AllItems() ([]*Item, error) {
	rows, err := db.Query("SELECT * FROM Items")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := make([]*Item, 0)
	for rows.Next() {
		item := new(Item)
		err := rows.Scan(&item.id, &item.title, &item.description, &item.price_per_unit, &item.unit_size, &item.image)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func InsertItem(item *Item) error {
	_, err := db.Exec("INSERT INTO items VALUES(?, ?, ?, ?, ?, ?, ?)", item.id, item.title, item.description, item.price_per_unit, item.unit_size, item.image)
	if err != nil {
		return err
	}

	return nil
}
