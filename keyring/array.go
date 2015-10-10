package keyring

type ArrayKeyring struct {
	items map[string]Item
}

func (k *ArrayKeyring) Get(key string) (Item, error) {
	if i, ok := k.items[key]; ok {
		return i, nil
	} else {
		return Item{}, ErrKeyNotFound
	}
}

func (k *ArrayKeyring) Set(i Item) error {
	if k.items == nil {
		k.items = map[string]Item{}
	}
	k.items[i.Key] = i
	return nil
}

func (k *ArrayKeyring) Remove(key string) error {
	delete(k.items, key)
	return nil
}

func (k *ArrayKeyring) Keys() ([]string, error) {
	var keys = []string{}
	for key := range k.items {
		keys = append(keys, key)
	}
	return keys, nil
}
