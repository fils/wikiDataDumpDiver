package dm

import "time"

type WikiData struct {
	Aliases      Aliases      `json:"aliases"`
	Claims       Claims       `json:"claims"`
	Descriptions Descriptions `json:"descriptions"`
	ID           string       `json:"id"`
	Labels       Labels       `json:"labels"`
	Lastrevid    int          `json:"lastrevid"`
	Modified     time.Time    `json:"modified"`
	Ns           int          `json:"ns"`
	Pageid       int          `json:"pageid"`
	Sitelinks    Sitelinks    `json:"sitelinks"`
	Title        string       `json:"title"`
	Type         string       `json:"type"`
}
type En struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Fr struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type ZhMo struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Aliases struct {
	En   []En   `json:"en"`
	Fr   []Fr   `json:"fr"`
	ZhMo []ZhMo `json:"zh-mo"`
}
type Value struct {
	EntityType string `json:"entity-type"`
	ID         string `json:"id"`
	NumericID  int    `json:"numeric-id"`
}
type Datavalue struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}
type Mainsnak struct {
	Datatype  string    `json:"datatype"`
	Datavalue Datavalue `json:"datavalue"`
	Property  string    `json:"property"`
	Snaktype  string    `json:"snaktype"`
}
type P1151 struct {
	ID       string   `json:"id"`
	Mainsnak Mainsnak `json:"mainsnak"`
	Rank     string   `json:"rank"`
	Type     string   `json:"type"`
}
type P150 struct {
	ID       string   `json:"id"`
	Mainsnak Mainsnak `json:"mainsnak"`
	Rank     string   `json:"rank"`
	Type     string   `json:"type"`
}
type P580 struct {
	Datatype  string    `json:"datatype"`
	Datavalue Datavalue `json:"datavalue"`
	Hash      string    `json:"hash"`
	Property  string    `json:"property"`
	Snaktype  string    `json:"snaktype"`
}

// type Qualifiers struct {
// 	P580 []P580 `json:"P580"`
// }
type P582 struct {
	Datatype  string    `json:"datatype"`
	Datavalue Datavalue `json:"datavalue"`
	Hash      string    `json:"hash"`
	Property  string    `json:"property"`
	Snaktype  string    `json:"snaktype"`
}
type Qualifiers struct {
	P580 []P580 `json:"P580"`
	P582 []P582 `json:"P582"`
}
type P6 struct {
	ID              string     `json:"id"`
	Mainsnak        Mainsnak   `json:"mainsnak"`
	Qualifiers      Qualifiers `json:"qualifiers,omitempty"`
	QualifiersOrder []string   `json:"qualifiers-order"`
	Rank            string     `json:"rank,omitempty"`
	Type            string     `json:"type,omitempty"`
	// Qualifiers      Qualifiers `json:"qualifiers,omitempty"`
}
type P143 struct {
	Datatype  string    `json:"datatype"`
	Datavalue Datavalue `json:"datavalue"`
	Property  string    `json:"property"`
	Snaktype  string    `json:"snaktype"`
}
type Snaks struct {
	P143 []P143 `json:"P143"`
}
type References struct {
	Hash       string   `json:"hash"`
	Snaks      Snaks    `json:"snaks"`
	SnaksOrder []string `json:"snaks-order"`
}
type P625 struct {
	ID         string       `json:"id"`
	Mainsnak   Mainsnak     `json:"mainsnak"`
	Rank       string       `json:"rank"`
	References []References `json:"references"`
	Type       string       `json:"type"`
}
type P856 struct {
	ID       string   `json:"id"`
	Mainsnak Mainsnak `json:"mainsnak"`
	Rank     string   `json:"rank"`
	Type     string   `json:"type"`
}
type Claims struct {
	P1151 []P1151 `json:"P1151"`
	P150  []P150  `json:"P150"`
	P6    []P6    `json:"P6"`
	P625  []P625  `json:"P625"`
	P856  []P856  `json:"P856"`
}
type It struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Pl struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Ro struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Descriptions struct {
	En En `json:"en"`
	It It `json:"it"`
	Pl Pl `json:"pl"`
	Ro Ro `json:"ro"`
}
type Ar struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type My struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Ps struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Labels struct {
	Ar Ar `json:"ar"`
	En En `json:"en"`
	Fr Fr `json:"fr"`
	My My `json:"my"`
	Ps Ps `json:"ps"`
}
type Afwiki struct {
	Badges []interface{} `json:"badges"`
	Site   string        `json:"site"`
	Title  string        `json:"title"`
}
type Dewiki struct {
	Badges []string `json:"badges"`
	Site   string   `json:"site"`
	Title  string   `json:"title"`
}
type Dewikinews struct {
	Badges []interface{} `json:"badges"`
	Site   string        `json:"site"`
	Title  string        `json:"title"`
}
type Elwiki struct {
	Badges []interface{} `json:"badges"`
	Site   string        `json:"site"`
	Title  string        `json:"title"`
}
type Enwiki struct {
	Badges []interface{} `json:"badges"`
	Site   string        `json:"site"`
	Title  string        `json:"title"`
}
type Zhwikivoyage struct {
	Badges []interface{} `json:"badges"`
	Site   string        `json:"site"`
	Title  string        `json:"title"`
}
type Zuwiki struct {
	Badges []interface{} `json:"badges"`
	Site   string        `json:"site"`
	Title  string        `json:"title"`
}
type Sitelinks struct {
	Afwiki       Afwiki       `json:"afwiki"`
	Dewiki       Dewiki       `json:"dewiki"`
	Dewikinews   Dewikinews   `json:"dewikinews"`
	Elwiki       Elwiki       `json:"elwiki"`
	Enwiki       Enwiki       `json:"enwiki"`
	Zhwikivoyage Zhwikivoyage `json:"zhwikivoyage"`
	Zuwiki       Zuwiki       `json:"zuwiki"`
}
