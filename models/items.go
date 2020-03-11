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

// Delete an item by id
func DeleteItem(itemId string) error {
	_, err := db.Exec("DELETE FROM items WHERE id = ?", itemId)
	if err != nil {
		return err
	}

	return nil
}

// Update an item
func UpdateItem(item *Item) error {
	_, err := db.Exec("UPDATE items SET title = ?, description = ?, price_per_unit = ?, unit_size = ?, image = ? WHERE id = ?", item.Title, item.Description, item.PricePerUnit, item.UnitSize, item.Image, item.Id)

	if err != nil {
		return err
	}

	return nil
}
