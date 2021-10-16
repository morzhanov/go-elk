package doc

type Document struct {
	Author string `fake:"{firstname lastname}"`
	Title  string `fake:"{sentence:3}"`
	Text   string `fake:"{randomstring:[cars,crooked,attend,ubiquitous,annoying,event,quarrelsome,cemetery,High,there,how,he,rung,upon,the,rein,of,a,wimpling,wing,an,err,some,any,bad,good,anything,warm,cold,final,start,word]}"`
}
