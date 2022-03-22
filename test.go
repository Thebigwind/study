package main

//LuLevelInfo  房源信息
type LuLevelInfo struct {
	LuID                    interface{} `json:"luId"` //int64
	LodgeBusiness           string      `json:"lodgeBusiness"`
	LandlordID              string      `json:"landlordId"` //string
	LodgeUnitName           string      `json:"lodgeUnitName"`
	MainImageURL            string      `json:"mainImageUrl"`
	LeaseType               string      `json:"leaseType"`
	LeaseTypeEn             string      `json:"leaseTypeEn"`
	MaxDays                 interface{} `json:"maxDays"`          //string
	MinDays                 interface{} `json:"minDays"`          //string
	HouseTypeRoom           interface{} `json:"houseTypeRoom"`    //int
	HouseTypeHall           interface{} `json:"houseTypeHall"`    //int
	HouseTypeToilet         interface{} `json:"houseTypeToilet"`  //int
	HouseTypeKitchen        interface{} `json:"houseTypeKitchen"` //int
	HouseTypeBalcony        interface{} `json:"houseTypeBalcony"` //int
	GuestNum                interface{} `json:"guestNum"`         //int
	State                   string      `json:"state"`
	RoomsNum                interface{} `json:"roomsNum"`   //int
	Crvp                    interface{} `json:"crvp"`       //int
	LuOtherFee              interface{} `json:"luOtherFee"` //string
	CreateType              string      `json:"createType"`
	Currency                string      `json:"currency"`
	Cashpledge              interface{} `json:"cashpledge"` //string
	CleanFee                interface{} `json:"cleanFee"`   //int
	PrePayRate              interface{} `json:"prePayRate"` //int
	IsOverSea               bool        `json:"isOverSea"`
	IsGangAoTai             bool        `json:"isGangAoTai"`
	BookFlow                string      `json:"bookFlow"`
	CancelRule              interface{} `json:"cancelRule"`
	SmartLockId             interface{} `json:"smartLockId"` //string
	Area                    interface{} `json:"area"`        //int
	CashPledgeOnline        string      `json:"cashPledgeOnline"`
	GuestGender             interface{} `json:"guestGender"`         //int
	HasLicense              interface{} `json:"hasLicense"`          //int
	BuildingProductType     interface{} `json:"buildingProductType"` //int
	BuildingProductTypeName string      `json:"buildingProductTypeName"`
	ActivitiesAllowed       string      `json:"activitiesAllowed"`
	ProductType             interface{} `json:"productType"` //int
	Attributes              interface{} `json:"attributes"`  //int
	LuRuzhuTips             string      `json:"luRuzhuTips"` //string

	CheckinTime    string     `json:"checkinTime"`    //
	CheckoutTime   string     `json:"checkoutTime"`   //
	CreateTime     string     `json:"createTime"`     //
	BusinessCircle string     `json:"businessCircle"` //
	IsSupportIm    bool       `json:"isSupportIm"`    //
	AddressInfo    AddresInfo `json:"addressInfo"`
	ExtraInfo      ExtraInfo  `json:"extraInfo"`
	BedInfo        BInfo      `json:"bedInfo"`
	PriceInfo      PriceInfo  `json:"priceInfo"`
	MultiPrice     MultiPrice `json:"multiPrice"`
}

///////////////// bedInfo /////////////
type BInfo struct {
	BedInfo        []BedInfo `json:"bedInfo"`
	BedNum         string    `json:"bedNum"`
	BedGuestTip    string    `json:"bedGuestTip"`
	BedGuestTipNew string    `json:"bedGuestTipNew"`
}
type BedInfo struct {
	ID     interface{} `json:"id"` //int
	TypeEn string      `json:"typeEn"`
	Type   string      `json:"type"`
	Width  interface{} `json:"width"`  //string
	Length interface{} `json:"length"` //string
	Nums   interface{} `json:"nums"`   //int
}

////////////// priceInfo //////////
type PriceInfo struct {
	LuID                interface{} `json:"luId"` //string
	UserIsVip           bool        `json:"userIsVip"`
	TotalCalendarPrice  interface{} `json:"totalCalendarPrice"`  //string
	AvgCalendarPrice    interface{} `json:"avgCalendarPrice"`    //string
	TotalPromotionPrice interface{} `json:"totalPromotionPrice"` //string
	AvgPromotionPrice   interface{} `json:"avgPromotionPrice"`   //string
	//Detail              Detail `json:"detail"`
}

