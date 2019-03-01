package regions

import (
	"encoding/json"
	"errors"
)

func List() (Regions, error) {
	// box, err := rice.FindBox("../")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := box.Bytes("regions.json")
	// if err != nil {
	// 	return nil, err
	// }

	// bytes, err := ioutil.ReadAll(data)
	// if err != nil {
	// 	return nil, err
	// }

	var regions Regions
	json.Unmarshal([]byte(REGION_DATA), &regions)

	return regions, nil
}

func ListPublic() (Regions, error) {
	regions, err := List()
	if err != nil {
		return nil, err
	}

	filtered := make([]Region, 0)
	for _, v := range regions {
		if v.Public {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}

func LookupByCode(code string) (Region, error) {
	return lookup(func(region Region) bool {
		return region.Code == code
	})
}

func LookupByName(name string) (Region, error) {
	return lookup(func(region Region) bool {
		return region.Name == name
	})
}

func lookup(filter func(Region) bool) (Region, error) {
	regions, err := List()
	if err != nil {
		return Region{}, err
	}

	for _, v := range regions {
		if filter(v) {
			return v, nil
		}
	}

	return Region{}, errors.New("Could not find region")
}
