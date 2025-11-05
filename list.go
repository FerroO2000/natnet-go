package natnetgo

type parseItemFn[T any] func(data []byte) (T, int, error)

type List[T any] struct {
	Count int32
	Items []T
}

const listMinLen = int32Len

func parseList[T any](parseItem parseItemFn[T], data []byte) (*List[T], int, error) {
	if len(data) < listMinLen {
		return nil, 0, ErrTooShort
	}

	l := new(List[T])

	count, _ := parseInt32(data)
	l.Count = count

	offset := listMinLen
	l.Items = make([]T, 0, l.Count)
	for range l.Count {
		item, tmpOffset, err := parseItem(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		l.Items = append(l.Items, item)
		offset += tmpOffset
	}

	return l, offset, nil
}

type SizedList[T any] struct {
	Count    int32
	DataSize int32
	Items    []T
}

const sizedListMinLen = 2 * listMinLen

func parseSizedList[T any](parseItem parseItemFn[T], data []byte) (*SizedList[T], int, error) {
	if len(data) < sizedListMinLen {
		return nil, 0, ErrTooShort
	}

	sl := new(SizedList[T])

	count, _ := parseInt32(data)
	sl.Count = count

	dataSize, _ := parseInt32(data[int32Len:])
	sl.DataSize = dataSize

	offset := sizedListMinLen
	sl.Items = make([]T, 0, sl.Count)
	for range sl.Count {
		item, tmpOffset, err := parseItem(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		sl.Items = append(sl.Items, item)
		offset += tmpOffset
	}

	return sl, offset, nil
}