//AddresInfo 房源地址信息
type AddresInfo struct {
	NationId           interface{}   `json:"nationId"`   //string
	ProvinceId         interface{}   `json:"provinceId"` //string
	CityId             interface{}   `json:"cityId"`     //string
	StreetId           interface{}   `json:"streetId"`   //string
	Latitude           string        `json:"latitude"`
	Longitude          string        `json:"longitude"`
	DisplayAddress     string        `json:"displayAddress"`
	DisplayAddress2    string        `json:"displayAddress2"`
	AddressCombination string        `json:"addressCombination"`
	DoorNumber         interface{}   `json:"doorNumber"` //string
	NationName         string        `json:"nationName"`
	ProvinceName       string        `json:"provinceName"`
	CityName           string        `json:"cityName"`
	Timezone           string        `json:"timezone"`
	DistrictName       string        `json:"districtName"`
	StreetName         string        `json:"streetName"`
	NationInfo         NationInfo    `json:"nationInfo"`
	ProvinceInfo       ProvinceInfo  `json:"provinceInfo"`
	CityInfo           LodgeCityInfo `json:"cityInfo"`
	DistrictInfo       DistrictInfo  `json:"districtInfo"`
	DetailAddress      string        `json:"DetailAddress"`
}

type ExtraInfo struct {
	AverageCommentScore    interface{} `json:"averageCommentScore"`
	AverageCommentScoreAll interface{} `json:"averageCommentScoreAll"` //int
	Comments               interface{} `json:"comments"`               //int
	ComeFrom               string      `json:"comeFrom"`
	FirstOnlineTime        string      `json:"firstOnlineTime"`
	LodgeUnitFromType      string      `json:"lodgeUnitFromType"`
	Telephone              interface{} `json:"telephone"`
	BdCheckedReal          interface{} `json:"bdCheckedReal"`  //int
	BdCheckedPhoto         interface{} `json:"bdCheckedPhoto"` //int
	Images                 interface{} `json:"images"`         //int
	Channel                interface{} `json:"channel"`        //int
	PartnerId              interface{} `json:"partnerId"`      //int
}

//NationInfo 国家信息
type NationInfo struct {
	Id             interface{} `json:"id"` //string
	Enname         string      `json:"enname"`
	Name           string      `json:"name"`
	Showname       string      `json:"showname"`
	Mobileareacode string      `json:"mobileareacode"`
	Shortname      string      `json:"shortname"`
	ContinentId    interface{} `json:"continent_id"` //string
	Py             string      `json:"py"`
}

//ProvinceInfo 省份信息
type ProvinceInfo struct {
	ProvinceId   interface{} `json:"province_id"` //string
	ProvinceName string      `json:"province_name"`
	Name         string      `json:"name"`
	ShortName    string      `json:"short_name"`
	Pinyin       string      `json:"pinyin"`
	Shortname    string      `json:"shortname"`
	NationId     interface{} `json:"nation_id"` //string
	StandardCode string      `json:"standard_code"`
}

//DistrictInfo 区域信息
type DistrictInfo struct {
	DistrictId   interface{} `json:"district_id"` //string
	DistrictName string      `json:"district_name"`
	ShortName    string      `json:"short_name"`
	Pinyin       string      `json:"pinyin"`
	Shortname    string      `json:"shortname"`
	CityId       interface{} `json:"city_id"`       //string
	StandardCode interface{} `json:"standard_code"` //string
}

type LodgeCityInfo struct {
	CityID       interface{} `json:"city_id"` // string
	CityName     string      `json:"city_name"`
	ShortName    string      `json:"short_name"`
	Pinyin       string      `json:"pinyin"`
	ShortPinyin  string      `json:"short_pinyin"`
	StandardCode string      `json:"standard_code"`
	ProvinceID   interface{} `json:"province_id"` // string
	Timezone     string      `json:"timezone"`
}

//detailInfo

//cancelRule

//imageInfo

//multiPrice
type MultiPrice struct {
	HolidayPrice      interface{} `json:"holidayPrice"`
	HolidayPriceH5Url string      `json:"holidayPriceH5Url"`
	Price             interface{} `json:"price"`
	WeekendPrice      interface{} `json:"weekendPrice"`
}

//brandInfo

type BrandInfo struct {
	BrandName string `json:"brandName"`
	BrandLogo string `json:"brandLogo"`
}
