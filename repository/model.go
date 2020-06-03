package repository

// Data struct
type Data struct {
	Selections struct {
		Makes []struct {
			Name   string `json:"name"`
			ID     string `json:"id"`
			Models []struct {
				Name  string `json:"name"`
				ID    string `json:"id"`
				Trims []struct {
					Name   string `json:"name"`
					ID     string `json:"id"`
					Styles []struct {
						Name string `json:"name"`
						ID   string `json:"id"`
					} `json:"styles"`
				} `json:"trims"`
			} `json:"models"`
		} `json:"makes"`
	} `json:"selections"`
}

type CarMake struct {
	MakeID               string   `protobuf:"bytes,3,opt,name=MakeID,proto3" json:"MakeID,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	LogoID               string   `protobuf:"bytes,2,opt,name=LogoID,proto3" json:"LogoID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type CarModel struct {
	ModelID              string   `protobuf:"bytes,4,opt,name=ModelID,proto3" json:"ModelID,omitempty"`
	MakeID               string   `protobuf:"bytes,3,opt,name=MakeID,proto3" json:"MakeID,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	LogoID               string   `protobuf:"bytes,2,opt,name=LogoID,proto3" json:"LogoID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type CarBodyTrim struct {
	BodyTrimID           string   `protobuf:"bytes,5,opt,name=BodyTrimID,proto3" json:"BodyTrimID,omitempty"`
	ModelID              string   `protobuf:"bytes,4,opt,name=ModelID,proto3" json:"ModelID,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	LogoID               string   `protobuf:"bytes,2,opt,name=LogoID,proto3" json:"LogoID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type CarStyle struct {
	StyleID              string   `protobuf:"bytes,6,opt,name=StyleID,proto3" json:"StyleID,omitempty"`
	BodyTrimID           string   `protobuf:"bytes,5,opt,name=BodyTrimID,proto3" json:"BodyTrimID,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	LogoID               string   `protobuf:"bytes,2,opt,name=LogoID,proto3" json:"LogoID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
