package winrm

import (
	"encoding/xml"
    "io"
    "io/ioutil"
)

type Selector struct{
    Value string  `xml:",innerxml"`
    Name  string  `xml:"Name,attr"`
}

type SelectorSet struct{
    Selector    *Selector      `xml"w:Selector"`
}

type ReferenceParameters struct{
    ResourceURI     string         `xml"w:ResourceURI"`
    SelectorSet     *SelectorSet   `xml"w:SelectorSet"`
}

type ResourceCreated struct{
    Address                 string                 `xml"a:Address"`
    ReferenceParameters     *ReferenceParameters   `xml"a:ReferenceParameters"`
}

type Shell struct{
    // xmlnsRsp           string  `xml:"xmlns:rsp,attr,omitempty"`
    ShellId            string          `xml"rsp:ShellId,omitempty"`
    ResourceUri        string          `xml"rsp:ResourceUri,omitempty"`
    Owner              string          `xml"rsp:Owner,omitempty"`
    ClientIP           string          `xml"rsp:ClientIP,omitempty"`
    IdleTimeOut        string          `xml"rsp:IdleTimeOut,omitempty"`
    InputStreams       string          `xml"rsp:InputStreams,omitempty"`
    OutputStreams      string          `xml"rsp:OutputStreams,omitempty"`
    ShellRunTime       string          `xml"rsp:OutputStreams,omitempty"`
    ShellInactivity    string          `xml"rsp:OutputStreams,omitempty"`
    WorkingDirectory   string          `xml:"rsp:WorkingDirectory,omitempty"`
    Environment        *Environment    `xml:"rsp:Environment,omitempty"`
}

type Body struct{
    ResourceCreated *ResourceCreated   `xml"x:ResourceCreated"`
    Shell           *Shell             `xml"rsp:Shell"`
}

type EnvelopeAttr struct{
  
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
    Header          *Header         `xml"s:Header"`
    Body            *Body           `xml"s:Body"`
}


func GetObjectFromXML(XMLinput io.Reader) (ResponseEnvelope) {    
    b, _ := ioutil.ReadAll(XMLinput)
    var response ResponseEnvelope
    xml.Unmarshal(b, &response)
    return response
}
