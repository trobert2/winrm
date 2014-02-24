package winrm

import (
	"encoding/xml"
    "io"
    "io/ioutil"
)

type ReferenceParameters struct{
    ResourceURI     string         `xml"w:ResourceURI"`
    SelectorSet     *SelectorSet   `xml"w:SelectorSet"`
}



type ResponseEnvelope struct{
    XMLName     xml.Name            `xml"s:Envelope,omitempty"`
    xmlnsS      string              `xml:"xmlns:s,attr,omitempty"`
    xmlnsA      string              `xml:"xmlns:a,attr,omitempty"`
    xmlnsX      string              `xml:"xmlns:x,attr,omitempty"`
    xmlnsW      string              `xml:"xmlns:w,attr,omitempty"`
    xmlnsRsp    string              `xml:"xmlns:rsp,attr,omitempty"`
    xmlnsP      string              `xml:"xmlns:p,attr,omitempty"`
    xmlnsLang   string              `xml:"xmlns:lang,attr,omitempty"`  
    Headers     *Headers            `xml"s:Header"`
    BodyStruct  *BodyStruct         `xml"s:Body"`
}


func GetObjectFromXML(XMLinput io.Reader) (ResponseEnvelope) {    
    b, _ := ioutil.ReadAll(XMLinput)
    var response ResponseEnvelope
    xml.Unmarshal(b, &response)
    return response
}
