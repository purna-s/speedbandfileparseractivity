package speedbandfileparseractivity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-speedbandfileparseractivity")

// MyActivity is a stub for your Activity implementation
type XMLParserActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XMLParserActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *XMLParserActivity) Metadata() *activity.Metadata {
	return a.metadata
}

//XSD

type TrafficInfo struct {
	XMLName            xml.Name         `xml:"TrafficInfo" json:"-"`
	SpeedBandInfoList  []SpeedBandInfo  `xml:"SpeedBandInfo" json:"SpeedBandInfo"`
	TDataList          []TData          `xml:"TData" json:"TData"`
}

type SpeedBandInfo struct {
	XMLName        xml.Name `xml:"SpeedBandInfo" json:"-"`
	Band           string   `xml:"Band" json:"Band"`
	MinimumSpeed   string   `xml:"MinimumSpeed" json:"MinimumSpeed"`
	MaximumSpeed   string   `xml:"MaximumSpeed" json:"MaximumSpeed"`
}

type TData struct {
	XMLName     xml.Name `xml:"TData" json:"-"`
	LinkID      string   `xml:"LinkID" json:"LinkID"`
	SpeedBand   string   `xml:"SpeedBand" json:"SpeedBand"`
}

// end of XSD

// Eval implements activity.Activity.Eval
func (a *XMLParserActivity) Eval(ctx activity.Context) (done bool, err error) {
	File := ctx.GetInput("file").(string)
	//XMLString := ctx.GetInput("xmlString").(string)

	activityLog.Debugf("File is : [%s]", File)
	//	activityLog.Debugf("XML String is : [%s]", XMLString)

	fmt.Println("File is : ", File)
	//fmt.Println("XML String is : ", XMLString)

	if len(File) == 0 { //&& (len(XMLString) == 0)

		activityLog.Debugf("value in both the fields is empty at least give one input ")
		fmt.Println("value in both the fields is empty at least give one input")

	}

	xmlFile, err := os.Open(File)
	// if we os.Open returns an error then handle it
	if err != nil {
		activityLog.Debugf("File Exception :  ", err)
		fmt.Println("File Exception ", err)
		return
	}

	fmt.Println("Successfully Opened ", File)
	activityLog.Debugf("Successfully Opened ", File)
	// defer the closing of our xmlFile so that we can parse it later on

	defer xmlFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		activityLog.Debugf("File Read Exception ", err)
		fmt.Println("File Read Exception ", err)
		return
	}

	// For File
	xmldata := SpeedBandInfo{}
	err = xml.Unmarshal(byteValue, &xmldata)
	jsonData, _ := json.Marshal(xmldata)
	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}

	//fmt.Println(string(jsonData)) // Printing Json Data

	//for XML String
	//xmlString := (string(XMLString))
	/*{	xml_data := SpeedBandInfo{}
		err = xml.Unmarshal([]byte(XMLString), &xml_data)
		jsonData, _ = json.Marshal(xml_data)
		if err != nil {
			activityLog.Debugf("Error ", err)
			fmt.Println("error: ", err)

		}
		fmt.Println("InSide XML String condition")
		fmt.Println(string(jsonData))
	}*/
	// Set the output as part of the context
	activityLog.Debugf("Activity has parsed SpeedBand file content Successfully")
	fmt.Println("Activity has parsed SpeedBand file content Successfully")

	ctx.SetOutput("output", string(jsonData))

	return true, nil
}
