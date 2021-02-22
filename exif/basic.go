package ImgMeta

// BasicInfo contains the most basic information that could be asked for
type BasicInfo struct {
	//Width    interface{}
	//Height   uint32
	//Title    string
	Description string
	Keywords    []string
}

// GetBasicInfo gets the basic information from the meta-information of the image
func GetBasicInfo(img Image) (info BasicInfo) {
	//width, err := img.ReadTagValue("SOF0", SOF0ImageWidth)
	//if err == nil {
	//	info.Width = width
	//} else {
	//	fmt.Println(err.Error())
	//}
	//height, err := img.ReadTagValue("SOF0", SOF0ImageHeight)
	//if err == nil {
	//	info.Height = height.(uint32)
	//} else {
	//	fmt.Println(err.Error())
	//}
	iptc, err := img.ReadTagValue("IPTC", IptcTagApplication2Keywords)
	if err == nil {
		iptcMap := iptc.(map[uint16]interface{})
		if value, ok := iptcMap[IptcTagApplication2Caption]; ok {
			info.Description = value.([]string)[0]
		}
		if value, ok := iptcMap[IptcTagApplication2Keywords]; ok {
			info.Keywords = value.([]string)
		}
	}
	//datetime, err := img.ReadTagValue("EXIF", ExifTagDateTimeOriginal)
	//if err == nil {
	//	fmt.Printf("datetime:%v\n", datetime)
	//}
	//height, err := img.ReadTagValue("IPTC", IptcTagApplication2Keywords)
	//if err == nil {
	//	info.Height = height.(float64)
	//} else {
	//	fmt.Println(err.Error())
	//}
	return
}
