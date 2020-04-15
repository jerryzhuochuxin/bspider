package model

type AuthorDo struct {
	ObjectId     string                 `bson:"objectId"`
	CRate        int                    `bson:"c_rate"`
	Mid          string                 `bson:"mid"`
	Name         string                 `bson:"name"`
	Face         string                 `bson:"face"`
	Official     string                 `bson:"official"`
	Sex          string                 `bson:"sex"`
	Data         map[string]interface{} `bson:"data"`
	Level        int                    `bson:"level"`
	Focus        bool                   `bson:"focus"`
	Pts          string                 `bson:"pts"`
	CFans        int                    `bson:"c_fans"`
	CAttention   int                    `bson:"c_attention"`
	CArchive     int                    `bson:"c_archive"`
	CArticle     int                    `bson:"c_article"`
	CArchiveView int                    `bson:"c_archiveView"`
	CArticleView int                    `bson:"c_articleView"`
	CLike        int                    `bson:"c_like"`
	CDatetime    int                    `bson:"c_dataTime"`
}

type VideoDo struct {
	Channel         string
	Aid             string
	Datetime        int
	Author          string
	Data            map[string]interface{}
	SubChannel      string
	Title           string
	Mid             string
	Pic             string
	CurrentView     int
	CurrentFavorite int
	CurrentDanmaku  int
	CurrentCoin     int
	CurrentShare    int
	CurrentLike     int
	CurrentDatetime int64
}

type VideoWithMidAidChannelsDo struct {
	Mid      string
	Aid      []string
	Channels map[string]interface{}
}

type SiteInfoDo struct {
	RegionCount map[string]interface{}
	AllCount    int
	WebOnline   int
	PlayOnline  int
}

type VideoOnlineDo struct {
	Title      string
	Author     string
	Data       map[string]interface{}
	Aid        string
	SubChannel string
	Channel    string
}
