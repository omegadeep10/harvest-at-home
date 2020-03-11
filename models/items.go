package models

// Get all Items as a list
func GetAllItems() ([]*Item, error) {
	rows, err := db.Query("SELECT * FROM Items")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := make([]*Item, 0)
	for rows.Next() {
		item := new(Item)
		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.PricePerUnit, &item.UnitSize, &item.Image)

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

// Create a new item
func CreateItem(item *Item) error {
	_, err := db.Exec("INSERT INTO items(title, description, price_per_unit, unit_size, image) VALUES(?, ?, ?, ?, ?)", item.Title, item.Description, item.PricePerUnit, item.UnitSize, item.Image)
	if err != nil {
		return err
	}

	return nil
}
