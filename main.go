package main

import (
	"bytes"
	"fmt"
	"os"
	"protobf-example/models"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
)

func main() {
	var user1 = &models.User{
		Id:       "u001",
		Name:     "Mulia Ichsan",
		Password: "f0r Th3 Mein",
		Gender:   models.UserGender_MALE,
	}

	// =========== original
	fmt.Printf("# ==== Original Proto Object\n       %#v \n", user1)

	// =========== as string
	fmt.Printf("# ==== Proto Object as String\n       %v \n", user1.String())

	var garage1 = &models.Garage{
		Id:   "g001",
		Name: "Texas",
		Coordinate: &models.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &models.GarageList{
		List: []*models.Garage{
			garage1,
		},
	}

	// converting proto object as JSON
	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonString := buf.String()
	fmt.Printf("# ==== Proto Object As JSON String\n       %v \n", jsonString)

	// converting JSON as String
	buf2 := strings.NewReader(jsonString)
	protoObject := new(models.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== JSON to Proto Object As String\n       %v \n", protoObject.String())

	var garageListByUser = &models.GarageListByUser{
		List: map[string]*models.GarageList{
			user1.Id: garageList,
		},
	}

	// converting garageListByUser Proto Object as JSOn
	var buf3 bytes.Buffer
	err3 := (&jsonpb.Marshaler{}).Marshal(&buf3, garageListByUser)
	if err1 != nil {
		fmt.Println(err3.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== Proto Object As JSON String\n       %v \n", buf3.String())

	// converting JSON object to Proto Object as String
	protoObject2 := new(models.GarageListByUser)
	err4 := jsonpb.UnmarshalString(buf3.String(), protoObject2)
	if err1 != nil {
		fmt.Println(err4.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== JSON Object to Proto As String\n       %v \n", protoObject2.String())
}
